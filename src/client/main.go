package main

import (
	"fmt"
)

func main() {
	config := GetConfiguration()
	fmt.Printf("%#v", config)
}
