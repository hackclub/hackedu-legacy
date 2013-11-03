package tests

import (
	"github.com/hackedu/hackedu/app/models"
	"github.com/robfig/revel"
	"net/url"
)

type AuthenticationTest struct {
	revel.TestSuite
}

func (t *AuthenticationTest) Before() {
	println("Setting up authentication test suite.")
}

func (t *AuthenticationTest) TestThatRegisterPageWorks() {
	t.Get("/register")
	t.AssertOk()
	t.AssertContentType("text/html")
}

func (t *AuthenticationTest) TestThatLoginPageWorks() {
	t.Get("/login")
	t.AssertOk()
	t.AssertContentType("text/html")
}

func (t *AuthenticationTest) TestRegistrationWithValidInfo() {
	user := newUser()
	submitUserRegisterForm(t, &user)

	t.AssertContains("Welcome")
}

func (t *AuthenticationTest) TestRegistrationWithInvalidInfo() {
	user := newUser()
	user.Email = "foo"

	submitUserRegisterForm(t, &user)

	t.AssertContains("You must provide a valid email address.")
}

func (t *AuthenticationTest) TestRegistrationWithDuplicateEmail() {
	user := newUser()

	submitUserRegisterForm(t, &user)
	submitUserRegisterForm(t, &user)

	t.AssertContains("Email is taken.")
}

func (t *AuthenticationTest) TestLoginWithValidInfo() {
	user := newUser()
	submitUserRegisterForm(t, &user)

	submitUserLoginForm(t, &user)

	t.AssertContains("Welcome")
}

func (t *AuthenticationTest) TestLoginWithInvalidInfo() {
	user := newUser()

	submitUserLoginForm(t, &user)

	t.AssertContains("Login failed!")
}

func (t *AuthenticationTest) TestLogout() {
	user := newUser()

	submitUserRegisterForm(t, &user)

	t.Get("/logout")
	t.AssertContains("Register")
}

func (t *AuthenticationTest) After() {
	println("Tearing down authentication test suite.")
}

func submitUserLoginForm(t *AuthenticationTest, user *models.User) {
	p := url.Values{}
	p.Set("email", user.Email)
	p.Set("password", user.Password)
	t.PostForm("/login", p)
}

func submitUserRegisterForm(t *AuthenticationTest, user *models.User) {
	p := url.Values{}
	p.Set("user.FirstName", user.FirstName)
	p.Set("user.LastName", user.LastName)
	p.Set("user.Email", user.Email)
	p.Set("user.Password", user.Password)
	p.Set("verifyPassword", user.Password)
	t.PostForm("/register", p)
}
