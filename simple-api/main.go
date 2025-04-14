package main

import (
	"net/http"

	"github.com/civilcoder55/go-lab/config"
	mongo "github.com/civilcoder55/go-lab/database"
	"github.com/civilcoder55/go-lab/routers"
)

func main() {
	config.LoadEnv()
	mongo.Connect()
	r := routers.GameRouter()
	http.Handle("/", r)
	http.ListenAndServe("127.0.0.1:3005", r)
}
