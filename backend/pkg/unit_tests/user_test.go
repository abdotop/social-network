package unittests

import (
	octopus "backend/app"
	"backend/pkg/models"
	"fmt"
	"testing"
	// _ "github.com/mattn/go-sqlite3"
)

func TestUserOBject(t *testing.T) {

	user := models.User{
		Email:     "ddldsfdo@outlook.com",
		FirstName: "lolo",
		LastName:  "juju",
		AboutMe:   "blablablablablablablablablablablablablablablablablabl",
	}
	if err := user.Create(octopus.AppTest.Db.Conn); err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	userCLone := models.User{}

	if err := userCLone.Get(octopus.AppTest.Db.Conn, user.ID); err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	results := []bool{
		user.Email == userCLone.Email,
		user.FirstName == userCLone.FirstName,
		user.AboutMe == userCLone.AboutMe,
	}
	for _, result := range results {
		if !result {
			t.Errorf("user created is different from user cloned by the methode Ged ")
		}
	}

	fmt.Println("succes✅:  user created is the same as user cloned")

}
