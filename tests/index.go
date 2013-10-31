package tests

func (t AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html")
}

func (t AppTest) TestThatIndexPageContainsHackEdu() {
  t.Get("/")
  t.AssertContains("HackEDU")
}

func (t AppTest) TestThatIndexPageContainsLoginButton() {
  t.Get("/")
  t.AssertContains("Login")
}

func (t AppTest) TestThatIndexPageContainsSignUpButton() {
  t.Get("/")
  t.AssertContains("Sign Up")
}
