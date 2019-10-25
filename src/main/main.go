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
}
