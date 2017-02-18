// schemaregistry provides go-bindigs for Schema Registry (https://github.com/confluentinc/schema-registry)

package schemaregistry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Latest represents the version "latest" in the operations SubjectVersion(subject, version) and
// TestCompatibility(subject, version, schema) of Registry
const Latest = 0

// Compatibility is the type of compatibility supported by the registry. The schema registry server can enforce certain
// compatibility rules when new schemas are registered in a subject.
//go:generate stringer -type=Compatibility
type Compatibility int

const (
	// None means no compatibility: A new schema can be any schema as long as it’s a valid Avro.
	None Compatibility = iota

	// Full means full compatibility: A new schema is fully compatible if it’s both backward and forward compatible
	// with the latest registered schema
	Full

	// Forward means forward compatibility: A new schema is forward compatible if the latest registered schema can
	// read data written in this schema.
	Forward

	// Backward means backward compatibility (default): A new schema is backwards compatible if it can be used to
	// read the data written in the latest registered schema.
	Backward
)

//go:generate stringer -type=ErrorCode
type ErrorCode int

const (
	// SubjectNotFound status code (Subject not found)
	SubjectNotFound ErrorCode = 40401

	// VersionNotFound status code (Version not found)
	VersionNotFound ErrorCode = 40402

	// SchemaNotFound status code (Schema not found)
	SchemaNotFound ErrorCode = 40403

	// InvalidAvroSchema status code (Invalid Avro schema)
	InvalidAvroSchema ErrorCode = 42201

	// InvalidVersion status code (Invalid version)
	InvalidVersion ErrorCode = 42202

	// InvalidCompatibilityLevel status code (Invalid compatibility level)
	InvalidCompatibilityLevel ErrorCode = 42203

	// BackendStoreErr status code (Error in the backend data store)
	BackendStoreErr ErrorCode = 50001

	// OperationTimedOut status code (Operation timed out)
	OperationTimedOut ErrorCode = 50002

	// FwdRequestToMasterErr status code (Error while forwarding the request to the master)
	FwdRequestToMasterErr ErrorCode = 50003
)

// APIError is an error returned by the Schema Registry API
type APIError struct {
	// Code is the error code
	Code ErrorCode `json:"error_code"`

	// Message is the error message
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Schema Registry API error, code: %d message: %s", e.Code, e.Message)
}

// SubjectSchema holds an Avro schema string along with its globally unique identifier and its version under a specific
// subject.
type SubjectSchema struct {
	// Subject is the subject name. A subject refers to the name under which the schema is registered. If you are
	// using the schema registry for Kafka, then a subject refers to either a "<topic>-key" or "<topic>-value"
	// depending on whether you are registering the key schema for that topic or the value schema.
	Subject string `json:"subject"`

	// ID is the unique id of the schema in the registry.
	ID int `json:"id"`

	// Version is the version of the schema in the subject.
	Version int `json:"version"`

	// Schema is the Avro schema string
	Schema string `json:"schema"`
}

// Config holds the configuration (global or of a subject)
type Config struct {
	// Compatibility is the compatibility level in use.
	Compatibility Compatibility
}

// Registry exposes the API operations of Schema Registry (https://github.com/confluentinc/schema-registry)
//
// Check http://docs.confluent.io/3.1.2/schema-registry/docs/api.html for details.
type Registry interface {
	// Schema gets a list of registered subjects.
	Schema(id int) (string, error)

	// Subjects gets a list of registered subjects.
	Subjects() ([]string, error)

	// SubjectVersions gets a list of versions registered under the specified subject.
	SubjectVersions(subject string) ([]int, error)

	// SubjectVersion gets a specific version of the schema registered under this subject.
	SubjectVersion(subject string, version int) (string, error)

	// RegisterSubjectSchema registers a new schema under the specified subject. If successfully registered, this
	// returns the unique identifier of this schema in the registry. The returned identifier should be used to
	// retrieve this schema from the schemas resource and is different from the schema’s version which is associated
	// with the subject. If the same schema is registered under a different subject, the same identifier will be
	// returned. However, the version of the schema may be different under different subjects.
	//
	// A schema should be compatible with the previously registered schema or schemas (if there are any) as per the
	// configured compatibility level. The configured compatibility level can be obtained by issuing a
	// SubjectConfig(subject). If that returns null, then Config()
	//
	// When there are multiple instances of schema registry running in the same cluster, the schema registration
	// request will be forwarded to one of the instances designated as the master. If the master is not available,
	// the client will get an error code indicating that the forwarding has failed.
	RegisterSubjectSchema(subject string, schema string) (int, error)

	// CheckSubjectSchema checks if a schema has already been registered under the specified subject. If so, this
	// returns the schema string along with its globally unique identifier, its version under this subject and the
	// subject name.
	CheckSubjectSchema(subject string, schema string) (*SubjectSchema, error)

	// TestCompatibility tests an input schema against a particular version of a subject’s schema for compatibility.
	// Note that the compatibility level applied for the check is the configured compatibility level for the subject
	// (SubjectConfig(subject)). If this subject’s compatibility level was never changed, then the global
	// compatibility level applies (Config()).
	TestCompatibility(subject string, version int, schema string) (bool, error)

	// SetConfig updates the global compatibility level.
	//
	// When there are multiple instances of schema registry running in the same cluster, the update request will be
	// forwarded to one of the instances designated as the master. If the master is not available, the client will
	// get an error code indicating that the forwarding has failed.
	SetConfig(config *Config) (*Config, error)

	// Config gets the global compatibility level.
	Config() (*Config, error)

	// SetSubjectConfig updates the compatibility level for the specified subject.
	SetSubjectConfig(subject string, config *Config) (*Config, error)

	// SubjectConfig gets the compatibility level for a subject.
	SubjectConfig(subject string) (*Config, error)
}

