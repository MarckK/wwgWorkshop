package animals

type Kitten struct {
	Name    string
	Hobbies []string
	Like    int
}

func (k *Kitten) SetName(name string) {
	k.Name = name
}

func (k *Kitten) GetName() string {
	return k.Name
}

func (k *Kitten) SetHobbies(hobbies []string) {
	k.Hobbies = hobbies
}

func (k *Kitten) GetHobbies() []string {
	return k.Hobbies
}

func (k *Kitten) IncrementLikeCounter() int {
	k.Like += 1
	return k.Like
}
