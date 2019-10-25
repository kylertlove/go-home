package main

import (
	media "./media"
)

func main() {
	a := media.CreateMovie("Jaws", 1978)
	println(a.Name)
	println(a.Good)

	b := media.CreateRecord("reads", []string{"one", "two", "three"})
	println(b.ArtistName)

	for index := 0; index < len(b.Albums); index++ {
		println("album arr:", b.Albums[index])
	}
}
