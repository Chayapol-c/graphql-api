package repository

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	data2 "graphql-tutorial/data"
	models2 "graphql-tutorial/data/models"
)

// ContactRepository represents the repository for contact operations
type ContactRepository struct {
	DB *data2.DB
}

// NewContactRepo creates a new instance of ContactRepo
func NewContactRepo() *ContactRepository {
	db := data2.NewDB()
	return &ContactRepository{DB: db}
}

// GetContactsBySearchText Get Contacts fetches contacts from the database with support for text search, limit, and offset
func (cr *ContactRepository) GetContactsBySearchText(searchText string, limit, offset int) ([]*models2.ContactModel, error) {
	var contacts []*models2.ContactModel

	query := fmt.Sprintf(`
            SELECT *
            FROM contact
            Where name like '%%%s%%'
               OR first_name like '%%%s%%' 
               OR last_name like '%%%s%%' 
               OR email like '%%%s%%' 
               OR phone like '%%%s%%' 
               OR address like '%%%s%%' 
               OR photo_path like '%%%s%%'
            LIMIT ? OFFSET ?
        `, searchText, searchText, searchText, searchText, searchText, searchText, searchText)

	rows, err := cr.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var contact models2.ContactModel
		err := rows.Scan(
			&contact.ContactId,
			&contact.Name,
			&contact.FirstName,
			&contact.LastName,
			&contact.GenderId,
			&contact.Dob,
			&contact.Email,
			&contact.Phone,
			&contact.Address,
			&contact.PhotoPath,
			&contact.CreatedAt,
			&contact.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, &contact)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}

// GetContactsBySearchTextPagination Get Contacts fetches contacts from the database with support for text search, limit, and offset
func (cr *ContactRepository) GetContactsBySearchTextPagination(searchText string, page, pageSize int) ([]*models2.ContactModel, *models2.PaginationModel, error) {
	var contacts []*models2.ContactModel

	query := fmt.Sprintf(`
            SELECT *
            FROM contact
            Where name like '%%%s%%' 
               OR first_name like '%%%s%%' 
               OR last_name like '%%%s%%' 
               OR email like '%%%s%%' 
               OR phone like '%%%s%%' 
               OR address like '%%%s%%' 
               OR photo_path like '%%%s%%'
        `, searchText, searchText, searchText, searchText, searchText, searchText, searchText)
	offset := (page - 1) * pageSize
	limit := pageSize

	pagination := data2.NewPagination(page, pageSize, query, limit, offset)

	pager, err := pagination.GetPageData(cr.DB)
	if err != nil {
		return nil, nil, err
	}

	query = query + " LIMIT ? OFFSET ?"

	rows, err := cr.DB.Query(query, limit, offset)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var contact models2.ContactModel
		err := rows.Scan(
			&contact.ContactId,
			&contact.Name,
			&contact.FirstName,
			&contact.LastName,
			&contact.GenderId,
			&contact.Dob,
			&contact.Email,
			&contact.Phone,
			&contact.Address,
			&contact.PhotoPath,
			&contact.CreatedAt,
			&contact.CreatedBy,
		)
		if err != nil {
			return nil, nil, err
		}
		contacts = append(contacts, &contact)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return contacts, pager, nil
}
