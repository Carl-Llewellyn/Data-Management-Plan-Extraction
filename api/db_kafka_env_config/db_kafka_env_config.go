// This pacage configures the database connection through a .env file
package db_kafka_env_config

import (
	apiconfig "dmp-api/api/api_config"
	"dmp-api/logger"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DMP_DB *gorm.DB

type Config struct {
	Host        string
	Port        string
	DB_Password string
	DB_User     string
	DB_Name     string
	SSLMode     string
	KAFA_IP     string
	KAFKA_PORT  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.DB_User, config.DB_Password, config.DB_Name, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

func setupKafka(config *Config) error {
	if ipError := apiconfig.SetKafkaBrokerIP(config.KAFA_IP); ipError != nil {
		logger.LogError("Kafka IP setup failed: " + ipError.Error())
		return ipError
	}

	if portError := apiconfig.SetKafkaBrokerIP(config.KAFA_IP); portError != nil {
		logger.LogError("Kafka port setup failed: " + portError.Error())
		return portError
	}

	return nil
}

func init() {
	//setup vars
	err := godotenv.Load(".env")
	if err != nil {
		logger.LogError(err.Error())
	}

	config := &Config{
		Host:        os.Getenv("DB_HOST"),
		Port:        os.Getenv("DB_PORT"),
		DB_Password: os.Getenv("DB_PASS"),
		DB_User:     os.Getenv("DB_USER"),
		SSLMode:     os.Getenv("DB_SSLMODE"),
		DB_Name:     os.Getenv("DB_NAME"),
		KAFA_IP:     os.Getenv("KAFKA_IP"),
		KAFKA_PORT:  os.Getenv("KAFKA_PORT"),
	}

	//set up Kafka
	kafkaErr := setupKafka(config)
	if kafkaErr != nil {
		logger.LogError(kafkaErr.Error())
	}

	//connect
	tmpDB, err := NewConnection(config)
	if err != nil {
		logger.LogError(err.Error())
	}
	DMP_DB = tmpDB
}
