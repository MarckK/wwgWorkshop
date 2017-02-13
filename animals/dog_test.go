package animals

import "testing"

//I can comment out almost all the dog testing and my coverage doesn't go down--because pets_tests.go is providing the coverage :)
//Still: 'PASS coverage: 69.2% of statements'
//I do need to keep test for Bark() as that functionality is specific to dog and is not being tested in pets_test.go

func TestBarkGetsWoofed(t *testing.T) {
	dog := Dog{}

	if dog.Bark() != "Woof!" {
		t.Fail()
	}
}

// func TestSetHobbiesSetsTheHobbies(t *testing.T) {
// 	hobbies := []string{"Barking"}
// 	dog := Dog{}
//
// 	dog.SetHobbies(hobbies)
//
// 	if !reflect.DeepEqual(dog.GetHobbies(), hobbies) {
// 		t.Fail()
// 	}
// }

// func TestGetHobbiesGetsTheHobbies(t *testing.T) {
// 	hobbies := []string{"Barking"}
// 	dog := Dog{Hobbies: hobbies}
//
// 	if dog.GetHobbies()[0] != hobbies[0] {
// 		t.Fail()
// 	}
// }

// func TestSetNameSetsTheName(t *testing.T) {
// 	dog := Dog{}
//
// 	dog.SetName("Spots")
//
// 	if dog.Name != "Spots" {
// 		t.Fail()
// 	}
// }

// func TestGetNameGetsTheName(t *testing.T) {
// 	dog := Dog{Name: "Spots"}
//
// 	if dog.GetName() != "Spots" {
// 		t.Fail()
// 	}
// }
