package reed

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

/*
Name: ReplyTextContent

Description:
  - Wraps a string in a text/plain response and replies back to
    the request sender.

Inputs:
  - w (http.ResponseWriter): Used to craft the HTTP response
  - r (*http.Request): Used to access information about the initial request
  - status (int): Specifies the HTTP response code
  - content (string): Text content to be included in the response body

Outputs:
  - None
*/
func ReplyTextContent(w http.ResponseWriter, r *http.Request, status int, content string) {
	w.Header().Set("Content-Type", "text/plain") // Set the header to plain/text
	w.WriteHeader(status)                        // Set the request status
	w.Write([]byte(content + "\n"))              // Attach the message to the body
}

/*
Name: ReplyError

Description:
  - Logs error message information to standard logger
  - Sends error message back to request sender

Inputs:
  - w (http.ResponseWriter): Used to craft the HTTP response
  - r (*http.Request): Used to access information about the initial request
  - status (int): Specifies the HTTP response code
  - errMessage (string): Text content to be included in the response body

Outputs:
  - None
*/
func ReplyError(w http.ResponseWriter, r *http.Request, status int, errMessage string) {
	// Log the error to the standard logger
	log.Printf("%s %s: Error %d %s", r.URL, r.Method, status, errMessage)
	// Reply to the request with an error
	http.Error(w, errMessage, status)
}

/*
Name: PutString

Description:
  - Wraps HTTP PUT request in a utility function
    for easier usage and clarity

Inputs:
  - httpUrl (string): URL to send request to
  - message (string): Message to send
  - expCodes ([]int): Expected response from server
    (nil if any response should be expected accepted)

Outputs:
  - error: Returned if there was an error in the process
*/
func PutString(httpUrl string, message string, expCodes []int) error {
	// Create the HTTP request
	req, err := http.NewRequest(
		http.MethodPut,
		httpUrl,
		bytes.NewBuffer([]byte(message)),
	)
	if err != nil {
		return fmt.Errorf("error while creating PUT request: %q", err)
	}
	req.Header.Set("Content-Type", "text/plain")
	// Send the request and get the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error while sending PUT request: %q", err)
	}
	// Check the reponse code if applicable
	if expCodes != nil {
		// Iterate through the provided expCodes
		for _, code := range expCodes {
			if resp.StatusCode == code {
				// If the response code was in our acceptable list
				// we can return a success
				return nil
			}
		}
		// If we didn't find an acceptable code, return an error
		return fmt.Errorf("response contained unacceptable status code: %s (acceptable options are: %+v)", resp.Status, expCodes)
	}
	return nil
}

/*
Name: DeleteString

Description:
  - Wraps HTTP DELETE request in utility function
    for easier usage and clarity

Inputs:
  - httpUrl (string): URL to send request to
  - message (string): Message to send
  - expCodes ([]int): Expected response from server
    (nil if any response should be expected accepted)

Outputs:
  - error: Returned if there was an error in the process
*/
func DeleteString(httpUrl string, message string, expCodes []int) error {
	req, err := http.NewRequest(
		http.MethodDelete,
		httpUrl,
		bytes.NewBuffer([]byte(message)),
	)
	if err != nil {
		return fmt.Errorf("error while creating DELETE request: %q", err)
	}
	req.Header.Set("Content-Type", "text/plain")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error while sending DELETE request: %q", err)
	}
	// Check the reponse code if applicable
	if expCodes != nil {
		// Iterate through the provided expCodes
		for _, code := range expCodes {
			if resp.StatusCode == code {
				// If the response code was in our acceptable list
				// we can return a success
				return nil
			}
		}
		// If we didn't find an acceptable code, return an error
		return fmt.Errorf("response contained unacceptable status code: %s (acceptable options are: %+v)", resp.Status, expCodes)
	}
	return nil
}
