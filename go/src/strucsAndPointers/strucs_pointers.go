package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "PONG")
}
func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}
func (myPc pc) String() string {
	return fmt.Sprintf("PC with %d GB de RAM %d GB of disk and %s brand", myPc.ram, myPc.disk, myPc.brand)
}
func main() {
	a := 50
	b := &a
	fmt.Println(a)
	fmt.Println(b)
	// if we want to print the value of [ b ] we should add * before de variable
	fmt.Println(*b)

	/* struct testing */
	myPc := pc{ram: 1, disk: 2, brand: "INTEL"}
	fmt.Println(myPc)
	myPc.ping()
	myPc.duplicateRam()
	myPc.duplicateRam()
	fmt.Println(myPc)

}
