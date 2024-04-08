package unittests

import (
	octopus "backend/app"
	"backend/pkg/models"
	"fmt"
	"testing"
)

func TestPostObject(t *testing.T) {

	post := models.Post{
		Title:   "hello world",
		Content: "hello senegal",
		Privacy: models.PrivacyPublic,
		//etc ........
	}
	if err := post.Create(octopus.AppTest.Db.Conn); err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}

	clonedpost := models.Post{}

	if err := clonedpost.Get(octopus.AppTest.Db.Conn, post.ID); err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	results := []bool{
		post.Title == clonedpost.Title,
		post.Content == clonedpost.Content,
		post.Privacy == clonedpost.Privacy,
	}
	for _, result := range results {
		if !result {
			t.Errorf("post  created is different from user cloned by the methode Ged ")
		}
	}

	fmt.Println("succes✅: the post created is the same as the post cloned")

}
