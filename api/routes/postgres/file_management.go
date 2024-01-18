// set up postgres specific routes for file management
package postgres_routes

import (
	apiconfig "dmp-api/api/api_config"
	jsonmodels "dmp-api/api/json_models"
	"dmp-api/logger"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

var uploadRootDir = "/tmp/uploads/"

func DownloadFile(context *fiber.Ctx) error {
	file := context.Query("file")
	logger.LogMessage("Getting file: " + file)

	// Get the file path
	filePath := filepath.Join(uploadRootDir, file)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		logger.LogError(err.Error())
		return errors.New("err_file_not_found")
	}

	// Open the file
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.LogError(err.Error())
		return errors.New("err_reading_file")
	}

	// Set the headers
	context.Set("Content-Type", "application/octet-stream")
	context.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, file))

	// Send the file data
	err = context.Send(fileData)
	if err != nil {
		logger.LogError(err.Error())
		return errors.New("err_sending_file")
	}

	logger.LogMessage("Download served.")
	return nil
}

func generateUserDir(context *fiber.Ctx, fullPath bool) (error, string) {
	uuid := context.FormValue("uuid")
	username := context.FormValue("username")
	sessionModel := jsonmodels.SessionRecord{}

	err := GetUserRecord(uuid, context, &sessionModel, username)
	if err != nil {
		logger.LogError("err_unable_to_get_usr_record_dir_" + err.Error())
		return err, ""
	}

	//return full path or just uuid path
	if fullPath {
		logger.LogMessage("succ_file_list")
		return nil, uploadRootDir + sessionModel.Uuid.String() + "/"
	} else {
		return nil, sessionModel.Uuid.String() + "/"
	}
}

func RetrieveFiles(context *fiber.Ctx) error {
	var files []jsonmodels.File
	_, dir := generateUserDir(context, true)
	uuid := context.Query("uuid")
	username := context.Query("username")
	sessionModel := jsonmodels.SessionRecord{}
	userRecordError := GetUserRecord(uuid, context, &sessionModel, username)

	if userRecordError != nil {
		logger.LogError("User record retrieval error for files: " + userRecordError.Error())
		return userRecordError
	}

	fmt.Println(dir)
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			logger.LogError("Cannot walk dir: " + err.Error())
			return err
		}
		if !d.IsDir() {
			relativePath, err := filepath.Rel(dir, path)
			if err != nil {
				logger.LogError("Cannot read dir: " + err.Error())
				return err
			}
			files = append(files, jsonmodels.File{
				Name: d.Name(),
				Path: apiconfig.GetFullAPIAddress() + "/api/postgres/download_file?file=" + sessionModel.Uuid.String() + "/" + url.QueryEscape(relativePath)}) //set the download link for each file
		}
		return nil
	})

	if err != nil {
		logger.LogError("Error retrieving files: " + err.Error())
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.LogMessage("succ_files_retrieved")
	return context.JSON(files)
}

func UploadFile(context *fiber.Ctx) error {
	err, uploadDir := generateUserDir(context, true)
	if err != nil {
		logger.LogError(err.Error())
		return err
	}

	file, err := context.FormFile("file")
	if err != nil {
		logger.LogError(err.Error())
		return err
	}

	src, err := file.Open()
	if err != nil {
		logger.LogError(err.Error())
		return err
	}
	defer src.Close()

	err1 := os.MkdirAll(uploadDir, os.ModePerm)
	if err1 != nil {
		logger.LogError(err.Error())
		return err
	}

	dst, err := os.Create(uploadDir + file.Filename)
	if err != nil {
		logger.LogError(err.Error())
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		logger.LogError(err.Error())
		return err
	}

	logger.LogMessage("file_upload_succ_" + uploadDir)
	return nil
}
