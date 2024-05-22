package bucket

import (
	"archetype/app/shared/infrastructure/observability"
	"archetype/app/shared/infrastructure/storj"
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"go.opentelemetry.io/otel/trace"
	"storj.io/uplink"
	"storj.io/uplink/edge"
)

type StorJBucket struct {
	fileExpiration  time.Duration
	sharedLinkCreds *edge.Credentials
	bucketName      string
	upLink          *storj.Uplink
}

func init() {
	ioc.Registry(
		NewStorJBucket,
		storj.NewUplink)
}

func NewStorJBucket(ul *storj.Uplink) (storj.UplinkManager, error) {
	sharedLinkExpiration := 2 * time.Minute
	fileExpiration := 7 * 24 * time.Hour
	bucketName := "insert-your-bucket-name"
	bucketFolderName := ""
	sharedLinkRestrictedAccess, err := ul.Access.Share(
		uplink.Permission{
			// only allow downloads
			AllowDownload: true,
			// this allows to automatically cleanup the access grants
			NotAfter: time.Now().Add(sharedLinkExpiration),
		}, uplink.SharePrefix{
			Bucket: bucketName,
			Prefix: bucketFolderName,
		},
	)
	if err != nil {
		return StorJBucket{}, fmt.Errorf("could not restrict access grant: %w", err)
	}

	// RegisterAccess registers the credentials to the linksharing and s3 sites.
	// This makes the data publicly accessible, see the security implications in https://docs.storj.io/dcs/concepts/access/access-management-at-the-edge.
	ctx := context.Background()
	credentials, err := ul.Config.RegisterAccess(ctx,
		sharedLinkRestrictedAccess,
		&edge.RegisterAccessOptions{Public: true})
	if err != nil {
		return StorJBucket{}, fmt.Errorf("could not register access: %w", err)
	}

	// Ensure the desired Bucket within the Project is created.
	_, err = ul.Project.EnsureBucket(ctx, bucketName)
	if err != nil {
		return StorJBucket{}, fmt.Errorf("could not ensure bucket: %v", err)
	}

	bucket := StorJBucket{
		fileExpiration:  fileExpiration,
		sharedLinkCreds: credentials,
		bucketName:      bucketName,
		upLink:          ul,
	}
	return bucket, nil
}

func (b StorJBucket) CreatePublicSharedLink(ctx context.Context, objectKey string) (string, error) {
	_, span := observability.Tracer.Start(ctx,
		"NewStorJBucketCreatePublicSharedLink",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()
	// Create a public link that is served by linksharing service.
	url, err := edge.JoinShareURL("https://link.storjshare.io",
		b.sharedLinkCreds.AccessKeyID,
		b.bucketName, objectKey, nil)
	if err != nil {
		span.RecordError(err)
		return "", fmt.Errorf("could not create a shared link: %w", err)
	}
	return url, nil
}

func (b StorJBucket) Upload(ctx context.Context, objectKey string, dataToUpload []byte) error {
	// Intitiate the upload of our Object to the specified bucket and key.
	upload, err := b.upLink.Project.UploadObject(ctx, b.bucketName, objectKey, &uplink.UploadOptions{
		// It's possible to set an expiration date for data.
		Expires: time.Now().Add(b.fileExpiration),
	})
	if err != nil {
		return fmt.Errorf("could not initiate upload: %v", err)
	}
	// Copy the data to the upload.
	buf := bytes.NewBuffer(dataToUpload)
	_, err = io.Copy(upload, buf)
	if err != nil {
		_ = upload.Abort()
		return fmt.Errorf("could not upload data: %v", err)
	}

	// Commit the uploaded object.
	err = upload.Commit()
	if err != nil {
		return fmt.Errorf("could not commit uploaded object: %v", err)
	}
	return nil
}
