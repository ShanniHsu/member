package sql

import "github.com/go-gormigrate/gormigrate/v2"

var List = []*gormigrate.Migration{
	CreateUser(),
}
