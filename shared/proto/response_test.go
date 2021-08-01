package proto

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockedWriter struct {
	code   int
	header http.Header
	bytes.Buffer
}

type TestError struct {
}

func (te *TestError) Error() string {
	return "test error"
}

func newWriter() *MockedWriter {
	return &MockedWriter{
		header: make(http.Header),
	}
}

func (mw *MockedWriter) Header() http.Header {
	return mw.header
}

func (mw *MockedWriter) WriteHeader(code int) {
	mw.code = code
}

func TestHttpError(t *testing.T) {
	cases := []struct {
		httpCode int
		appCode  int64
		error    error
	}{
		{
			http.StatusBadRequest,
			1,
			new(TestError),
		},
	}

	for _, each := range cases {
		w := newWriter()
		HttpError(w, each.httpCode, each.appCode, each.error.Error(), nil)

		assert.Equal(t, each.httpCode, w.code)
		expect := response{
			Code: each.appCode,
			Data: each.error,
		}
		var actual response
		decoder := json.NewDecoder(&w.Buffer)
		if err := decoder.Decode(&actual); err != nil {
			assert.Fail(t, err.Error())
		}

		assert.Equal(t, expect.Code, actual.Code)
		assert.ObjectsAreEqualValues(expect.Data, actual.Data)
	}
}

func TestHttpOk(t *testing.T) {
	cases := []interface{}{
		nil,
		"name",
		3,
		struct {
			Name  string
			Value int
		}{
			Name:  "kevin",
			Value: 3,
		},
	}

	for _, each := range cases {
		w := newWriter()
		HttpOk(w, each)

		var actual response
		decoder := json.NewDecoder(&w.Buffer)
		if err := decoder.Decode(&actual); err != nil {
			assert.Fail(t, err.Error())
		}

		assert.Equal(t, http.StatusOK, w.code)
		assert.Equal(t, int64(0), actual.Code)
		assert.ObjectsAreEqualValues(each, actual.Data)
	}
}

func TestHttpParamError(t *testing.T) {
	cases := []string{
		"first error",
		"second error",
	}

	for _, each := range cases {
		w := newWriter()
		HttpParamError(w, each)

		var actual response
		decoder := json.NewDecoder(&w.Buffer)
		if err := decoder.Decode(&actual); err != nil {
			assert.Fail(t, err.Error())
		}

		assert.Equal(t, int64(http.StatusBadRequest), actual.Code)
		assert.ObjectsAreEqualValues(each, actual.Data)
	}
}
