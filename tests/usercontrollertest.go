package tests

import (
	"github.com/hackedu/hackedu/app/routes"
	"github.com/robfig/revel"
	"net/url"
)

type UserControllerTest struct {
	revel.TestSuite
}

func (t *UserControllerTest) Before() {
	println("Setting up user controller test suite.")
}

func (t *UserControllerTest) TestIndexWhenNotLoggedIn() {
	t.Get(routes.Users.Index())
	t.AssertContains("Please log in first.")
}

func (t *UserControllerTest) TestIndexWhenLoggedIn() {
	user := newUser()

	p := url.Values{}
	p.Set("user.FirstName", user.FirstName)
	p.Set("user.LastName", user.LastName)
	p.Set("user.Email", user.Email)
	p.Set("user.Password", user.Password)
	p.Set("verifyPassword", user.Password)
	t.PostForm("/register", p)

	t.Get(routes.Users.Index())
	t.AssertContains(user.FullName())
}

func (t *UserControllerTest) After() {
	println("Tearing down user controller test suite.")
}
