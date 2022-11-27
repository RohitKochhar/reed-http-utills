package reed

import "net/http"

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
