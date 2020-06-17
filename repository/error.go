package repository

import "github.com/pkg/errors"

var ErrNoAccountFound = errors.New("no account found")
var ErrShopNotFound = errors.New("shop not found")
