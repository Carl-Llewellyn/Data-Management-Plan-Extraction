// routes for kafka submissions
package kafka_routes

import (
	apiconfig "dmp-api/api/api_config"
	jsonmodels "dmp-api/api/json_models"
	postgresroutes "dmp-api/api/routes/postgres"
	"dmp-api/logger"
	"encoding/json"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

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

	err := postgresroutes.GetUserRecord(uuid, context, sessionModel, username)
	if err != nil {
		logger.LogError(err.Error())
	}

	//set variables - no need to check as the separation of concerns is on the databases
	dataCollectionModel.UuidFK = sessionModel.Uuid
	dataCollectionModel.PowerPointFileName = powerpointFileName
	dataCollectionModel.IsNewData = isNewData
	dataCollectionModel.DatasetName = datasetName
	dataCollectionModel.DataStorageRequirements = dataStorageRequirements

	// Create a Kafka producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": apiconfig.GetFullKafkaAddress()})
	if err != nil {
		logger.LogError("Failed to create producer: " + err.Error())
		return err
	}

	// Marshal the dataCollectionModel to JSON
	data, err := json.Marshal(dataCollectionModel)
	if err != nil {
		logger.LogError("Failed to marshal data: " + err.Error())
		return err
	}

	// Produce a message to the Kafka topic
	topic := "data_collection_topic"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	logger.LogMessage("Successfully produced message to Kafka topic")

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

	err := postgresroutes.GetUserRecord(uuid, context, sessionModel, username)
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

	// Create a Kafka producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": apiconfig.GetFullKafkaAddress()})
	if err != nil {
		logger.LogError("Failed to create producer: " + err.Error())
		return err
	}

	// Marshal the dataPlanningModel to JSON
	data, err := json.Marshal(dataCollectionModel)
	if err != nil {
		logger.LogError("Failed to marshal data: " + err.Error())
		return err
	}

	// Produce a message to the Kafka topic
	topic := "data_planning_topic"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	logger.LogMessage("Successfully produced message to Kafka topic")

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

	err := postgresroutes.GetUserRecord(uuid, context, sessionModel, username)
	if err != nil {
		logger.LogError(err.Error())
	}

	//set variables - no need to check as the separation of concerns is on the databases
	dataProcessingAnalysisModel.UuidFK = sessionModel.Uuid
	dataProcessingAnalysisModel.DataProcessingTech = DataProcessingTech
	dataProcessingAnalysisModel.DataTypes = DataTypes

	// Create a Kafka producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": apiconfig.GetFullKafkaAddress()})
	if err != nil {
		logger.LogError("Failed to create producer: " + err.Error())
		return err
	}

	// Marshal the dataProcessingAnalysisModel to JSON
	data, err := json.Marshal(dataProcessingAnalysisModel)
	if err != nil {
		logger.LogError("Failed to marshal data: " + err.Error())
		return err
	}

	// Produce a message to the Kafka topic
	topic := "data_processing_analysis_topic"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	logger.LogMessage("Successfully produced message to Kafka topic")

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

	err := postgresroutes.GetUserRecord(uuid, context, sessionModel, username)
	if err != nil {
		logger.LogError(err.Error())
	}

	//set variables - no need to check as the separation of concerns is on the databases
	dataQaQcModel.UuidFK = sessionModel.Uuid
	dataQaQcModel.QaQcMethods = QaQcMethods
	dataQaQcModel.QaQcStrategies = QaQcStrategies
	dataQaQcModel.ValidationCalibrationMethods = ValidationCalibrationMethods

	// Create a Kafka producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": apiconfig.GetFullKafkaAddress()})
	if err != nil {
		logger.LogError("Failed to create producer: " + err.Error())
		return err
	}

	// Marshal the dataQaQcModel to JSON
	data, err := json.Marshal(dataQaQcModel)
	if err != nil {
		logger.LogError("Failed to marshal data: " + err.Error())
		return err
	}

	// Produce a message to the Kafka topic
	topic := "data_qa_qc_topic"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	logger.LogMessage("Successfully produced message to Kafka topic")

	return nil
}
