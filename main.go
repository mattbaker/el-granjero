package main

import (
	"fmt"
	"os"
)

func main() {
	key := os.Getenv("DESTINY_KEY")
	fmt.Printf("Key is: %s\n", key)
}
