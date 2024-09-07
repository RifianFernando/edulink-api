package main

import (
	"github.com/skripsi-be/connections"
	"github.com/skripsi-be/models"
)

func init() {
	connections.LoadEnvVariables()
	connections.ConnecToDB()
}

func main() {
	connections.DB.AutoMigrate(&models.Student{})
	connections.DB.AutoMigrate(&models.Class{})
}
