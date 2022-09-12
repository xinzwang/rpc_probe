package main

import (
	"flag"
	"fmt"
)


func main(){
	name := flag.String("name", "", "Name of RPC function")
	flag.Parse()

	// reflect.Call()
}