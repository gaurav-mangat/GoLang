package main

import "fmt"

type dog struct {
	name string
}

func (d *dog) walk() {
	fmt.Println(d.name, "is walking....")
}
func (d *dog) run() {
	fmt.Println(d.name, "is running....")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()

}

func main() {
	d1 := &dog{"Tommy"}
	d1.run()
	d1.walk()
	youngRun(d1)
	d2 := dog{"Kaalu"}
	d2.run()
	d2.walk()
}
