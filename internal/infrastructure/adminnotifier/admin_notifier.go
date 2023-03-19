package adminnotifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const notificationPath = "api/notify"

type AdminNotifier struct {
	httpClient           *http.Client
	notificationEndpoint string
}

func New(httpClient *http.Client, server string) *AdminNotifier {
	notificationEndpoint, err := url.JoinPath(server, notificationPath)
	if err != nil {
		log.Fatalln("Invalid server set for admin notifications:", err)
	}
	log.Println("Will send admin notifications to endpoint", notificationEndpoint)
	return &AdminNotifier{httpClient: httpClient, notificationEndpoint: notificationEndpoint}
}

func (n *AdminNotifier) Notify(level string, employeeAbbreviation string, message string) error {

	requestBody, err := json.Marshal(map[string]string{"level": level, "employeeAbbreviation": employeeAbbreviation, "message": message})
	if err != nil {
		return fmt.Errorf("could not serialize request body: %w", err)
	}

	response, err := n.httpClient.Post(n.notificationEndpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("could send request: %w", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("could parse response body: %w", err)
	}
	if response.StatusCode > 299 || response.StatusCode < 200 {
		return fmt.Errorf("request failed: %s, %s", response.Status, string(responseBody))
	}

	// no need to parse response if request was successful
	return nil
}
