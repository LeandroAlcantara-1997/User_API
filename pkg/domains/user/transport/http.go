package transport

import (
	"context"
	"encoding/json"
	"log"
	stdHTTP "net/http"

	"github.com/facily-tech/go-scaffold/pkg/domains/user"
	userErr "github.com/facily-tech/go-scaffold/pkg/domains/user/error"
	"github.com/facily-tech/go-scaffold/pkg/domains/user/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(svc user.ServiceI) stdHTTP.Handler {
	options := []http.ServerOption{
		http.ServerErrorEncoder(errorHandler),
	}

	postUser := http.NewServer(
		user.PostUser(svc),
		decodeCreateUser,
		codeHTTP{200}.encodeResponse,
		options...,
	)
	r := chi.NewRouter()

	r.Post("/create", postUser.ServeHTTP)
	return r
}

func decodeCreateUser(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var userCreate model.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&userCreate); err != nil {
		return nil, userErr.ErrEmptyRepository
	}
	return userCreate, nil
}

type codeHTTP struct {
	int
}

func (c codeHTTP) encodeResponse(_ context.Context, w stdHTTP.ResponseWriter, input interface{}) error {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(c.int)
	return json.NewEncoder(w).Encode(input)
}

func errorHandler(_ context.Context, err error, w stdHTTP.ResponseWriter) {
	resp, code := userErr.RESTErrorBussines.ErrorProcess(err)

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": resp}); err != nil {
		log.Printf("Encoding error, nothing much we can do: %v", err)
	}
}
