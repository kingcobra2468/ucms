package message

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrBadRequest = errors.New("unable to process request")
)

// Error handling.
type errorer interface {
	error() error
}

func MakeHTTPHandler(ms MessageService) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").Path("/message").Handler(httptransport.NewServer(
		makeNewNotificationEndpoint(ms),
		decodeNewNotificationRequest,
		encodeResponse,
	))

	return r
}

// Process the token registration request.
func decodeNewNotificationRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req NewNotificationRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, ErrBadRequest
	}
	return req, nil
}

// Handle the encoding of response data post endpoint logic.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// Handle for situations if an error exists.
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

// Process error codes for responding with the correct status code.
func codeFrom(err error) int {
	switch err {
	case ErrBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
