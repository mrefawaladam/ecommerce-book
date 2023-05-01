package main

import (
	r "ebook/internal/adapters/http"
)

func main() {
	e := r.InitRoutes()
	e.Logger.Fatal(e.Start(":1323"))
}
