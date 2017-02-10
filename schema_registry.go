// schemaregistry provides go-bindigs for Schema Registry (https://github.com/confluentinc/schema-registry)

package schemaregistry

import (
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

// SubjectSchema holds an Avro schema string along with its globally unique identifier and its version under a specific
// subject.
type SubjectSchema struct {
	// Subject is the subject name. A subject refers to the name under which the schema is registered. If you are
	// using the schema registry for Kafka, then a subject refers to either a "<topic>-key" or "<topic>-value"
	// depending on whether you are registering the key schema for that topic or the value schema.
	Subject string

	// ID is the unique id of the schema in the registry.
	ID int

	// Version is the version of the schema in the subject.
	Version int

	// Schema is the Avro schema string
	Schema string
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
	CheckSubjectSchema(subject string, schema string) (SubjectSchema, error)

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
	SetConfig(config Config) (Config, error)

	// Config gets the global compatibility level.
	Config() (Config, error)

	// SetSubjectConfig updates the compatibility level for the specified subject.
	SetSubjectConfig(subject string, config Config) (Config, error)

	// SubjectConfig gets the compatibility level for a subject.
	SubjectConfig(subject string) (Config, error)
}

// New returns the default Registry implementation.
func New(endpoint string) (Registry, error) {
	_, err := url.ParseRequestURI(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint URL")
	}
	r := &registry{}
	return r, nil
}

// ErrNotImplemented is returned by Registry operations not yet implemented
var ErrNotImplemented = errors.New("Not implemented yet :(")

type registry struct {
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
	return 0, ErrNotImplemented
}

func (r *registry) CheckSubjectSchema(subject string, schema string) (SubjectSchema, error) {
	return SubjectSchema{}, ErrNotImplemented
}

func (r *registry) TestCompatibility(subject string, version int, schema string) (bool, error) {
	return false, ErrNotImplemented
}

func (r *registry) SetConfig(config Config) (Config, error) {
	return Config{}, ErrNotImplemented
}

func (r *registry) Config() (Config, error) {
	return Config{}, ErrNotImplemented
}

func (r *registry) SetSubjectConfig(subject string, config Config) (Config, error) {
	return Config{}, ErrNotImplemented
}

func (r *registry) SubjectConfig(subject string) (Config, error) {
	return Config{}, ErrNotImplemented
}
