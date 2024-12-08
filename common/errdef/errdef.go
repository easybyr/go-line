package errdef

import "errors"

var (
	PARAMS_ERROR = errors.New("params error")
	ETCD_CLIENT_NOT_INIT = errors.New("etcd client not initilized")
	ETCD_KEY_NOT_EXISTED = errors.New("etcd key not existed")
)


