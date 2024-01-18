package api

import (
	kafkaroutes "dmp-api/api/routes/kafka"
	postgresroutes "dmp-api/api/routes/postgres"

	"github.com/gofiber/fiber/v2"
)

func SetupPostgresRoutes(app *fiber.App) { //set up routes
	dmp_api := app.Group("/api/postgres")

	//------------get-------------
	//users
	dmp_api.Get("/get_user_by_username", postgresroutes.GetUserRecordByUsername)
	dmp_api.Get("/user_files", postgresroutes.RetrieveFiles)
	//files
	dmp_api.Get("/download_file", postgresroutes.DownloadFile)

	//------------post------------
	//users
	dmp_api.Post("/create_new_users", postgresroutes.CreateNewUsers)
	//files
	dmp_api.Post("/upload", postgresroutes.UploadFile)

	//forms
	dmp_api.Post("/submit_form_data_collection", postgresroutes.SubmitDataCollection)
	dmp_api.Post("/submit_form_data_planning", postgresroutes.SubmitDataPlanning)
	dmp_api.Post("/submit_form_data_processing_analysis", postgresroutes.SubmitDataProcessingAnalysis)
	dmp_api.Post("/submit_form_data_aq_qc", postgresroutes.SubmitDataQaQc)
}

func SetupKafkaRoutes(app *fiber.App) { //set up routes
	dmp_api := app.Group("/api/kafka")

	//------------post------------
	//forms
	dmp_api.Post("/submit_form_data_collection", kafkaroutes.SubmitDataCollection)
	dmp_api.Post("/submit_form_data_planning", kafkaroutes.SubmitDataPlanning)
	dmp_api.Post("/submit_form_data_processing_analysis", kafkaroutes.SubmitDataProcessingAnalysis)
	dmp_api.Post("/submit_form_data_aq_qc", kafkaroutes.SubmitDataQaQc)
}
