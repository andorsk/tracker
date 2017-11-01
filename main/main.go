package main

import u "andortracker/controller/user"

func main() {
	a := u.App{}
	a.Initialize("root", "c0raline", "rest-api-example")
	a.Run(":8080")
}
