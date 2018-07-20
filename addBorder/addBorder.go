package main

import "fmt"

func main() {
	fmt.Println(addBorder([]string{
		"abc",
		"ded",
	}))
}

func addBorder(picture []string) (pic []string) {
	size := len(picture[0]) + 2
	var edge string

	for i := 0; i < size; i++ {
		edge += "*"
	}

	pic = append(pic, edge)

	for _, item := range picture {
		pic = append(pic, "*"+item+"*")
	}

	pic = append(pic, edge)
	return
}
