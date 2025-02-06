package sql

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"member/sql/v2412"
	v2501 "member/sql/v2501"
	v2502 "member/sql/v2502"
)

var List = []*gormigrate.Migration{
	v2412.CreateUser(),
	v2501.AddUserEmail(),
	v2501.CreateRestaurant(),
	v2502.CreateUserRestaurant(),
}
