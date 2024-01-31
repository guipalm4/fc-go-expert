package main

// Declaração GLOBAL
var (
	b bool
	c int
	d string
	e float32
)

func main() {

	println(b)
	println(c)
	println(d)
	println(e)
	println("Atribuição...")
	b = true
	println(b)
	a := "Declaração shorthand"
	println(a)

}
