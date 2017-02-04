package animals

import (
	"reflect"
	"testing"
)

func TestSetHobbiesSetsTheHobbies(t *testing.T) {
	hobbies := []string{"Barking"}
	dog := Dog{}

	dog.SetHobbies(hobbies)

	if !reflect.DeepEqual(dog.GetHobbies(), hobbies) {
		t.Fail()
	}
}
