package main

import "fmt"

type animal interface {
	move() string
}

type dog struct{}
type bird struct{}
type fish struct{}

func (fish) move() string {
	return "Am a fish and I swim"
}
func (bird) move() string {
	return "Am a bird and a fly"
}

func (dog) move() string {
	return "Am a dog and I walk"
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func main() {
	dog := dog{}
	fish := fish{}
	bird := bird{}

	moveAnimal(dog)
	moveAnimal(fish)
	moveAnimal(bird)

}
