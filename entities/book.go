package entities

// Books :struct model for book
type Books struct {
	BookID        int         `json:"bookId"`
	AuthorID      int         `json:"authorId"`
	Title         string      `json:"title"`
	Publications  Publication `json:"publication"`
	PublishedDate string      `json:"publishedDate"`
	Author        Author      `json:"author,omitempty"`
}

// Publication :enum
type Publication string

const (
	Arihant    Publication = "arihant"
	Scholastic Publication = "scholastic"
	Penguin    Publication = "penguin"
)
