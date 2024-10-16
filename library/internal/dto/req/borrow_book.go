package req

import "time"

type BorrowBookPostDto struct {
	StudentId         string                    `json:"student_id" binding:"required,min=1"`
	ReturnDate        time.Time                 `json:"return_date" binding:"required"`
	BorrowBookDetails []BorrowBookDetailPostDto `json:"borrow_book_details" binding:"required"`
}
type BorrowBookDetailPostDto struct {
	BookId   int32 `json:"book_id" binding:"required,gt=0"`
	Quantity int32 `json:"quantity" binding:"required,gt=0"`
}
