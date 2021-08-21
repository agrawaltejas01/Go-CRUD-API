module crud-api

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	models v0.0.0-00010101000000-000000000000
	routes v0.0.0-00010101000000-000000000000
)

replace models => ./models

replace routes => ./routes
