package app

import (
	"fmt"
	"go-cms/schema"
)

type User struct {
	*schema.User
	Website string `form:"website"`
}

func (p *User) Login()  {
	fmt.Println(p.Username)
	fmt.Println(p.Website)
}