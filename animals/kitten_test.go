package animals

//I had to put a K in front of all these test names or else golang complained (the compiler?)--in any case got an error message.  Is there a way to refactor these tests so test all the Pets for common functionality and then test animals, like Dog, for seperate functionality?

//Also: even with these additional tests my coverage has not gone above 38.5%

import "testing"

// func TestKSetHobbiesSetsTheHobbies(t *testing.T) {
// 	hobbies := []string{"Hunting"}
// 	kitten := Kitten{}
// 	AssertSetsHobbies(t, &kitten, hobbies)
// }
//
// func AssertSetsHobbies(t *testing.T, pet Pet, hobbies []string) {
// 	pet.SetHobbies(hobbies)
//
// 	if !reflect.DeepEqual(pet.GetHobbies(), hobbies) {
// 		t.Fail()
// 	}
// }

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
