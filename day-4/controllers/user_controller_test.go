package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Jiran03/agmc/task/day4/config"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Method           string
	Name             string
	Path             string
	ExpectedStatus   int
	ExpectedResponse string
}

func InitEchoTest() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	// e.Start(":8800")
	return e
}

func TestCreateUsersController(t *testing.T) {
	testingCase := []TestCase{{
		Method:           http.MethodPost,
		Name:             "Create User",
		Path:             "/users",
		ExpectedStatus:   http.StatusOK,
		ExpectedResponse: "jiran",
	},
	}

	e := InitEchoTest()
	reqStr := `{
		"name":"jiran",
		"email": "jiran@mail.com",
		"password": "sayatetau"
	}`
	fmt.Println(reqStr)

	for _, testCase := range testingCase {
		req := httptest.NewRequest(testCase.Method, testCase.Path, strings.NewReader(reqStr))
		fmt.Println("req: ", req)
		//set request header
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		fmt.Println("context", ctx)

		//pengecekan test
		if assert.NoError(t, CreateUserController(ctx)) {
			assert.Equal(t, testCase.ExpectedStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), testCase.ExpectedResponse)
		}
	}

}

func TestGetAllUserController(t *testing.T) {
	testingCase := []TestCase{{
		Method:           http.MethodGet,
		Name:             "Get All User",
		Path:             "/users",
		ExpectedStatus:   http.StatusOK,
		ExpectedResponse: "jiran",
	},
	}

	e := InitEchoTest()
	for _, testCase := range testingCase {
		req := httptest.NewRequest(testCase.Method, testCase.Path, nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		//pengecekan test
		if assert.NoError(t, GetAllUserController(ctx)) {
			assert.Equal(t, testCase.ExpectedStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), testCase.ExpectedResponse)
		}
	}

}

func TestGetUserByIdController(t *testing.T) {
	testingCase := []TestCase{{
		Method:           http.MethodGet,
		Name:             "Get User",
		Path:             "/users",
		ExpectedStatus:   http.StatusOK,
		ExpectedResponse: "jiran",
	},
	}

	e := InitEchoTest()
	for _, testCase := range testingCase {
		req := httptest.NewRequest(testCase.Method, testCase.Path, nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		//pengecekan test
		if assert.NoError(t, GetUserByIdController(ctx)) {
			assert.Equal(t, testCase.ExpectedStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), testCase.ExpectedResponse)
		}
	}
}

func TestUpdateUserController(t *testing.T) {
	testingCase := []TestCase{{
		Method:           http.MethodPut,
		Name:             "Update Users",
		Path:             "/users",
		ExpectedStatus:   http.StatusOK,
		ExpectedResponse: "bambang",
	},
	}

	reqStr := `{
		"name":"bambang saja",
		"email": "jiran@mail.com",
		"password": "sayatetau"
	}`

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjM0MDAxMjIsInVzZXJFbWFpbCI6ImppcmFuQG1haWwuY29tIiwidXNlcklEIjo1fQ.-ol201NXYRH09JO22AKQz_9zs7Ncs2EOwuAdKt_vSJw"
	bearer := "Bearer " + token

	e := InitEchoTest()
	for _, testCase := range testingCase {
		req := httptest.NewRequest(testCase.Method, testCase.Path, strings.NewReader(reqStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, bearer)
		fmt.Println(req.Header)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		fmt.Println(ctx.ParamValues())
		//pengecekan test
		if assert.NoError(t, UpdateUserController(ctx)) {
			assert.Equal(t, testCase.ExpectedStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), testCase.ExpectedResponse)
		}
	}
}

func TestLoginUserController(t *testing.T) {
	testingCase := []TestCase{{
		Method:           http.MethodPost,
		Name:             "Login Users",
		Path:             "/login",
		ExpectedStatus:   http.StatusOK,
		ExpectedResponse: "bambang",
	},
	}

	reqStr := `{
	"email": "bambang@mail.com",
	"password": "gantipasswordsaya"
	}`

	e := InitEchoTest()

	for _, testCase := range testingCase {
		req := httptest.NewRequest(testCase.Method, testCase.Path, strings.NewReader(reqStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)

		//pengecekan test
		if assert.NoError(t, LoginUserController(ctx)) {
			assert.Equal(t, testCase.ExpectedStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), testCase.ExpectedResponse)
		}
	}
}

func TestDeleteUserController(t *testing.T) {
	testingCase := []TestCase{{
		Method:           http.MethodPost,
		Name:             "Delete User",
		Path:             "/users/?id=1",
		ExpectedStatus:   http.StatusOK,
		ExpectedResponse: "user deleted",
	},
	}

	e := InitEchoTest()

	for _, testCase := range testingCase {
		req := httptest.NewRequest(testCase.Method, testCase.Path, nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		//pengecekan test
		if assert.NoError(t, DeleteUserController(ctx)) {
			assert.Equal(t, testCase.ExpectedStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), testCase.ExpectedResponse)
		}
	}

}
