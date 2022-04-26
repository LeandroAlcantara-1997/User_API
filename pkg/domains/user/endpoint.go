package user

import (
	"context"

	userErr "github.com/facily-tech/go-scaffold/pkg/domains/user/error"
	"github.com/facily-tech/go-scaffold/pkg/domains/user/model"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

func PostUser(svc ServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(model.CreateUserRequest)
		if !ok {
			return nil, errors.Wrap(userErr.ErrTypeAssertion, "cannot convert request-> CreateUserRequest")
		}

		resp, err := svc.CreateUser(ctx, req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func GetUserByID(svc ServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(model.GetUserByIDRequest)
		if !ok {
			return nil, errors.Wrap(userErr.ErrTypeAssertion, "cannot convert request-> GetUserByIDRequest")
		}

		resp, err := svc.FindUserByID(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func UpdateUser(svc ServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(model.UpdateUserRequest)
		if !ok {
			return nil, errors.Wrap(userErr.ErrTypeAssertion, "cannot convert request-> UpdateUserRequest")
		}

		resp, err := svc.UpdateUser(ctx, req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func DeleteUser(svc ServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(model.DeleteUserByIDRequest)
		if !ok {
			return nil, errors.Wrap(userErr.ErrTypeAssertion, "cannot convert request-> UpdateUserRequest")
		}

		err := svc.DeleteUserByID(ctx, req)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}
