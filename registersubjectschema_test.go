package schemaregistry_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/larixsource/go-schema-registry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistry_RegisterSubjectSchemaOK(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/subjects/frames-value/versions", r.URL.String())
		assert.Equal(t, "application/vnd.schemaregistry.v1+json", r.Header.Get("Content-Type"))

		var msg map[string]string
		err := json.NewDecoder(r.Body).Decode(&msg)
		require.Nil(t, err)
		assert.Equal(t, testSchema, msg["schema"])

		json.NewEncoder(w).Encode(map[string]interface{}{"id": 1})
	}))
	defer ts.Close()

	registry, err := schemaregistry.New(ts.URL)
	require.Nil(t, err)

	id, err := registry.RegisterSubjectSchema("frames-value", testSchema)
	require.Nil(t, err)
	assert.Equal(t, 1, id)
}

func TestRegistry_RegisterSubjectSchemaErrIncompatibleAvroSchema(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/subjects/frames-value/versions", r.URL.String())
		assert.Equal(t, "application/vnd.schemaregistry.v1+json", r.Header.Get("Content-Type"))

		var msg map[string]string
		err := json.NewDecoder(r.Body).Decode(&msg)
		require.Nil(t, err)
		assert.Equal(t, testSchema, msg["schema"])

		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&schemaregistry.APIError{
			Code:    schemaregistry.ErrorCode(http.StatusConflict),
			Message: "Incompatible Avro schema",
		})
	}))
	defer ts.Close()

	registry, err := schemaregistry.New(ts.URL)
	require.Nil(t, err)

	_, err = registry.RegisterSubjectSchema("frames-value", testSchema)
	apiErr, ok := err.(*schemaregistry.APIError)
	require.True(t, ok)
	assert.EqualValues(t, http.StatusConflict, apiErr.Code)
	assert.Equal(t, "Incompatible Avro schema", apiErr.Message)
}

func TestRegistry_RegisterSubjectSchemaErrInvalidAvroSchema(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/subjects/frames-value/versions", r.URL.String())
		assert.Equal(t, "application/vnd.schemaregistry.v1+json", r.Header.Get("Content-Type"))

		var msg map[string]string
		err := json.NewDecoder(r.Body).Decode(&msg)
		require.Nil(t, err)
		assert.Equal(t, testSchema, msg["schema"])

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(&schemaregistry.APIError{
			Code:    schemaregistry.InvalidAvroSchema,
			Message: "Invalid Avro schema",
		})
	}))
	defer ts.Close()

	registry, err := schemaregistry.New(ts.URL)
	require.Nil(t, err)

	_, err = registry.RegisterSubjectSchema("frames-value", testSchema)
	apiErr, ok := err.(*schemaregistry.APIError)
	require.True(t, ok)
	assert.Equal(t, schemaregistry.InvalidAvroSchema, apiErr.Code)
	assert.Equal(t, "Invalid Avro schema", apiErr.Message)
}

func TestRegistry_RegisterSubjectSchemaErrBackendStore(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/subjects/frames-value/versions", r.URL.String())
		assert.Equal(t, "application/vnd.schemaregistry.v1+json", r.Header.Get("Content-Type"))

		var msg map[string]string
		err := json.NewDecoder(r.Body).Decode(&msg)
		require.Nil(t, err)
		assert.Equal(t, testSchema, msg["schema"])

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&schemaregistry.APIError{
			Code:    schemaregistry.BackendStoreErr,
			Message: "Error in the backend data store",
		})
	}))
	defer ts.Close()

	registry, err := schemaregistry.New(ts.URL)
	require.Nil(t, err)

	_, err = registry.RegisterSubjectSchema("frames-value", testSchema)
	apiErr, ok := err.(*schemaregistry.APIError)
	require.True(t, ok)
	assert.Equal(t, schemaregistry.BackendStoreErr, apiErr.Code)
	assert.Equal(t, "Error in the backend data store", apiErr.Message)
}

func TestRegistry_RegisterSubjectSchemaErrOperationTimeout(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/subjects/frames-value/versions", r.URL.String())
		assert.Equal(t, "application/vnd.schemaregistry.v1+json", r.Header.Get("Content-Type"))

		var msg map[string]string
		err := json.NewDecoder(r.Body).Decode(&msg)
		require.Nil(t, err)
		assert.Equal(t, testSchema, msg["schema"])

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&schemaregistry.APIError{
			Code:    schemaregistry.OperationTimedOut,
			Message: "Operation timed out",
		})
	}))
	defer ts.Close()

	registry, err := schemaregistry.New(ts.URL)
	require.Nil(t, err)

	_, err = registry.RegisterSubjectSchema("frames-value", testSchema)
	apiErr, ok := err.(*schemaregistry.APIError)
	require.True(t, ok)
	assert.Equal(t, schemaregistry.OperationTimedOut, apiErr.Code)
	assert.Equal(t, "Operation timed out", apiErr.Message)
}

func TestRegistry_RegisterSubjectSchemaErrFwdRequestToMaster(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/subjects/frames-value/versions", r.URL.String())
		assert.Equal(t, "application/vnd.schemaregistry.v1+json", r.Header.Get("Content-Type"))

		var msg map[string]string
		err := json.NewDecoder(r.Body).Decode(&msg)
		require.Nil(t, err)
		assert.Equal(t, testSchema, msg["schema"])

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&schemaregistry.APIError{
			Code:    schemaregistry.FwdRequestToMasterErr,
			Message: "Error while forwarding the request to the master",
		})
	}))
	defer ts.Close()

	registry, err := schemaregistry.New(ts.URL)
	require.Nil(t, err)

	_, err = registry.RegisterSubjectSchema("frames-value", testSchema)
	apiErr, ok := err.(*schemaregistry.APIError)
	require.True(t, ok)
	assert.Equal(t, schemaregistry.FwdRequestToMasterErr, apiErr.Code)
	assert.Equal(t, "Error while forwarding the request to the master", apiErr.Message)
}
