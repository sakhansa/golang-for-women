package main

import "assignment-project-rest-api/routers"

var PORT = ":8080"

func main() {
	routers.StartServer().Run(PORT)

}
