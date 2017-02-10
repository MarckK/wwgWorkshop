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

func TestGetHobbiesGetsTheHobbies(t *testing.T) {
	hobbies := []string{"Barking"}
	dog := Dog{Hobbies: hobbies}

	if dog.GetHobbies()[0] != hobbies[0] {
		t.Fail()
	}
}

func TestSetNameSetsTheName(t *testing.T) {
	dog := Dog{}

	dog.SetName("Spots")

	if dog.Name != "Spots" {
		t.Fail()
	}
}

func TestGetNameGetsTheName(t *testing.T) {
	dog := Dog{Name: "Spots"}

	if dog.GetName() != "Spots" {
		t.Fail()
	}
}

func TestBarkGetsWoofed(t *testing.T) {
	dog := Dog{}

	if dog.Bark() != "Woof!" {
		t.Fail()
	}
}
