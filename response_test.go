package response

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewResponse(t *testing.T) {
	// Act
	got := CreateResponse()

	// Assert
	if got.Data != nil {
		t.Errorf("\nExpected %v\nGot %v", nil, got.Data)
	}

	if got.Errors != nil {
		t.Errorf("\nExpected %v\nGot %v", nil, got.Errors)
	}
}

func TestSetData(t *testing.T) {
	// Arrange
	data := map[string]string{"field": "value"}

	got := CreateResponse()
	expected := &Response{Data: data}

	// Act
	got.SetData(data)

	// Assert
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\nExpected %v\nGot %v", expected, got)
	}
}

func TestAppendError(t *testing.T) {
	// Arrange
	err := errors.New("message")

	got := CreateResponse()
	expected := &Response{Errors: []Error{{Message: err.Error()}}}

	// Act
	got.AppendError(err)

	// Assert
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\nExpected %v\nGot %v", expected, got)
	}
}

func TestSendWithStatusCode(t *testing.T) {
	// Arrange
	data := map[string]string{"field": "value"}
	err := errors.New("message")
	w := httptest.NewRecorder()
	responseToSend := CreateResponse()

	expectedModel, _ := json.Marshal(&Response{Data: data, Errors: []Error{{Message: err.Error()}}})
	expectedStatusCode := 200

	// Act
	responseToSend.SetData(data)
	responseToSend.AppendError(err)
	responseToSend.SendWithStatusCode(w, 200)
	resp := w.Result()
	got, _ := io.ReadAll(resp.Body)
	// encoder adds 0 byte at end for delimiter
	got = got[:len(got)-1]

	// Assert
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected %v\nGot %v", expectedStatusCode, got)
	}

	if !bytes.Equal(got, expectedModel) {
		t.Errorf("\nExpected %v\nGot %v", expectedModel, got)
	}
}

func TestSendDataWithStatusCode(t *testing.T) {
	// Arrange
	data := map[string]string{"field": "value"}
	w := httptest.NewRecorder()
	responseToSend := CreateResponse()

	expectedModel, _ := json.Marshal(&Response{Data: data})
	expectedStatusCode := 200

	// Act
	responseToSend.SendDataWithStatusCode(w, data, 200)
	resp := w.Result()
	got, _ := io.ReadAll(resp.Body)
	// encoder adds 0 byte at end for delimiter
	got = got[:len(got)-1]

	// Assert
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected %v\nGot %v", expectedStatusCode, got)
	}

	if !bytes.Equal(got, expectedModel) {
		t.Errorf("\nExpected %v\nGot %v", expectedModel, got)
	}
}

func TestSendErrorWithStatusCode(t *testing.T) {
	// Arrange
	err := errors.New("message")
	w := httptest.NewRecorder()
	responseToSend := CreateResponse()

	expectedModel, _ := json.Marshal(&Response{Errors: []Error{{Message: err.Error()}}})
	expectedStatusCode := 200

	// Act
	responseToSend.SendErrorWithStatusCode(w, err, 200)
	resp := w.Result()
	got, _ := io.ReadAll(resp.Body)
	// encoder adds 0 byte at end for delimiter
	got = got[:len(got)-1]

	// Assert
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected %v\nGot %v", expectedStatusCode, got)
	}

	if !bytes.Equal(got, expectedModel) {
		t.Errorf("\nExpected %v\nGot %v", expectedModel, got)
	}
}
