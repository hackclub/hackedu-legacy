package tests

import (
	"fmt"
	"github.com/hackedu/hackedu/app/models"
	"math/rand"
	"time"
)

func newUser() models.User {
	rand.Seed(time.Now().UnixNano())

	return models.User{
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     fmt.Sprintf("%s%d%s", "foo", rand.Intn(999999), "@bar.com"),
		Password:  "foobar",
	}
}
