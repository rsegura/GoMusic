package server

import (
	//"github.com/vsouza/go-gin-boilerplate/config"
	"fmt"
	"GoMusic/rest"
)

func Init() {
	fmt.Println("Entramos en el INIT");
	//config := config.GetConfig()
	//r := NewRouter()
	rest.RunAPI(":9090")
	//r.Run(config.GetString("server.port"))
}
