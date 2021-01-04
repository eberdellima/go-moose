package tests

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"go-moose/database"
	"go-moose/src"
	"go-moose/tests/utilities"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var server *httptest.Server
var db *gorm.DB

func TestMain(m *testing.M) {

	//utilities.LoadEnv("../.env")

	db = database.InitializeDBConnection()
	defer db.Close()

	server = httptest.NewServer(src.ConfigureAPI())

	os.Exit(m.Run())
}

func TestRegistration(t *testing.T) {

	body, err := utilities.CreateRequestBody(utilities.CreateRegistrationInput())

	if err != nil {
		t.Fatal(err.Error())
	}

	httpClient := server.Client()
	response, err := httpClient.Post(server.URL+"/register", "application/json", body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestLogin(t *testing.T) {

	password := "sample Password"
	user := utilities.CreateUser(password)
	body, err := utilities.CreateRequestBody(utilities.CreateLoginInput(user.Email, password))

	if err != nil {
		t.Fatal(err.Error())
	}

	httpClient := server.Client()
	response, err := httpClient.Post(server.URL+"/login", "application/json", body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestLogout(t *testing.T) {

	user := utilities.CreateUser("")
	tokenPair := utilities.CreateTokenPair(user)

	req, err := http.NewRequest("POST", server.URL+"/logout", nil)

	if err != nil {
		t.Fatal(err.Error())
	}

	req.Header.Add("Authorization", "Bearer "+tokenPair.AccessToken)

	httpClient := server.Client()
	response, err := httpClient.Do(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}

func TestRefreshToken(t *testing.T) {

	user := utilities.CreateUser("")
	body, err := utilities.CreateRequestBody(utilities.CreateRefreshTokenInput(user))

	if err != nil {
		t.Fatal(err.Error())
	}

	req, err := http.NewRequest(http.MethodPatch, server.URL+"/refresh-token", body)

	if err != nil {
		t.Fatal(err.Error())
	}

	httpClient := server.Client()
	response, err := httpClient.Do(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
