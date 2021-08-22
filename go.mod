module crud-api

go 1.16

require (
	controller v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	models v0.0.0-00010101000000-000000000000
)

replace models => ./models

replace controller => ./app/controller
