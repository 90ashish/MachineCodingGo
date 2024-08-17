package main

import (
	basic "cache/Basic"
	moderate "cache/moderate"
	"fmt"
)

func main(){
	fmt.Println("Machine Coding Round - Go")

	// basic
	basic.BasicRedisCache()

	// Moderate 
	moderate.ModerateLRUCache()
}

