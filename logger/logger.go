// package to simplify logging to a file and to differentiate logging types, e.g. error, warning, etc.
// This also provides a way of a sort of stack trace with the logs, where the caller is logged as well
package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var verbose = false
var ex, _ = os.Executable()
var logdir = filepath.Dir(ex) + "/logs/"

var initDir = false

// set verbosity, this is just if we want to output to the console as well as a file
func SetVerbosity(verbosity bool) {
	verbose = verbosity
}

// change the log directory
func ChangeDir(newDir string) {
	logdir = newDir
}

// get caller for stack trace
func retrieveCallInfo() (packageName string, funcName string, line int, fileName string) {
	pc, file, line, _ := runtime.Caller(3)
	_, fileName = path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)

	funcName = parts[pl-1]

	if len(parts[pl-2]) > 0 {
		if parts[pl-2][0] == '(' {
			funcName = parts[pl-2] + "." + funcName
			packageName = strings.Join(parts[0:pl-2], ".")
		} else {
			packageName = strings.Join(parts[0:pl-1], ".")
		}
	} else {
		packageName = parts[1]
	}

	return packageName, funcName, line, fileName
}

// prefox log level to the message
func LogError(message string) {
	logToFile("ERROR:", message)
}

func LogMessage(message string) {
	logToFile("INFO:", message)
}

func LogWarning(message string) {
	logToFile("WARN:", message)
}

// set up logging directories
func init() {
	if err := os.Mkdir(logdir, os.ModePerm); err != nil && !errors.Is(err, os.ErrExist) {
		fmt.Println(err.Error())
	} else {
		if !initDir {
			fmt.Println(fmt.Sprintf("All logs can be found in: %v\n", logdir))
			initDir = true
		}
	}
}

// copy logs to a file
func logToFile(logLevel string, message string) {
	LOG_FILE := fmt.Sprintf("%v%v%v%v.log", logdir, time.Now().Local().Day(), time.Now().Local().Month(), time.Now().Local().Year())
	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Cannot print to log.")
	}
	defer logFile.Close()

	// Set log out put
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	// log.SetFlags(log.Lshortfile | log.LstdFlags)
	//get the caller and package for logging
	pkgName, _, _, _ := retrieveCallInfo()
	log.Println(fmt.Sprintf("%v %v in package %v: %v ", time.Now().Local(), logLevel, pkgName, message))
	//if verbose mode on, also output to the console
	if verbose {
		fmt.Println(fmt.Sprintf("%v %v in package %v: %v ", time.Now().Local(), logLevel, pkgName, message))
	}
}
