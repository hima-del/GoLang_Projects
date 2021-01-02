package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../00_images/78771293a.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println(f)
}
