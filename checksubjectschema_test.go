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

const testSchema = `{
  "type": "record",
  "name": "Frame",
  "fields": [
    {
      "name": "data",
      "type": "bytes"
    }
  ]
}`

func TestRegistry_CheckSubjectSchemaOK(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/subjects/frames-value", r.URL.String())
		assert.Equal(t, "application/vnd.schemaregistry.v1+json", r.Header.Get("Content-Type"))

		var msg map[string]string
		err := json.NewDecoder(r.Body).Decode(&msg)
		require.Nil(t, err)
		assert.Equal(t, testSchema, msg["schema"])

		json.NewEncoder(w).Encode(schemaregistry.SubjectSchema{
			Subject: "frames-value",
			ID:      1,
			Version: 2,
			Schema:  testSchema,
		})
	}))
	defer ts.Close()

	registry, err := schemaregistry.New(ts.URL)
	require.Nil(t, err)

	ss, err := registry.CheckSubjectSchema("frames-value", testSchema)
	require.Nil(t, err)
	assert.Equal(t, "frames-value", ss.Subject)
	assert.Equal(t, 1, ss.ID)
	assert.Equal(t, 2, ss.Version)
	assert.Equal(t, testSchema, ss.Schema)
}

func TestRegistry_CheckSubjectSchemaErrSubjectNotFound(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/subjects/frames-value", r.URL.String())
		assert.Equal(t, "application/vnd.schemaregistry.v1+json", r.Header.Get("Content-Type"))

		var msg map[string]string
		err := json.NewDecoder(r.Body).Decode(&msg)
		require.Nil(t, err)
		assert.Equal(t, testSchema, msg["schema"])

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&schemaregistry.APIError{
			Code:    schemaregistry.SubjectNotFound,
			Message: "Subject not found",
		})
	}))
	defer ts.Close()

	registry, err := schemaregistry.New(ts.URL)
	require.Nil(t, err)

	_, err = registry.CheckSubjectSchema("frames-value", testSchema)
	apiErr, ok := err.(*schemaregistry.APIError)
	require.True(t, ok)
	assert.Equal(t, schemaregistry.SubjectNotFound, apiErr.Code)
	assert.Equal(t, "Subject not found", apiErr.Message)
}

func TestRegistry_CheckSubjectSchemaErrSchemaNotFound(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/subjects/frames-value", r.URL.String())
		assert.Equal(t, "application/vnd.schemaregistry.v1+json", r.Header.Get("Content-Type"))

		var msg map[string]string
		err := json.NewDecoder(r.Body).Decode(&msg)
		require.Nil(t, err)
		assert.Equal(t, testSchema, msg["schema"])

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&schemaregistry.APIError{
			Code:    schemaregistry.SchemaNotFound,
			Message: "Schema not found",
		})
	}))
	defer ts.Close()

	registry, err := schemaregistry.New(ts.URL)
	require.Nil(t, err)

	_, err = registry.CheckSubjectSchema("frames-value", testSchema)
	apiErr, ok := err.(*schemaregistry.APIError)
	require.True(t, ok)
	assert.Equal(t, schemaregistry.SchemaNotFound, apiErr.Code)
	assert.Equal(t, "Schema not found", apiErr.Message)
}

func TestRegistry_CheckSubjectSchemaErrInternalServer(t *testing.T) {
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/subjects/frames-value", r.URL.String())
		assert.Equal(t, "application/vnd.schemaregistry.v1+json", r.Header.Get("Content-Type"))

		var msg map[string]string
		err := json.NewDecoder(r.Body).Decode(&msg)
		require.Nil(t, err)
		assert.Equal(t, testSchema, msg["schema"])

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&schemaregistry.APIError{
			Code:    500,
			Message: "Internal server error",
		})
	}))
	defer ts.Close()

	registry, err := schemaregistry.New(ts.URL)
	require.Nil(t, err)

	_, err = registry.CheckSubjectSchema("frames-value", testSchema)
	apiErr, ok := err.(*schemaregistry.APIError)
	require.True(t, ok)
	assert.EqualValues(t, 500, apiErr.Code)
	assert.Equal(t, "Internal server error", apiErr.Message)
}
