package Day12_Assign

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

// SomeStructWithTags ...
type SomeStructWithTags struct {
	UserName    string `faker:"username"`
	Email       string `faker:"email"`
	Password    string `faker:"password"`
	PhoneNumber string `faker:"phone_number"`
}

func Example_withTags() {

	a := SomeStructWithTags{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
}
