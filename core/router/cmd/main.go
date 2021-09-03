package main

func main()  {
	router := Router{
		ControllerPath: "controllers",
		RouterDir: "routers",
	}
	router.Build("","")
}