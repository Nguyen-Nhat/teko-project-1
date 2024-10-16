package response

import "net/http"

const (
	CodeSuccess                    = 200000
	CodeCreated                    = 201000
	CodeInvalidRequestParam        = 400001
	CodeInvalidRequestBody         = 400002
	CodeInvalidPathVariable        = 400003
	CodeCannotCreateBook           = 400004
	CodeCannotCreateAuthor         = 400005
	CodeCannotCreateGenre          = 400006
	CodeBookGenreExists            = 400007
	CodeBookAuthorExists           = 400008
	CodeCannotInsertGenreToBook    = 400009
	CodeCannotInsertAuthorToBook   = 400010
	CodeCannotRemoveGenreFromBook  = 400011
	CodeCannotRemoveAuthorFromBook = 400012
	CodeInvalidReturnBookDate      = 400013
	CodeCannotReturnBorrowBook     = 400014
	CodeBookNotFound               = 404001
	CodeAuthorNotFound             = 404002
	CodeGenreNotFound              = 404003
	CodeBorrowBookNotFound         = 404004
	CodeInternalServerError        = 500000
)

type responseMapping struct {
	HttpCode int
	Message  string
}

var msg = map[int]responseMapping{
	CodeSuccess:                    {HttpCode: http.StatusOK, Message: "Success"},
	CodeCreated:                    {HttpCode: http.StatusCreated, Message: "Successfully created"},
	CodeInvalidRequestParam:        {HttpCode: http.StatusBadRequest, Message: "Invalid Request Param"},
	CodeInvalidRequestBody:         {HttpCode: http.StatusBadRequest, Message: "Invalid Request Body"},
	CodeInvalidPathVariable:        {HttpCode: http.StatusBadRequest, Message: "Invalid Path Variable"},
	CodeCannotCreateBook:           {HttpCode: http.StatusBadRequest, Message: "Cannot create book"},
	CodeCannotCreateAuthor:         {HttpCode: http.StatusBadRequest, Message: "Cannot create author"},
	CodeCannotCreateGenre:          {HttpCode: http.StatusBadRequest, Message: "Cannot create genre"},
	CodeBookGenreExists:            {HttpCode: http.StatusBadRequest, Message: "The book has this genre"},
	CodeBookAuthorExists:           {HttpCode: http.StatusBadRequest, Message: "The book has this author"},
	CodeCannotInsertGenreToBook:    {HttpCode: http.StatusBadRequest, Message: "Cannot insert genre to book"},
	CodeCannotInsertAuthorToBook:   {HttpCode: http.StatusBadRequest, Message: "Cannot insert author to book"},
	CodeCannotRemoveAuthorFromBook: {HttpCode: http.StatusBadRequest, Message: "Cannot remove author from book"},
	CodeCannotRemoveGenreFromBook:  {HttpCode: http.StatusBadRequest, Message: "Cannot remove genre from book"},
	CodeInvalidReturnBookDate:      {HttpCode: http.StatusBadRequest, Message: "Invalid return book date"},
	CodeBookNotFound:               {HttpCode: http.StatusNotFound, Message: "Cannot find book"},
	CodeAuthorNotFound:             {HttpCode: http.StatusNotFound, Message: "Cannot find author"},
	CodeGenreNotFound:              {HttpCode: http.StatusNotFound, Message: "Cannot find genre"},
	CodeCannotReturnBorrowBook:     {HttpCode: http.StatusBadRequest, Message: "Cannot return borrow book"},
	CodeBorrowBookNotFound:         {HttpCode: http.StatusNotFound, Message: "Cannot return borrow book"},
	CodeInternalServerError:        {HttpCode: http.StatusInternalServerError, Message: "Internal Server Error"},
}
