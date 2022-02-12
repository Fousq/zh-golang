package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {
	helloMessage := emoji.Sprint("Hello :world_map:!")
	fmt.Println(helloMessage)
}
