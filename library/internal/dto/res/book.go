package res

import "library/internal/database"

type BookDto struct {
	Id                int32  `json:"id"`
	Title             string `json:"title"`
	PublishedYear     int32  `json:"published_year"`
	AvailableQuantity int32  `json:"available_quantity"`
	BorrowQuantity    int32  `json:"borrow_quantity"`
}

func (b *BookDto) FromModel(book database.Book) {
	b.Id = book.ID
	b.Title = book.Title
	b.PublishedYear = book.PublishedYear
	b.AvailableQuantity = book.AvailableQuantity
	b.BorrowQuantity = book.BorrowQuantity
}

type BookDetailDto struct {
	*BookDto
	Authors []*AuthorDto `json:"authors"`
	Genres  []*GenreDto  `json:"genres"`
}

func (b *BookDetailDto) FromModel(book database.Book, authors []database.Author, genres []database.Genre) {
	b.BookDto = &BookDto{}
	b.BookDto.FromModel(book)
	authorList := make([]*AuthorDto, len(authors))
	for i, author := range authors {
		temp := &AuthorDto{}
		temp.FromModel(&author)
		authorList[i] = temp
	}
	b.Authors = authorList

	genreList := make([]*GenreDto, len(genres))
	for i, genre := range genres {
		temp := &GenreDto{}
		temp.FromModel(&genre)
		genreList[i] = temp
	}
	b.Genres = genreList
}
