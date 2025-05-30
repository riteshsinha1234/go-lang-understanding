package main

import (
	"fmt"

	"github.com/codersritesh/auth"
	"github.com/codersritesh/user"
)

func main() {
	auth.LoginWithCredintials("coderritesh","secret")

	session := auth.GetSession()
	fmt.Println("session",session)

	user := user.User{
		Email: "ritesh.sinha.737001@gmail.com",
		Name: "Ritesh",
	}
	fmt.Println(user.Email,user.Name)
}