package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/hackedu/hackedu/app/models"
	"github.com/hackedu/hackedu/app/routes"
	"github.com/robfig/revel"
)

type Authentication struct {
	GorpController
}

func (c Authentication) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}

func (c Authentication) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if email, ok := c.Session["user"]; ok {
		return c.getUser(email)
	}
	return nil
}

func (c Authentication) Login() revel.Result {
	return c.Render()
}

func (c Authentication) Register() revel.Result {
	return c.Render()
}

func (c Authentication) LoginUser(email, password string) revel.Result {
	user := c.getUser(email)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = email
			c.Flash.Success("Welcome " + user.FirstName + "!")
			return c.Redirect(routes.App.Index())
		}
	}

	c.Flash.Out["email"] = email
	c.Flash.Error("Login failed!")
	return c.Redirect(routes.Authentication.Login())
}

func (c Authentication) RegisterUser(user models.User, verifyPassword string) revel.Result {
	email, err := c.Txn.SelectStr("select Email from User where Email = ?",
		user.Email)
	if email != "" && err == nil {
		c.Validation.Error("Email is taken.")
	}

	c.Validation.Required(verifyPassword).Message("You must verify your password.")
	c.Validation.Required(verifyPassword == user.Password).
		Message("Password does not match")
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Authentication.Register())
	}

	user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	err = c.Txn.Insert(&user)
	if err != nil {
		panic(err)
	}

	c.Session["user"] = user.Email
	c.Flash.Success("Welcome " + user.FirstName + "!")
	return c.Redirect(routes.App.Index())
}

func (c Authentication) LogoutUser() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}

func (c Authentication) getUser(email string) *models.User {
	users, err := c.Txn.Select(models.User{}, `select * from User where Email = ?`, email)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*models.User)
}
