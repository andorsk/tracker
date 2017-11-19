package main

import (
	"fmt"
	"log"
	"tracker/config"
	"tracker/controller"
	hbc "tracker/controller/heartbeat"
	hmi "tracker/model/heartbeat"
	umi "tracker/model/user"
	hbmodel "tracker/proto/heartbeat"
	umodel "tracker/proto/user"
	"tracker/server"
)

var s server.Server

func genFakeHeartbeats(count int) {

	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		var hb hbmodel.Heartbeat

		for j := 0; j < count*10; j++ {
			jf := int64(j)
			hb = hbmodel.Heartbeat{
				Timestamp: jf,
				Longitude: float64(jf + 50),
				Latitude:  float64(jf + 50),
				Altitude:  10,
				UserId:    int64(i),
			}

			if j == 0 {
				hbc.AddHeartbeatTrack(s.DB, hb)
			}
			hmi.Push(s.DB, hb)
		}
	}
}

func genFakeUserData(count int) {

	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		var user umodel.User
		var uname = fmt.Sprintf("user %v", i)
		var username = fmt.Sprintf("username %v", i)
		var email = fmt.Sprintf("user%v@email.com", i)

		user = umodel.User{
			Name:     uname,
			Age:      int32(i),
			Username: username,
			Email:    email,
		}
		fmt.Println("user is ", user)
		umi.Push(s.DB, &user)
	}
}

func initDB() {
	s = server.Server{}
	conf, err := config.LoadConfig("../../config/config.json")
	if err != nil {
		log.Fatal("Failed to load config")
	}

	fmt.Println("DB is ", conf.GetDB().DBName)
	s.Initialize(conf.GetDB().User, conf.GetDB().Password, conf.GetDB().DBName)
	mc := controller.MasterController{DB: s.DB, Router: s.Router}
	mc.InitializeRoutes(s.Router)

}

func clearTables() {
	s.DB.Exec("DELETE from users")
	s.DB.Exec("AlTER TABLE users MODIFY AUTO_INCREMENT=1")
	s.DB.Exec("DELETE from heartbeats")
	s.DB.Exec("ALTER TAbLE heartbeats MODIFY  AUTO_INCREMENT=1")
	s.DB.Exec("DELETE from tracks")
}

func main() {
	initDB()
	clearTables()
	genFakeUserData(5)
	genFakeHeartbeats(10)
}
