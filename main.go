package main

import (
	"fmt"
	"gitlab.com/wzcourses/demo-gads-app/routers"
)

func main() {
	fmt.Print("hello belajar gitlab administration")

	r := routers.AppRouter()

	err := r.Run(":8080")
	if err != nil {
		panic("error handle run")
	}

}
