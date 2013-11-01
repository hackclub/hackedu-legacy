package tests

import "github.com/robfig/revel"

type UserModelTest struct {
	revel.TestSuite
}

func (t *UserModelTest) Before() {
	println("Setting up user model test suite.")
}

func (t *UserModelTest) TestString() {
	user := newUser()
	user.Email = "foo@bar.com"

	t.AssertEqual("User(foo@bar.com)", user.String())
}

func (t *UserModelTest) TestFullName() {
	user := newUser()
	user.FirstName = "Foo"
	user.LastName = "Bar"

	t.AssertEqual("Foo Bar", user.FullName())
}

func (t *UserModelTest) After() {
	println("Tearing down user godel test suite.")
}
