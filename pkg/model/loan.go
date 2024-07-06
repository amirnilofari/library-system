package model

import "time"

type Loan struct {
	ID         int       `json:"id"`
	BookID     int       `json:"book_id"`
	UserID     int       `json:"user_id"`
	LoanDate   time.Time `json:"loan_date"`
	ReturnDate time.Time `json:"return_date"`
	Returned   bool      `json:"returned"`
}

// NewLoan creates a new Loan instance
func NewLoan(bookID, userID int) *Loan {
	return &Loan{
		BookID:   bookID,
		UserID:   userID,
		LoanDate: time.Now(),
		Returned: false,
	}
}

func (l *Loan) MarkAsReturned() {
	l.Returned = true
	l.ReturnDate = time.Now()
}
