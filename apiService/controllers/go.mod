module example.com/controllers

go 1.16

replace example.com/models => ../models

require (
	example.com/models v0.0.0-00010101000000-000000000000 // indirect
	github.com/gin-gonic/gin v1.7.2
)
