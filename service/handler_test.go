package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/oentoro/ms.account/dbclient"
	"github.com/oentoro/ms.account/model"
	. "github.com/smartystreets/goconvey/convey"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func TestGetAccount(t *testing.T) {
	// Create a mock instance that implements the IBoltClient interface
	mockRepo := &dbclient.MockBoltClient{}

	// Declare two mock behaviours. For "123" as input, return a proper Account struct and nil as error.
	// For "456" as input, return an empty Account object and a real error.
	mockRepo.On("QueryAccount", "123").Return(model.Account{Id: "123", Name: "Person_123"}, nil)
	mockRepo.On("QueryAccount", "456").Return(model.Account{}, fmt.Errorf("Some error"))

	// Finally, assign mockRepo to the DBClient field (it's in _handlers.go_, e.g. in the same package)
	DBClient = mockRepo

	Convey("Given a HTTP request for /accounts/123", t, func() {
		Convey("When the request is handled by the Router", func() {
			router := setupRouter()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/accounts/123", nil)
			router.ServeHTTP(w, req)

			Convey("Then the response should be a 200", func() {
				So(w.Code, ShouldEqual, 200)

				account := model.Account{}
				json.Unmarshal(w.Body.Bytes(), &account)
				So(account.Id, ShouldEqual, "123")
				So(account.Name, ShouldEqual, "Person_123")
			})
		})
	})
}

func TestGetAccountWrongPath(t *testing.T) {
	Convey("Given a HTTP request for /invalid/123", t, func() {
		Convey("When the request is handled by the Router", func() {
			// NewRouter().ServeHTTP(resp, req)
			router := setupRouter()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/invalid/123", nil)
			router.ServeHTTP(w, req)

			Convey("Then the response should be a 404", func() {
				So(w.Code, ShouldEqual, 404)
			})
		})
	})
}
