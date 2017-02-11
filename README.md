# go-schema-registry

> GO bindings for Confluent [schema-registry](https://github.com/confluentinc/schema-registry).

This package provides bindings for the operations described in http://docs.confluent.io/3.1.2/schema-registry/docs/api.html

**This is work in progress!**

API operation | Binding func | Implemented
--- | --- | ---
GET /schemas/ids/{int: id} | Schema(id int) (string, error) | No
GET /subjects | Subjects() ([]string, error) | No
GET /subjects/(string: subject)/versions | SubjectVersions(subject string) ([]int, error) | No
GET /subjects/(string: subject)/versions/(versionId: version) | SubjectVersion(subject string, version int) (string, error) | No
POST /subjects/(string: subject)/versions | RegisterSubjectSchema(subject string, schema string) (int, error) | No
POST /subjects/(string: subject) | CheckSubjectSchema(subject string, schema string) (*SubjectSchema, error) | Yes
POST /compatibility/subjects/(string: subject)/versions/(versionId: version) | TestCompatibility(subject string, version int, schema string) (bool, error) | No
PUT /config | SetConfig(config *Config) (*Config, error) | No
GET /config | Config() (*Config, error) | No
PUT /config/(string: subject) | SetSubjectConfig(subject string, config *Config) (*Config, error) | No
GET /config/(string: subject) | SubjectConfig(subject string) (*Config, error) | No


Usage:

```go
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
registry, err := schemaregistry.New("http://localhost:8081")
if err != nil {
        // handle err
}
ss, err := registry.CheckSubjectSchema("frames-value", schema)
if err != nil {
        // handle err
}
log.Printf("subject: %s", ss.Subject)
log.Printf("schema ID: %s", ss.ID)
log.Printf("schema version (in subject): %s", ss.Version)
log.Printf("returned schema: %s", ss.Schema)
```

API errors are returned as an *APIError instance, giving access to the error code and message:

```go
_, err := registry.CheckSubjectSchema("inexistent-subject", schema)
if err != nil {
        apiErr, ok := err.(*schemaregistry.APIError)
        if ok {
                switch apiErr.Code {
                case schemaregistry.SubjectNotFound:
                        // subject not found
                case schemaregistry.SchemaNotFound:
                        // schema not found
                default:
                        // other API error, like an Internal server error
                }
        }
        // then, a non-API error, like a connectivity error, JSON encoding error, etc.
}
```


Also, there is a [Testify](https://github.com/stretchr/testify) mock (MockRegistry) available for testing:

```go
testSchema := `{
  "type": "record",
  "name": "Frame",
  "fields": [
    {
      "name": "data",
      "type": "bytes"
    }
  ]
}`
registry := &schemaregistry.MockRegistry{}
registry.On("CheckSubjectSchema", "test-frames-value", testSchema).Return(&schemaregistry.SubjectSchema{
    Subject: "test-frames-value",
    ID:      1,
    Version: 3,
    Schema:  testSchema,
}, nil)

// use the mock registry
ss, err := registry.CheckSubjectSchema("test-frames-value", testSchema)
assert.Nil(t, err)
assert.Equal(t, "test-frames-value", ss.Subject)
assert.Equal(t, 1, ss.ID)
assert.Equal(t, 3, ss.Version)
assert.Equal(t, testSchema, ss.Schema)
```
