package animals

import (
	"reflect"
	"testing"
)

func KTestSetHobbiesSetsTheHobbies(t *testing.T) {
	hobbies := []string{"Hunting"}
	kitten := Kitten{}

	kitten.SetHobbies(hobbies)

	if !reflect.DeepEqual(kitten.GetHobbies(), hobbies) {
		t.Fail()
	}
}

func KTestGetHobbiesGetsTheHobbies(t *testing.T) {
	hobbies := []string{"Hunting"}
	kitten := Kitten{Hobbies: hobbies}

	if kitten.GetHobbies()[0] != hobbies[0] {
		t.Fail()
	}
}

func KTestSetNameSetsTheName(t *testing.T) {
	kitten := Kitten{}

	kitten.SetName("Pearl")

	if kitten.Name != "Pearl" {
		t.Fail()
	}
}

func KTestGetNameGetsTheName(t *testing.T) {
	kitten := Kitten{Name: "Pearl"}

	if kitten.GetName() != "Pearl" {
		t.Fail()
	}
}
