package transport

import (
	"context"
	"encoding/json"
	"log"
	stdHTTP "net/http"
	"strconv"

	"github.com/facily-tech/go-scaffold/pkg/domains/user"
	userErr "github.com/facily-tech/go-scaffold/pkg/domains/user/error"
	"github.com/facily-tech/go-scaffold/pkg/domains/user/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func NewHTTPHandler(svc user.ServiceI) stdHTTP.Handler {
	options := []http.ServerOption{
		http.ServerErrorEncoder(errorHandler),
	}

	postUser := http.NewServer(
		user.PostUser(svc),
		decodePostUser,
		codeHTTP{200}.encodeResponse,
		options...,
	)

	getUserByID := http.NewServer(
		user.GetUserByID(svc),
		decodeGetUserByID,
		codeHTTP{200}.encodeResponse,
		options...,
	)

	putUser := http.NewServer(
		user.UpdateUser(svc),
		decodePutUser,
		codeHTTP{200}.encodeResponse,
		options...,
	)

	deleteUser := http.NewServer(
		user.DeleteUser(svc),
		decodeDelete,
		codeHTTP{200}.encodeResponse,
		options...,
	)
	r := chi.NewRouter()

	r.Post("/", postUser.ServeHTTP)
	r.Get("/{id}", getUserByID.ServeHTTP)
	r.Put("/{id}", putUser.ServeHTTP)
	r.Delete("/{id}", deleteUser.ServeHTTP)
	return r
}

// decodifica o body para CreateUserRequest
func decodePostUser(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var userCreate model.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&userCreate); err != nil {
		return nil, userErr.ErrEmptyRepository
	}

	// O validate verifica se todos os campos obrigatorios declarados na struct vieram na request
	if err := validate.Struct(userCreate); err != nil {
		return nil, userErr.ErrInvalidBody
	}

	return userCreate, nil
}

// Pega o id na URL do endpoint e cria uma GetUserByIDRequest
func decodeGetUserByID(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var (
		userGet model.GetUserByIDRequest
		err     error
	)
	if userGet.ID, err = strconv.Atoi(chi.URLParam(r, "id")); err != nil {
		return nil, userErr.ErrInvalidPath
	}

	return userGet, err
}

// decodifica um UpdateUserRequest
func decodePutUser(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var (
		userUpdate model.UpdateUserRequest
		err        error
	)
	if err := json.NewDecoder(r.Body).Decode(&userUpdate); err != nil {
		return nil, userErr.ErrInvalidBody
	}

	if userUpdate.ID, err = strconv.Atoi(chi.URLParam(r, "id")); err != nil {
		return nil, userErr.ErrInvalidPath
	}
	if err = validate.Struct(userUpdate); err != nil {
		return nil, err
	}
	return userUpdate, nil
}

func decodeDelete(_ context.Context, r *stdHTTP.Request) (interface{}, error) {
	var (
		userDelete model.DeleteUserByIDRequest
		err        error
	)
	if userDelete.ID, err = strconv.Atoi(chi.URLParam(r, "id")); err != nil {
		return nil, err
	}

	return userDelete, nil
}

// Struct que vai definir o status HTTP de reposta se tudo der certo
type codeHTTP struct {
	int
}

// O codeHTTP implementa uma função de resposta da request
func (c codeHTTP) encodeResponse(_ context.Context, w stdHTTP.ResponseWriter, input interface{}) error {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(c.int)
	return json.NewEncoder(w).Encode(input)
}

// O erroHandler é chamado caso algo dê errado e ele é chamado no options na função de NewHTTPHandler
func errorHandler(_ context.Context, err error, w stdHTTP.ResponseWriter) {
	resp, code := userErr.RESTErrorBussines.ErrorProcess(err)

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": resp}); err != nil {
		log.Printf("Encoding error, nothing much we can do: %v", err)
	}
}
