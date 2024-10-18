package res

import (
	"library/internal/database"
	"library/internal/util"
	"time"
)

type BorrowBookDto struct {
	Id         int32      `json:"id"`
	StudentID  string     `json:"student_id"`
	BorrowDate time.Time  `json:"borrow_date"`
	DueDate    time.Time  `json:"due_date"`
	ReturnDate *time.Time `json:"return_date"`
	IsReturn   bool       `json:"is_return"`
}

func (bb *BorrowBookDto) FromModel(borrowBook database.BorrowBook) {
	bb.Id = borrowBook.ID
	bb.StudentID = borrowBook.StudentID
	bb.BorrowDate = borrowBook.BorrowDate
	bb.DueDate = borrowBook.DueDate
	bb.IsReturn = borrowBook.IsReturn
	bb.ReturnDate = util.FromNullTime(borrowBook.ReturnDate)
}

type BorrowBookItemDto struct {
	Id           int32 `json:"id"`
	BorrowBookID int32 `json:"borrow_book_id"`
	BookID       int32 `json:"book_id"`
	Quantity     int32 `json:"quantity"`
}

type BorrowItemWithBookInfoDto struct {
	BorrowBookItemDto
	Title      string    `json:"title"`
	StudentId  string    `json:"student_id"`
	BorrowDate time.Time `json:"borrow_date"`
	DueDate    time.Time `json:"due_date"`
}

func (b *BorrowBookItemDto) FromModel(detail database.BorrowBookDetail) {
	b.Id = detail.ID
	b.BookID = detail.BookID
	b.Quantity = detail.Quantity
	b.BorrowBookID = detail.BorrowBookID
}

type BorrowBookDetailDto struct {
	BorrowBookDto
	Items []BorrowBookItemDto `json:"items"`
}

func (b *BorrowBookDetailDto) FromModel(borrow database.BorrowBook, details []database.BorrowBookDetail) {
	b.BorrowBookDto.FromModel(borrow)
	b.Items = make([]BorrowBookItemDto, len(details))
	for i, item := range details {
		b.Items[i].FromModel(item)
	}
}
