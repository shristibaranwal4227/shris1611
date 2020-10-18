package main

func main() {
	scheduling := scheduling.Getscheduling()

	app := &app.Application{}
	app.Initialize(config)
	app.Run(":3000")
}
