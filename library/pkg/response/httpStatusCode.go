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
	CodeCannotCreateBorrowBook     = 400007
	CodeBookGenreExists            = 400008
	CodeBookAuthorExists           = 400009
	CodeCannotInsertGenreToBook    = 400010
	CodeCannotInsertAuthorToBook   = 400011
	CodeCannotRemoveGenreFromBook  = 400012
	CodeCannotRemoveAuthorFromBook = 400013
	CodeInvalidReturnBookDate      = 400014
	CodeCannotReturnBorrowBook     = 400015
	CodeBookNotFound               = 404001
	CodeAuthorNotFound             = 404002
	CodeGenreNotFound              = 404003
	CodeBorrowBookNotFound         = 404004
	CodeGenreNotFoundByBookId      = 404005
	CodeAuthorNotFoundByBookId     = 404006
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
	CodeCannotCreateBorrowBook:     {HttpCode: http.StatusBadRequest, Message: "Cannot create borrow book"},
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
	CodeGenreNotFoundByBookId:      {HttpCode: http.StatusNotFound, Message: "Cannot find genre by book id"},
	CodeAuthorNotFoundByBookId:     {HttpCode: http.StatusNotFound, Message: "Cannot find author by book id"},
	CodeInternalServerError:        {HttpCode: http.StatusInternalServerError, Message: "Internal Server Error"},
}
