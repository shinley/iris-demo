package main

import (
	"iris-demo/bootstrap"
	//"github.com/yz124/superstar/web/middleware/identity"
	"iris-demo/router"
)

func main() {
	app := bootstrap.New("Superstar database", "一凡Sir")
	app.Bootstrap()
	//app.Configure(identity.Configure, routes.Configure)
	app.Configure(routes.Configure)
	app.Listen(":8081")
}
