package main

import (
	"fmt"
	"strconv"

	"github.com/abhaydiwan/sweatbead/samplemgr/db"
	"github.com/abhaydiwan/sweatbead/samplemgr/service"

	"github.com/urfave/negroni"
)

//it is a main file....
func main() {
	//Initialize DB
	err, db := db.GetDB()
	if err != nil {
		println(err)
		panic(err)
	}
	fmt.Println("Db initialized: ", db)
	//mux router
	router := service.InitRouter()
	//init web server
	server := negroni.Classic()
	server.UseHandler(router)

	port := 3301
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	server.Run(addr)
}
