package server

import (
	"FirstSite/router"
)

func Init() {
	r := router.Router()
	r.Run()
}
