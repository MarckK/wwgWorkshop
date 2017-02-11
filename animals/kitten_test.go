package animals

import (
	"reflect"
	"testing"
)

func TestKSetHobbiesSetsTheHobbies(t *testing.T) {
	hobbies := []string{"Hunting"}
	kitten := Kitten{}

	kitten.SetHobbies(hobbies)

	if !reflect.DeepEqual(kitten.GetHobbies(), hobbies) {
		t.Fail()
	}
}

func TestKGetHobbiesGetsTheHobbies(t *testing.T) {
	hobbies := []string{"Hunting"}
	kitten := Kitten{Hobbies: hobbies}

	if kitten.GetHobbies()[0] != hobbies[0] {
		t.Fail()
	}
}

func TestKSetNameSetsTheName(t *testing.T) {
	kitten := Kitten{}

	kitten.SetName("Pearl")

	if kitten.Name != "Pearl" {
		t.Fail()
	}
}

func TestKGetNameGetsTheName(t *testing.T) {
	kitten := Kitten{Name: "Pearl"}

	if kitten.GetName() != "Pearl" {
		t.Fail()
	}
}
