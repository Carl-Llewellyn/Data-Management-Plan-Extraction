package postgres_routes

import (
	"dmp-api/logger"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"dmp-api/api/db_kafka_env_config"
	jsonmodels "dmp-api/api/json_models"

	"github.com/gofiber/fiber/v2"
	GUUID "github.com/google/uuid"
)

func GetUserRecordByUsername(context *fiber.Ctx) error {
	username := context.Query("username")
	sessionModel := &jsonmodels.SessionRecord{}
	if username == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "err_username_empty",
		})
		err := errors.New("username cannot be empty")
		logger.LogError(err.Error())
		return err
	}

	err := db_kafka_env_config.DMP_DB.Where("username = ?", username).First(sessionModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "err_user_not_found"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "succ_user_found",
			"data": sessionModel})

	return nil
}

func CreateNewUsers(context *fiber.Ctx) error {
	var newUsers jsonmodels.SessionRecords

	err := json.Unmarshal(context.Body(), &newUsers)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "err_parsing_user"})
		logger.LogError("Error parsing new user: " + err.Error())
		return err
	}

	for _, user := range newUsers.Data {
		blankUUID, _ := GUUID.Parse("00000000-0000-0000-0000-000000000000")
		if user.Uuid == blankUUID {
			user.Uuid = GUUID.New()
		}

		err := db_kafka_env_config.DMP_DB.Create(&user).Error
		if err != nil {
			context.Status(http.StatusUnprocessableEntity).JSON(
				&fiber.Map{"message": "err_creating_user"})
			logger.LogError("Error creating new user: " + err.Error())
			return err
		}
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "New user(s) created"})

	return nil
}

// tihs function is used by the others to get the user record
func GetUserRecord(uuid string, context *fiber.Ctx, sessionModel *jsonmodels.SessionRecord, username string) error {
	uuid = strings.ReplaceAll(uuid, " ", "")
	username = strings.ReplaceAll(username, " ", "")
	blankUUID, _ := GUUID.Parse("00000000-0000-0000-0000-000000000000")

	if uuid == "" {
		uuid = blankUUID.String()
	}

	uuidParsed, err := GUUID.Parse(uuid)
	if err != nil {
		logger.LogError("UUID is invalid.")
		return errors.New("err_invalid_uuid")
	}

	//if none of the parameters are found
	if username == "" && uuidParsed == blankUUID {
		logger.LogError("Cannot find username or UUID in query.")
		return errors.New("err_uuid_username_not_in_query")
	}

	//if search wants to use both username and uuid
	if username != "" && uuidParsed != blankUUID {
		err := db_kafka_env_config.DMP_DB.Where("username = ?", username).Where("uuid = ?", uuidParsed).First(sessionModel).Error
		if err != nil {
			context.Status(http.StatusBadRequest).JSON(
				&fiber.Map{"message": "err_user_not_found"})
			logger.LogError(err.Error())
			return err
		}

		context.Status(http.StatusOK).JSON(
			&fiber.Map{"message": "succ_user_found",
				"data": sessionModel})

		return nil

	}

	if username != "" {
		err := db_kafka_env_config.DMP_DB.Where("username = ?", username).First(sessionModel).Error
		if err != nil {
			context.Status(http.StatusBadRequest).JSON(
				&fiber.Map{"message": "err_user_not_found"})
			logger.LogError(err.Error())
			return err
		}

		context.Status(http.StatusOK).JSON(
			&fiber.Map{"message": "succ_user_found",
				"data": sessionModel})

		return nil
	}

	if uuidParsed != blankUUID {
		err := db_kafka_env_config.DMP_DB.Where("uuid = ?", uuidParsed).First(sessionModel).Error
		if err != nil {
			context.Status(http.StatusBadRequest).JSON(
				&fiber.Map{"message": "err_user_not_found"})
			logger.LogError(err.Error())
			return err
		}

		context.Status(http.StatusOK).JSON(
			&fiber.Map{"message": "succ_user_found",
				"data": sessionModel})

		return nil
	}

	return nil
}
