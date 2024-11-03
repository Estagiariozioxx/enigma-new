package main

import (
    "enigma-new/config"
    "enigma-new/routes"
)

// @title Enigma API
// @version 1.0
// @description Documentação da API do Enigma.
// @termsOfService http://swagger.io/terms/
// @contact.name Leonardo Lopes
// @contact.url https://www.linkedin.com/in/leonardo-lopes-49730a146/
// @contact.email leonardolopes-@hotmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /

func main() {
    config.InitDB()
    router := routes.SetupRouter()
    router.Run(":8080")
}


