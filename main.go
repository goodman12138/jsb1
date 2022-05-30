package main

import (
	"fmt"
	"net/http"

	"go/conf/utils"
	"go/route"
)

func main() {

	route.Init()

	fmt.Println("监听开始")
	err := http.ListenAndServe(utils.Getport(), nil)
	if err != nil {
		panic(err)
	}
}
