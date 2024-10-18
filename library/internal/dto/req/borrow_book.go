package req

import "time"

type BorrowBookPostDto struct {
	StudentId         string                    `json:"student_id" binding:"required,min=1"`
	DueDate           time.Time                 `json:"due_date" binding:"required"`
	BorrowBookDetails []BorrowBookDetailPostDto `json:"borrow_book_details" binding:"required"`
}
type BorrowBookDetailPostDto struct {
	BookId   int32 `json:"book_id" binding:"required,gt=0"`
	Quantity int32 `json:"quantity" binding:"required,gt=0"`
}

type BorrowBookDetailPageDto struct {
	StudentID *string `json:"student_id" form:"student_id" binding:"omitempty,min=1"`
	DayRange  *int32  `json:"day_range" form:"day_range" binding:"omitempty,gt=0"`
	PageInfo
}
