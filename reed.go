package reed

import (
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
