// these are the handlers for the form submission
package postgres_routes

import (
	"dmp-api/api/db_kafka_env_config"
	jsonmodels "dmp-api/api/json_models"
	"dmp-api/logger"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SubmitDataCollection(context *fiber.Ctx) error {
	//uuid or username is accepted as it's handled hy the GetUserRecord function below
	uuid := context.FormValue("uuid_fk")
	username := context.FormValue("username")
	powerpointFileName := context.FormValue("powerpoint_file_name")
	isNewData := context.FormValue("is_this_new_data")
	datasetName := context.FormValue("dataset_name")
	dataStorageRequirements := context.FormValue("data_storage_requirements")
	sessionModel := &jsonmodels.SessionRecord{}
	dataCollectionModel := &jsonmodels.FormDataCollection{}

	err := GetUserRecord(uuid, context, sessionModel, username)
	if err != nil {
		logger.LogError(err.Error())
	}

	//set variables - no need to check as the separation of concerns is on the databases
	dataCollectionModel.UuidFK = sessionModel.Uuid
	dataCollectionModel.PowerPointFileName = powerpointFileName
	dataCollectionModel.IsNewData = isNewData
	dataCollectionModel.DatasetName = datasetName
	dataCollectionModel.DataStorageRequirements = dataStorageRequirements

	dbCreateErr := db_kafka_env_config.DMP_DB.Create(&dataCollectionModel).Error
	if dbCreateErr != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "err_creating_user"})
		logger.LogError("Error creating new user: " + err.Error())
		return err
	}
	return nil
}

func SubmitDataPlanning(context *fiber.Ctx) error {
	//uuid or username is accepted as it's handled hy the GetUserRecord function below
	uuid := context.FormValue("uuid_fk")
	username := context.FormValue("username")
	projectTitle := context.FormValue("project_title")
	publicationTitle := context.FormValue("publication_title")
	projectDescription := context.FormValue("project_description")
	dataOwnership := context.FormValue("data_ownership")
	//should add parser to this to custom type
	projectStatus := jsonmodels.ProjectStatus(context.FormValue("project_status"))
	projectTimeFrameStart := context.FormValue("project_time_frame_start")
	projectTimeFrameEnd := context.FormValue("project_time_frame_end")
	shapeFileName := context.FormValue("shapefile_name")

	sessionModel := &jsonmodels.SessionRecord{}
	dataCollectionModel := &jsonmodels.FormDataPlanning{}

	err := GetUserRecord(uuid, context, sessionModel, username)
	if err != nil {
		logger.LogError(err.Error())
	}

	//set variables - no need to check as the separation of concerns is on the databases
	dataCollectionModel.UuidFK = sessionModel.Uuid
	dataCollectionModel.ProjectTitle = projectTitle
	dataCollectionModel.PublicationTitle = publicationTitle
	dataCollectionModel.ProjectDescription = projectDescription
	dataCollectionModel.DataOwnership = dataOwnership
	dataCollectionModel.ProjectStatus = projectStatus
	dataCollectionModel.ProjectTimeFrameStart = projectTimeFrameStart
	dataCollectionModel.ProjectTimeFrameEnd = projectTimeFrameEnd
	dataCollectionModel.ShapeFileName = shapeFileName

	dbCreateErr := db_kafka_env_config.DMP_DB.Create(&dataCollectionModel).Error
	if dbCreateErr != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "err_creating_data_planning"})
		logger.LogError("Error creating error planning: " + err.Error())
		return err
	}
	return nil
}

func SubmitDataProcessingAnalysis(context *fiber.Ctx) error {
	//uuid or username is accepted as it's handled hy the GetUserRecord function below
	uuid := context.FormValue("uuid_fk")
	username := context.FormValue("username")
	DataProcessingTech := context.FormValue("data_processing_tech")
	DataTypes := context.FormValue("data_types")

	sessionModel := &jsonmodels.SessionRecord{}
	dataProcessingAnalysisModel := &jsonmodels.FormDataProcessingAnalysis{}

	err := GetUserRecord(uuid, context, sessionModel, username)
	if err != nil {
		logger.LogError(err.Error())
	}

	//set variables - no need to check as the separation of concerns is on the databases
	dataProcessingAnalysisModel.UuidFK = sessionModel.Uuid
	dataProcessingAnalysisModel.DataProcessingTech = DataProcessingTech
	dataProcessingAnalysisModel.DataTypes = DataTypes

	dbCreateErr := db_kafka_env_config.DMP_DB.Create(&dataProcessingAnalysisModel).Error
	if dbCreateErr != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "err_creating_data_proc_analysis"})
		logger.LogError("Error creating data processing analysis: " + err.Error())
		return err
	}
	return nil
}

func SubmitDataQaQc(context *fiber.Ctx) error {
	//uuid or username is accepted as it's handled hy the GetUserRecord function below
	uuid := context.FormValue("uuid_fk")
	username := context.FormValue("username")
	QaQcMethods := context.FormValue("qa_qc_methods")
	QaQcStrategies := context.FormValue("qa_qc_strategies")
	ValidationCalibrationMethods := context.FormValue("validation_calibration_methods")

	sessionModel := &jsonmodels.SessionRecord{}
	dataQaQcModel := &jsonmodels.FormDataQaQc{}

	err := GetUserRecord(uuid, context, sessionModel, username)
	if err != nil {
		logger.LogError(err.Error())
	}

	//set variables - no need to check as the separation of concerns is on the databases
	dataQaQcModel.UuidFK = sessionModel.Uuid
	dataQaQcModel.QaQcMethods = QaQcMethods
	dataQaQcModel.QaQcStrategies = QaQcStrategies
	dataQaQcModel.ValidationCalibrationMethods = ValidationCalibrationMethods

	dbCreateErr := db_kafka_env_config.DMP_DB.Create(&dataQaQcModel).Error
	if dbCreateErr != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "err_creating_qa_qc"})
		logger.LogError("Error creating QA QC: " + err.Error())
		return err
	}
	return nil
}
