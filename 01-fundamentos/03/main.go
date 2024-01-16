package main

type ID int

// Declaração GLOBAL
var (
	b bool = false
	c int = 10
	d string = "guipalm4"
	e float32 = 1.2
	f ID = 1
)

func main() {

	println(b)
	println(c)
	println(d)
	println(e)
	println("Atribuição...")
	b = true
	println(b)
	a:="Declaração shorthand"
	println(a)
	println(f)

}