// New returns the default Registry implementation.
func New(endpoint string) (Registry, error) {
	_, err := url.ParseRequestURI(endpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid endpoint URL: %s", endpoint)
	}
	r := &registry{
		endpoint: endpoint,
	}
	return r, nil
}

// ErrNotImplemented is returned by Registry operations not yet implemented
var ErrNotImplemented = errors.New("Not implemented yet :(")

type schemaJSON struct {
	Schema string `json:"schema"`
}

type schemaIDJSON struct {
	ID int `json:"id"`
}

type registry struct {
	endpoint string
}

func (r *registry) Schema(id int) (string, error) {
	return "", ErrNotImplemented
}

func (r *registry) Subjects() ([]string, error) {
	return nil, ErrNotImplemented
}

func (r *registry) SubjectVersions(subject string) ([]int, error) {
	return nil, ErrNotImplemented
}

func (r *registry) SubjectVersion(subject string, version int) (string, error) {
	return "", ErrNotImplemented
}

func (r *registry) RegisterSubjectSchema(subject string, schema string) (int, error) {
	msg := schemaJSON{
		Schema: schema,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(&msg)
	if err != nil {
		return 0, errors.Wrap(err, "error creating JSON msg in RegisterSubjectSchema")
	}

	operationURL := r.endpoint + "/subjects/" + subject + "/versions"
	resp, err := http.Post(operationURL, "application/vnd.schemaregistry.v1+json", &buf)
	if err != nil {
		return 0, errors.Wrapf(err, "error in POST %s", operationURL)
	}

	if resp.StatusCode != http.StatusOK {
		var errMsg APIError
		err = json.NewDecoder(resp.Body).Decode(&errMsg)
		if err != nil {
			err = errors.Wrapf(err, "error decoding error response, status=%d", resp.StatusCode)
			return 0, err
		}
		return 0, &errMsg
	}

	var respMsg schemaIDJSON
	err = json.NewDecoder(resp.Body).Decode(&respMsg)
	if err != nil {
		return 0, errors.Wrap(err, "error decoding response in RegisterSubjectSchema")
	}
	return respMsg.ID, nil
}

func (r *registry) CheckSubjectSchema(subject string, schema string) (*SubjectSchema, error) {
	msg := schemaJSON{
		Schema: schema,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(&msg)
	if err != nil {
		return nil, errors.Wrap(err, "error creating JSON msg in CheckSubjectSchema")
	}

	operationURL := r.endpoint + "/subjects/" + subject
	resp, err := http.Post(operationURL, "application/vnd.schemaregistry.v1+json", &buf)
	if err != nil {
		return nil, errors.Wrapf(err, "error in POST %s", operationURL)
	}

	if resp.StatusCode != http.StatusOK {
		var errMsg APIError
		err = json.NewDecoder(resp.Body).Decode(&errMsg)
		if err != nil {
			err = errors.Wrapf(err, "error decoding error response, status=%d", resp.StatusCode)
			return nil, err
		}
		return nil, &errMsg
	}

	var ss SubjectSchema
	err = json.NewDecoder(resp.Body).Decode(&ss)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding response in CheckSubjectSchema")
	}
	return &ss, nil
}

func (r *registry) TestCompatibility(subject string, version int, schema string) (bool, error) {
	return false, ErrNotImplemented
}

func (r *registry) SetConfig(config *Config) (*Config, error) {
	return nil, ErrNotImplemented
}

func (r *registry) Config() (*Config, error) {
	return nil, ErrNotImplemented
}

func (r *registry) SetSubjectConfig(subject string, config *Config) (*Config, error) {
	return nil, ErrNotImplemented
}

func (r *registry) SubjectConfig(subject string) (*Config, error) {
	return nil, ErrNotImplemented
}
