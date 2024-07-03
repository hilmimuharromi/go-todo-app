package main

import (
	"todo-app/configs"
	"todo-app/routers"
)

func main() {
	configs.DatabaseConnect()
	defer configs.DisconnectDB(configs.DB)
	routers.SetupRouters()
}
