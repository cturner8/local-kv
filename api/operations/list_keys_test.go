package operations

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	utils "cturner8/local-kv/testing"
)

func TestListKeysHandlerWithNoData(t *testing.T) {
	db := utils.SetupDatabase("list_keys")

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	recorder := httptest.NewRecorder()
	listKeysController := NewListKeysController(db)
	handler := http.HandlerFunc(listKeysController.ListKeysHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Keys":[],"NextMarker":null,"Truncated":false}`
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

func TestListKeysHandlerWithData(t *testing.T) {
	t.Skip("fix read only db")

	db := utils.SetupDatabase("list_keys")

	// Create some test data
	now := time.Now()
	_, err := db.Exec(
		`INSERT INTO 
			KeyMetadata (
				id,
				arn,
				awsAccountId,
				createdDate,
				updatedDate,
				customerMasterKeySpec,
				customKeyStoreId,
				description,
				enabled,
				keyManager,
				multiRegion,
				keySpec,
				keyUsage,
				origin
			) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		"1234567890",
		"arn:aws:kms:eu-west-2:000000000000:key/1234567890",
		"000000000000",
		now,
		now,
		nil,
		nil,
		"test key",
		true,
		"CUSTOMER",
		false,
		"RSA_2048",
		"ENCRYPT_DECRYPT",
		"AWS_KMS")
	if err != nil {
		t.Fatal(err)
	}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	recorder := httptest.NewRecorder()
	listKeysController := NewListKeysController(db)
	handler := http.HandlerFunc(listKeysController.ListKeysHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(recorder, req)

	// Check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Keys":[],"NextMarker":null,"Truncated":false}`
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}
