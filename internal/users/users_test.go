package users_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/blsmxiu47/fft-api-go/internal/core/app"
	"github.com/blsmxiu47/fft-api-go/internal/utils"
)

var a app.App;

const usersTableCreationSQL = `
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
	a.Initialize(
		utils.GetEnv("APP_DB_USERNAME"),
		utils.GetEnv("APP_DB_PASSWORD"),
		utils.GetEnv("APP_DB_NAME"),
	)

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(usersTableCreationSQL); err != nil {
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

func TestUsers_TableEmpty(t *testing.T) {
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

func TestUsers_GetUser_Fail(t *testing.T) {
    clearTable()

    req, _ := http.NewRequest("GET", "/user/9999", nil)
    response := executeRequest(req)

    checkResponseCode(t, http.StatusNotFound, response.Code)

    var m map[string]string
    json.Unmarshal(response.Body.Bytes(), &m)

    got := m["error"]
    want := "User not found"
    if got != want {
        t.Errorf("Expected the 'error' key of the response to be set to '%s'. Got '%s'", want, got)
    }
}
