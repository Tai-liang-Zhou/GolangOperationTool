package main

import (
	"fmt"

	"github.com/melbahja/goph"
)

func main() {
	client, err := goph.New("tlchoud", "192.168.1.114", goph.Password("Tomuwygnnr2A"))

	if err != nil {
		fmt.Println(err)
	}
	out, err := client.Run("ls ~/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	fmt.Println("test")

}
