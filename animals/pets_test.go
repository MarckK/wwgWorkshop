package animals

import (
	"reflect"
	"testing"
)

func TestSetHobbiesSetsTheHobbiesForAllPets(t *testing.T) {
	kittenHobbies := []string{"Hunting"}
	kitten := Kitten{}
	AssertSetsHobbies(t, &kitten, kittenHobbies)

	dogHobbies := []string{"Barking"}
	dog := Dog{}
	AssertSetsHobbies(t, &dog, dogHobbies)
}

func AssertSetsHobbies(t *testing.T, pet Pet, hobbies []string) {
	pet.SetHobbies(hobbies)

	if !reflect.DeepEqual(pet.GetHobbies(), hobbies) {
		t.Fail()
	}
}

func TestGetHobbiesGetsTheHobbiesForAllPets(t *testing.T) {
	kittenHobbies := []string{"Hunting"}
	kitten := Kitten{Hobbies: kittenHobbies}
	AssertGetsHobbies(t, &kitten, kittenHobbies)

	dogHobbies := []string{"Barking"}
	dog := Dog{Hobbies: dogHobbies}
	AssertSetsHobbies(t, &dog, dogHobbies)
}

func AssertGetsHobbies(t *testing.T, pet Pet, hobbies []string) {
	pet.GetHobbies()

	if pet.GetHobbies()[0] != hobbies[0] {
		t.Fail()
	}
}

func TestSetNameSetsTheNameForAllPets(t *testing.T) {
	kitten := Kitten{}
	AssertSetNameSetsTheName(t, &kitten, "Pearl")

	dog := Dog{}
	AssertSetNameSetsTheName(t, &dog, "Beau")

}

func AssertSetNameSetsTheName(t *testing.T, pet Pet, name string) {
	pet.SetName(name)

	switch p := pet.(type) {
	case *Kitten:
		if p.Name != name {
			t.Fail()
		}
	case *Dog:
		if p.Name != name {
			t.Fail()
		}
	}
}

func TestGetNameGetsTheNameForAllPets(t *testing.T) {
	kitten := Kitten{Name: "Pearl"}

	if kitten.GetName() != "Pearl" {
		t.Fail()
	}
}

// func AssertGetNameGetsTheName(t *testing.T, pet Pet, name string) {
//   pet.GetName(name)
//
//   if kitten, ok := pet.(*Kitten); ok {
//     if
//   }
// }

// func TestSetNameSetsTheNameAndTestGetNameGetsTheNameForAllPets(t *testing.T) {
// 	kitten := Kitten{}
// 	AssertSetNameSetsTheNameAndTestGetNameGetsTheNameForAllPets(t, &kitten, "Pearl")
//
// 	dog := Dog{}
// 	AssertSetNameSetsTheNameAndTestGetNameGetsTheNameForAllPets(t, &dog, "Beau")
// }
//
// func AssertSetNameSetsTheNameAndTestGetNameGetsTheNameForAllPets(t *testing.T, pet Pet, name string) {
// 	pet.SetName(name)
//
// 	if pet.GetName() != name {
// 		t.Fail()
// 	}
// }
