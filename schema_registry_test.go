package schemaregistry_test

import (
	"testing"

	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/larixsource/go-schema-registry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var defaultEndpoint = "http://localhost:8081"

func TestNewInvalidEndpoint(t *testing.T) {
	t.Parallel()
	_, err := schemaregistry.New("asdf")
	if assert.Error(t, err) {
		assert.Equal(t, "invalid endpoint URL: parse asdf: invalid URI for request", err.Error())
	}
}

func TestRegistry_SchemaNotImpl(t *testing.T) {
	t.Parallel()
	registry, err := schemaregistry.New(defaultEndpoint)
	require.Nil(t, err)

	_, err = registry.Schema(0)
	if assert.Error(t, err) {
		assert.Equal(t, schemaregistry.ErrNotImplemented, err)
	}
}

func TestRegistry_SubjectsNotImpl(t *testing.T) {
	t.Parallel()
	registry, err := schemaregistry.New(defaultEndpoint)
	require.Nil(t, err)

	_, err = registry.Subjects()
	if assert.Error(t, err) {
		assert.Equal(t, schemaregistry.ErrNotImplemented, err)
	}
}

func TestRegistry_SubjectVersionsNotImpl(t *testing.T) {
	t.Parallel()
	registry, err := schemaregistry.New(defaultEndpoint)
	require.Nil(t, err)

	_, err = registry.SubjectVersions("")
	if assert.Error(t, err) {
		assert.Equal(t, schemaregistry.ErrNotImplemented, err)
	}
}

func TestRegistry_SubjectVersionNotImpl(t *testing.T) {
	t.Parallel()
	registry, err := schemaregistry.New(defaultEndpoint)
	require.Nil(t, err)

	_, err = registry.SubjectVersion("", 0)
	if assert.Error(t, err) {
		assert.Equal(t, schemaregistry.ErrNotImplemented, err)
	}
}

func TestRegistry_RegisterSubjectSchemaNotImpl(t *testing.T) {
	t.Parallel()
	registry, err := schemaregistry.New(defaultEndpoint)
	require.Nil(t, err)

	_, err = registry.RegisterSubjectSchema("", "")
	if assert.Error(t, err) {
		assert.Equal(t, schemaregistry.ErrNotImplemented, err)
	}
}

func TestRegistry_CheckSubjectSchemaOK(t *testing.T) {
	t.Parallel()
	schema := `{
  "type": "record",
  "name": "Frame",
  "fields": [
    {
      "name": "data",
      "type": "bytes"
    }
  ]
}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/subjects/frames-value", r.URL.String())
		assert.Equal(t, "application/vnd.schemaregistry.v1+json", r.Header.Get("Content-Type"))

		var msg map[string]string
		err := json.NewDecoder(r.Body).Decode(&msg)
		require.Nil(t, err)
		assert.Equal(t, schema, msg["schema"])

		json.NewEncoder(w).Encode(schemaregistry.SubjectSchema{
			Subject: "frames-value",
			ID:      1,
			Version: 2,
			Schema:  schema,
		})
	}))
	defer ts.Close()

	registry, err := schemaregistry.New(ts.URL)
	require.Nil(t, err)

	ss, err := registry.CheckSubjectSchema("frames-value", schema)
	require.Nil(t, err)
	assert.Equal(t, "frames-value", ss.Subject)
	assert.Equal(t, 1, ss.ID)
	assert.Equal(t, 2, ss.Version)
	assert.Equal(t, schema, ss.Schema)
}

func TestRegistry_TestCompatibilityNotImpl(t *testing.T) {
	t.Parallel()
	registry, err := schemaregistry.New(defaultEndpoint)
	require.Nil(t, err)

	_, err = registry.TestCompatibility("", 0, "")
	if assert.Error(t, err) {
		assert.Equal(t, schemaregistry.ErrNotImplemented, err)
	}
}

func TestRegistry_SetConfigNotImpl(t *testing.T) {
	t.Parallel()
	registry, err := schemaregistry.New(defaultEndpoint)
	require.Nil(t, err)

	_, err = registry.SetConfig(schemaregistry.Config{})
	if assert.Error(t, err) {
		assert.Equal(t, schemaregistry.ErrNotImplemented, err)
	}
}

func TestRegistry_ConfigNotImpl(t *testing.T) {
	t.Parallel()
	registry, err := schemaregistry.New(defaultEndpoint)
	require.Nil(t, err)

	_, err = registry.Config()
	if assert.Error(t, err) {
		assert.Equal(t, schemaregistry.ErrNotImplemented, err)
	}
}

func TestRegistry_SetSubjectConfigNotImpl(t *testing.T) {
	t.Parallel()
	registry, err := schemaregistry.New(defaultEndpoint)
	require.Nil(t, err)

	_, err = registry.SetSubjectConfig("", schemaregistry.Config{})
	if assert.Error(t, err) {
		assert.Equal(t, schemaregistry.ErrNotImplemented, err)
	}
}

func TestRegistry_SubjectConfigNotImpl(t *testing.T) {
	t.Parallel()
	registry, err := schemaregistry.New(defaultEndpoint)
	require.Nil(t, err)

	_, err = registry.SubjectConfig("")
	if assert.Error(t, err) {
		assert.Equal(t, schemaregistry.ErrNotImplemented, err)
	}
}
