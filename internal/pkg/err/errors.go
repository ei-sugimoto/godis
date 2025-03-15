package err

import "errors"

var ErrPortEmpty = errors.New("port is empty")

var ErrConfigPathEmpty = errors.New("config path is empty")

var ErrNoListener = errors.New("no listener")
