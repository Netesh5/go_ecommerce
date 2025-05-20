package main

import "github.com/netesh5/go_ecommerce/internal/config"

func main() {
	config := config.LoadConfig()
	println("Config loaded successfully", config.Env)
	// if _, err := os.Stat(".env"); err != nil {
	// 	log.Println("Error loading .env file")
	// }
	// log.Println("Loading .env file")

}
