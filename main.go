package main

import (
	"dmp-api/api"
	apiconfig "dmp-api/api/api_config"
	"dmp-api/logger"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var version = "0.8"

func main() {
	//extract flags, set their defaults, and description
	var ex, _ = os.Executable()
	var defaultLogDir = filepath.Dir(ex) + "/logs/"
	verbose := flag.Bool("verb", false, "True outputs logs to console AND file. False only outputs to file.")
	loggerDir := flag.String("ld", defaultLogDir, "Set log directory.")
	help := flag.Bool("h", false, "Show help screen.")
	apiPort := flag.String("p", "8080", "Sets the port to bind the API to. Default is 8080.")
	apiIP := flag.String("ip", "0.0.0.0", "Sets the IP for the API to listen on. Default is 0.0.0.0")

	flag.Parse()

	//show all of the flags and quit if help requested
	if *help {
		flag.Usage()
		return
	}

	//init the logger
	logger.SetVerbosity(*verbose)
	logger.ChangeDir(*loggerDir)

	logger.LogMessage("Starting... version: " + version)

	//init the API ports and address
	apiconfig.SetAPIIP(*apiIP)
	apiconfig.SetAPIPort(*apiPort)

	//set up new connection, along with what API calls are allowed
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*", //note that this is for an intranet
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	//set up routes
	api.SetupPostgresRoutes(app)
	api.SetupKafkaRoutes(app)
	fmt.Println(*apiPort)

	//API listening start
	err := app.Listen(apiconfig.GetFullAPIAddress())
	if err != nil {
		logger.LogError(err.Error())
		return
	}
}
