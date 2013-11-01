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

func (t AppTest) TestThatIndexPageContainsAuthenticationButtons() {
  t.Get("/")
  t.AssertContains("Login")
  t.AssertContains("Sign Up")
}
