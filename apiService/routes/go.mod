module example.com/routes

go 1.16

replace example.com/controllers => ../controllers

require (
	example.com/controllers v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.2
)
