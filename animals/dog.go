package animals

type Dog struct {
	Name  string
	Sound string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d *Dog) GetName() string {
	return d.Name
}

func (d *Dog) Bark() string {
	return "Woof!"
}
