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
