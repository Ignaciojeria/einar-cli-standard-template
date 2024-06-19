package systemerr

import "errors"

var EXTERNAL_SERVER_ERROR = errors.New("EXTERNAL_SERVER_ERROR")
var EXTERNAL_UNKNOWN_ERROR = errors.New("EXTERNAL_UNKNOWN_ERROR")
var HTTP_NETWORK_ERROR = errors.New("HTTP_NETWORK_ERROR")
var INTERNAL_SERVER_ERROR = errors.New("INTERNAL_SERVER_ERROR")
var PUBSUB_BROKER_ERROR = errors.New("PUBSUB_BROKER_ERROR")