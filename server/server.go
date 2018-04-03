package server

func Start() {
	router := InitRoutes()
	router.Run(":8888")
}
