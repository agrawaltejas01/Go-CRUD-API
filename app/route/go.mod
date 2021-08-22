module route

go 1.16

replace models => ../../models

replace controller => ../controller

replace service => ../service

replace data => ../data

require (
	controller v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)
