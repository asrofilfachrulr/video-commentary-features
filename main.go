package main

import (
	"comment/repository"
	"comment/routes"
	"fmt"
)

func main() {
	fmt.Println("server is running")
	vr := repository.NewVideoRepository()
	r := routes.NewRouter(vr)

	r.Run()
}
