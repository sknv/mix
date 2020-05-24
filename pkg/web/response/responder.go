package response

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"mix/pkg/log"
	"mix/pkg/response"
)

const (
	MsgRenderError = "render http error"
)

type webError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type webResponse struct {
	Data  interface{} `json:"data"`
	Error *webError   `json:"error"`
}

// RenderError renders JSON error response.
func RenderError(w http.ResponseWriter, r *http.Request, err error) {
	logger := log.ExtractLogger(r.Context()).With("error", err.Error()) // log the rendered error
	defer func() { logger.Warn(MsgRenderError) }()

	// Prepare the error to render
	webErr := webError{Message: err.Error()}

	var cause *response.Error
	if errors.As(err, &cause) {
		webErr.Code = cause.Code
		webErr.Message = cause.Message

		if cause.Internal != nil { // log the internal error if present
			logger = logger.With("internal", cause.Internal.Error())
		}
	}

	render.Status(r, http.StatusBadRequest) // default status for errors
	render.JSON(w, r, webResponse{Error: &webErr})
}

// RenderData renders JSON data reponse.
func RenderData(w http.ResponseWriter, r *http.Request, data interface{}) {
	render.JSON(w, r, webResponse{Data: data})
}
