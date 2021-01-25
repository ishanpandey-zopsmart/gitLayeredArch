package service
//
//import (
//	"github.com/DATA-DOG/go-sqlmock"
//	"net/http"
//	"testing"
//)
//
//func TestGetByName(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("An error '%s' was not expected when opening connection with Mock database connection", err)
//	}
//	defer  db.Close()
//
//	mock.ExpectBegin()
//	res, err := http.NewRequest(http.MethodGet, "http://localhost:8000/customer", nil)
//	if err != nil {
//		t.Fatalf("An error occurred while creating a request")
//	}
//
//
//}
