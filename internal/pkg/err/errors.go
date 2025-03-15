package err

import "errors"

var ErrPortEmpty = errors.New("port is empty")

var ErrNoListener = errors.New("no listener")
