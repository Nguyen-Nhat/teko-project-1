package response

import "net/http"

const (
	CodeSuccess             = 200000
	CodeCreated             = 201000
	CodeInvalidRequestParam = 400001
	CodeInvalidRequestBody  = 400002
	CodeInvalidPathVariable = 400003
	CodeUniversityNotFound  = 404001
	CodeStudentNotFound     = 404002
	CodeInternalServerError = 500000
)

type responseMapping struct {
	HttpCode int
	Message  string
}

var msg = map[int]responseMapping{
	CodeSuccess:             {HttpCode: http.StatusOK, Message: "Success"},
	CodeCreated:             {HttpCode: http.StatusCreated, Message: "Successfully created"},
	CodeInvalidRequestParam: {HttpCode: http.StatusBadRequest, Message: "Invalid Request Param"},
	CodeInvalidRequestBody:  {HttpCode: http.StatusBadRequest, Message: "Invalid Request Body"},
	CodeInvalidPathVariable: {HttpCode: http.StatusBadRequest, Message: "Invalid Path Variable"},
	CodeUniversityNotFound:  {HttpCode: http.StatusNotFound, Message: "University not found"},
	CodeStudentNotFound:     {HttpCode: http.StatusNotFound, Message: "Student not found"},
	CodeInternalServerError: {HttpCode: http.StatusInternalServerError, Message: "Internal Server Error"},
}
