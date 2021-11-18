package utils

import "errors"

// AppError defines general error's throughout the system
type AppError error

var (
	ErrInvalidCredentials AppError = errors.New("invalid credentials")
	ErrServerError        AppError = errors.New("server_error")
	ErrInvalidRequest     AppError = errors.New("invalid_request")
	ErrNotAllowed         AppError = errors.New(
		"not allowed to access this server")
	ErrStrictPasswordPolicyViolation AppError = errors.New("password_policy_violation")
	ErrUnauthorized                  AppError = errors.New("unauthorized")
	ErrUserExists                    AppError = errors.New("user_exists")
	ErrUserNotFound                  AppError = errors.New("user does not exists")
	ErrWrongPassword                 AppError = errors.New("password doesn't match")
	ErrInvalidToken                  AppError = errors.New("invalid_credential")
)

// ErrorStatusCodes holds the http status codes for every AppError
var ErrorStatusCodes = map[AppError]int{
	ErrInvalidRequest:                400,
	ErrInvalidCredentials:            401,
	ErrServerError:                   500,
	ErrUnauthorized:                  401,
	ErrUserExists:                    401,
	ErrStrictPasswordPolicyViolation: 401,
	ErrUserNotFound:                  400,
	ErrWrongPassword:                 400,
	ErrNotAllowed:                    401,
}

// ErrorDescriptions holds detailed error description for every AppError
var ErrorDescriptions = map[AppError]string{
	ErrServerError:                   "The authorization server encountered an unexpected condition that prevented it from fulfilling the request",
	ErrInvalidCredentials:            "Invalid Credentials",
	ErrInvalidRequest:                "The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed",
	ErrUnauthorized:                  "The user does not have requested authorization to access this resource",
	ErrUserExists:                    "This email is already assigned to another user",
	ErrStrictPasswordPolicyViolation: "Please ensure the password is 8 characters long and has 1 digit, 1 lowercase alphabet, 1 uppercase alphabet and 1 special character",
	ErrInvalidToken:                  "This credential is not a valid credential",
	ErrNotAllowed: "You are not allowed to access this server, " +
		"contact server admin",
}
