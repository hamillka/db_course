package repositories

import (
	goErrors "errors"
)

var (
	ErrRecordNotFound        = goErrors.New("record was not found")
	ErrDatabaseWritingError  = goErrors.New("error while writing to DB")
	ErrDatabaseReadingError  = goErrors.New("error while reading from DB")
	ErrRecordAlreadyExists   = goErrors.New("record with this data already exists")
	ErrDatabaseDeletingError = goErrors.New("record was not deleted")
	ErrDatabaseUpdatingError = goErrors.New("record was not updated")
)
