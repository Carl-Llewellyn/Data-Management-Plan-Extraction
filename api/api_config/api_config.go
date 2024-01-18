// package to manage and enforce validation on the API and Kafka address and port
package apiconfig

import (
	"errors"
	"net"
	"strconv"
)

// default if not initialized
var apiPort = "8080"
var apiIP = "0.0.0.0"

var kafkaBrokerPort = "9092"
var kafaBrokerIP = "0.0.0.0"

// getters and setters to enforce validation of addresses and ports
func SetAPIPort(newPort string) error {
	err := validatePort(newPort)
	if err == nil {
		apiPort = newPort
		return nil
	}
	return err
}

func SetAPIIP(newAddress string) error {
	err := validateIP(newAddress)
	if err == nil {
		apiIP = newAddress
		return nil
	}
	return err
}

// return the full address - allows getting without exposing the address and port individually
func GetFullAPIAddress() string {
	return apiIP + ":" + apiPort
}

// getters and setters to enforce validation of addresses and ports
func SetKafkaBrokerPort(newPort string) error {
	err := validatePort(newPort)
	if err == nil {
		kafkaBrokerPort = newPort
		return nil
	}
	return err
}

func SetKafkaBrokerIP(newAddress string) error {
	err := validateIP(newAddress)
	if err == nil {
		kafaBrokerIP = newAddress
		return nil
	}
	return err
}

// return the full address - allows getting without exposing the address and port individually
func GetFullKafkaAddress() string {
	return apiIP + ":" + apiPort
}

// validation functions
func validatePort(newPort string) error {
	_, err := strconv.Atoi(newPort)
	if err != nil {
		return err
	}
	apiPort = newPort
	return nil
}

func validateIP(newIP string) error {
	if net.ParseIP(newIP) == nil {
		return errors.New("Error invalid API address.")
	}
	apiIP = newIP
	return nil
}
