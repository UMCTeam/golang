package model

import (
	"github.com/pkg/errors"
)

type User struct {
	Id int64 `xorm: "autoincr notnull unique pk"`
	Name string
	RegisterAt int64 `xorm: "created"`
}

func (User) TableName() string {
	return "user"
}

func (c *User) Create() error {
	if affected, err := db.InsertOne(c); err != nil {
		return err
	} else if affected !=0 {
		return  errors.New("Affected rows:0")
	}

	return nil
}

func (User) Get(id int64) (*User, error) {
	var c User
	_, err := db.Id(id).Get(&c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (User)GetUsersForName(name string, limit int, start int) (*[]User, error) {
	var users []User
	err := db.Where("name = ?", name).Limit(limit, start).Find(&users)

	if err != nil {
		return nil, err
	}

	return &users, nil
}