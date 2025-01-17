package data

import (
	"fmt"
	"graphql-tutorial/data/models"
	"math"

	_ "github.com/mattn/go-sqlite3"
)

// Pagination struct for handling paginated results
type Pagination struct {
	Page        int
	PageSize    int
	TotalPages  int
	TotalItems  int
	HasNext     bool
	HasPrevious bool
	query       string
	parameters  []interface{}
}

// NewPagination creates a new Pagination object with the given page and page size
func NewPagination(page, pageSize int, query string, parameters ...interface{}) *Pagination {
	return &Pagination{
		Page:       page,
		PageSize:   pageSize,
		query:      query,
		parameters: parameters,
	}
}

// GetPageData fetches the paginated data and calculates pagination info
func (p *Pagination) GetPageData(db *DB) (*models.PaginationModel, error) {
	// Get total count of items
	totalCountQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s)", p.query)

	// fmt.Println( "parameters")
	fmt.Println(p.parameters...)
	row, err := db.QueryRow(totalCountQuery, p.parameters...)
	if err != nil {
		return nil, err
	}
	pager := &models.PaginationModel{}
	row.Scan(&p.TotalItems)
	// fmt.Println("Total Items", p.TotalItems)
	pager.TotalItems = p.TotalItems

	// Calculate total pages
	p.TotalPages = int(math.Ceil(float64(p.TotalItems) / float64(p.PageSize)))
	pager.TotalPages = p.TotalPages
	// Check if there's a next or previous page
	p.HasNext = p.Page < p.TotalPages
	pager.HasNext = p.HasNext
	p.HasPrevious = p.Page > 1
	pager.HasPrevious = p.HasPrevious
	pager.Page = p.Page
	pager.PageSize = p.PageSize

	return pager, nil
}

// Next increments the page number
func (p *Pagination) Next() {
	if p.HasNext {
		p.Page++
	}
}

// Previous decrements the page number
func (p *Pagination) Previous() {
	if p.HasPrevious {
		p.Page--
	}
}
