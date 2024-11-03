package lab6 

import "fmt"

type Mouse struct {
	Name  string
	Age   int
	Breed string
}

func NewMouse(name string, age int, size string) Mouse {
	return Mouse{
		Name:  name,
		Age:   age,
		Size: size,
	}
}
func (d *Mouse) GetAge() int {
	return d.Age
}
func (d *Mouse) SetAge(age int) {
	d.Age = age
}
func (d *Mouse) Squeak() string {
	return "Peep!"
}

func Startlab6() {
	dog := NewMouse("Ster", 5, "dnsmw")
	fmt.Println("Возраст мышки:", mouse.GetAge())
	dog.SetAge(6)
	fmt.Println("Новый возраст мышки:", mouse.GetAge())
	fmt.Println("Мышка пищит:", mouse.Squeak())
}