package tests

import "github.com/robfig/revel"

type AppTest struct {
	revel.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html")
}

func (t AppTest) TestThatIndexPageContainsHackEdu() {
	t.Get("/")
	t.AssertContains("HackEDU")
}

func (t AppTest) TestThatIndexPageContainsAuthenticationButtons() {
	t.Get("/")
	t.AssertContains("Login")
	t.AssertContains("Register")
}

func (t *AppTest) After() {
	println("Tear down")
}
