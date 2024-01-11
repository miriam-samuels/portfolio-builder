package helper

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
)

// This function takes in an interface for the request and returns a function which takes in a handler function and returns a handler
func ParseRequestBody(w http.ResponseWriter, r *http.Request, i interface{}) error {
	err := json.NewDecoder(r.Body).Decode(i)
	if err != nil {
		if err.Error() != "EOF" {
			return err
		}
		return err
	}
	return nil
}

func ParseMultipartRequestBody(w http.ResponseWriter, r *http.Request) (multipart.File, error) {
	//  PARSE formdata including uploaded file
	err := r.ParseMultipartForm(10 << 20) // 10mb limit
	if err != nil {
		SendResponse(w, http.StatusBadRequest, false, "error parsing body:", nil)
		return nil, err
	}

	f, _, err := r.FormFile("file")
	if err != nil {
		SendResponse(w, http.StatusBadRequest, false, "Unable to retrieve file from form", nil)
		return nil, err
	}
	return f, err
}

