module example.com/apiService

go 1.16

replace example.com/routes => ./routes

replace example.com/controllers => ./controllers

replace example.com/models => ./models

require (
	example.com/controllers v0.0.0-00010101000000-000000000000 // indirect
	example.com/models v0.0.0-00010101000000-000000000000 // indirect
	example.com/routes v0.0.0-00010101000000-000000000000
	github.com/mattn/go-isatty v0.0.12 // indirect
	gorm.io/driver/postgres v1.1.0 // indirect
	gorm.io/gorm v1.21.11 // indirect
)
