package users_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/blsmxiu47/fft-api-go/internal/core/app"
)

var a app.App;

const tableCreationSQL = `
CREATE TABLE IF NOT EXISTS users
(
	id SERIAL PRIMARY KEY,
	email VARCHAR (50) UNIQUE,
	firstName VARCHAR (50)
);
`

func TestMain(m *testing.M) {
	// TODO: can we like not call this block in every single script plz
	//   replace with project utils function to get env value given key
	err := godotenv.Load("../../env/.env.local")
	if err != nil{
		log.Fatalf("Error loading .env file: %s", err)
	}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationSQL); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

func checkResponseCode(t *testing.T, want int, got int) {
	if want != got {
		t.Errorf("Expected response code %d. Got %d\n", want, got)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func TestTableEmpty(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	got := response.Body.String()
	want := "[]"

	if got != want {
		t.Errorf("Expected an empty array. Got %s", got)
	}
}
