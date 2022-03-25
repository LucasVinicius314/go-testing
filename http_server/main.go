package main

import "fmt"

type Thing struct {
	id   int
	name string
}

func main() {
	var t = 3

	var k = Thing{id: 3, name: "sapato"}

	fmt.Println(k.id)

	var stuff = []Thing{{id: 3, name: "sapato"}, {id: 4, name: "marcelo"}}

	for u := 0; u < len(stuff); {
		var thing = stuff[u]

		fmt.Println(thing)

		u++
	}

	g := 4

	for t >= 0 {
		fmt.Println(g)

		t--
	}

	var e = make([]int, 5)

	for i := 0; i < len(e); {
		fmt.Println(e[i])

		i++
	}

	defer fmt.Println("b")

	fmt.Println("a")
}
