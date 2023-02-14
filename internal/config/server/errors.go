package server

import "errors"

var ErrorNoAddressCert = errors.New("No path for the cert")
var ErrorNoAddressCertKey = errors.New("No path for the cert key")
var ErrorNoDSN = errors.New("No DSN for database")
var ErrorNoJWTKEY = errors.New("No key for JWT")
