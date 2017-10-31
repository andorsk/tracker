package main

func main() {
	a := App{}
	a.Initialize("root", "c0raline", "rest-api-example")
	a.Run(":8080")
}
