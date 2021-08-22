module controller

go 1.16

replace models => ../../models

replace data => ../data

replace service => ../service

require (
	github.com/gorilla/mux v1.8.0
	service v0.0.0-00010101000000-000000000000
)
