package schemaregistry_test

import (
	"testing"

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

func TestRegistry_CheckSubjectSchemaNotImpl(t *testing.T) {
	t.Parallel()
	registry, err := schemaregistry.New(defaultEndpoint)
	require.Nil(t, err)

	_, err = registry.CheckSubjectSchema("", "")
	if assert.Error(t, err) {
		assert.Equal(t, schemaregistry.ErrNotImplemented, err)
	}
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
