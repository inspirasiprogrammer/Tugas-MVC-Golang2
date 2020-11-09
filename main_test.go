package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	BadRequestCode     = 400
	SuccessRequestCode = 200
)

type TestStruct struct {
	requestBody        string
	expectedStatusCode int
	responseBody       string
	observedStatusCode int
}

func TestUserSignup(t *testing.T) {

	url := "http://localhost:8081/api/v1/account/add"

	tests := []TestStruct{
		{`{}`, BadRequestCode, "", 0},
		{`{"name":""}`, BadRequestCode, "", 0},
		{`{"name":"Irwan Syahputra","password":"123456"}`, BadRequestCode, "", 0},
		{`{"name":"Irwan Syahputra","password":"123456","saldo" : 1000000}`, SuccessRequestCode, "", 0},
		{`{"name":"Irwan Syahputra","password":"123456","saldo" : 1000000}`, BadRequestCode, "", 0}, // request should fail because of
	}

	for i, testCase := range tests {

		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader

		request, err := http.NewRequest("POST", url, reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err) //Something is wrong while sending request
		}
		body, _ := ioutil.ReadAll(res.Body)

		tests[i].responseBody = strings.TrimSpace(string(body))
		tests[i].observedStatusCode = res.StatusCode

	}

	DisplayTestCaseResults("UserSignup", tests, t)

}

func TestUserSignin(t *testing.T) {

	url := "http://localhost:8081/api/v1/login"

	tests := []TestStruct{
		{`{}`, BadRequestCode, "", 0},
		{`{"name":""}`, BadRequestCode, "", 0},
		{`{"name":"Irwan Syahputra","password":"123456"}`, SuccessRequestCode, "", 0},
		{`{"name":"Irwan Syahputra","password":"123456"}`, BadRequestCode, "", 0}, // request should fail because of
	}

	for i, testCase := range tests {

		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader

		request, err := http.NewRequest("POST", url, reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err) //Something is wrong while sending request
		}
		body, _ := ioutil.ReadAll(res.Body)

		tests[i].responseBody = strings.TrimSpace(string(body))
		tests[i].observedStatusCode = res.StatusCode

	}

	DisplayTestCaseResults("UserSignin", tests, t)
}

func TestUserGetAccountTrans(t *testing.T) {

	url := "http://localhost:8081/api/v1/account/"

	tests := []TestStruct{
		{`{}`, BadRequestCode, "", 0},
		{`{"name":""}`, BadRequestCode, "", 0},
		{`{"name":"Irwan Syahputra","password":"123456"}`, SuccessRequestCode, "", 0},
		{`{"name":"Irwan Syahputra","password":"123456"}`, BadRequestCode, "", 0}, // request should fail because of
	}

	for i, testCase := range tests {

		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader

		request, err := http.NewRequest("POST", url, reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err) //Something is wrong while sending request
		}
		body, _ := ioutil.ReadAll(res.Body)

		tests[i].responseBody = strings.TrimSpace(string(body))
		tests[i].observedStatusCode = res.StatusCode

	}

	DisplayTestCaseResults("UserSignin", tests, t)
}

func DisplayTestCaseResults(functionalityName string, tests []TestStruct, t *testing.T) {

	for _, test := range tests {

		if test.observedStatusCode == test.expectedStatusCode {
			t.Logf("Passed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		} else {
			t.Errorf("Failed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		}
	}
}
