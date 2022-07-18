package entities

// Author : struct model for author
type Author struct {
	AuthorID    int    `json:"authorId,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	DateOfBirth string `json:"DOB,omitempty"`
	PenName     string `json:"penName,omitempty"`
}
