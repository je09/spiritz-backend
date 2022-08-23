package backend

func StartApp() {
	router := InitRoutes()
	router.Run(":8080")
}
