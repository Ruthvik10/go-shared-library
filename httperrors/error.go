package httperrors

import (
	"net/http"

	"github.com/Ruthvik10/go-shared-library/jsonutil"
	"github.com/Ruthvik10/go-shared-library/logger"
)

type envelope map[string]any

func logError(r *http.Request, err error) {
	logger.Error(err, map[string]any{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	err := jsonutil.WriteJSON(w, envelope{"error": message}, status, nil)
	if err != nil {
		logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)
	message := "the server encountered a problem and could not process your request"
	errorResponse(w, r, http.StatusInternalServerError, message)
}

func NotFoundErrorResponse(w http.ResponseWriter, r *http.Request) {
	message := "could not find the requested resource"
	errorResponse(w, r, http.StatusNotFound, message)
}

func BadRequestErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusNotFound, err.Error())
}
