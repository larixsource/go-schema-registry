package schemaregistry

import (
	"github.com/stretchr/testify/mock"
)

// MockRegistry is a test double for Registry.
// Generated with github.com/xeger/mongoose; do not edit by hand.
type MockRegistry struct {
	mock.Mock
}

func (_m *MockRegistry) CheckSubjectSchema(subject string, schema string) (SubjectSchema, error) {
	ret := _m.Called(subject, schema)

	var r0 SubjectSchema

	if r0f, ok := ret.Get(0).(func(string, string) SubjectSchema); ok {
		r0 = r0f(subject, schema)
	} else {
		r0 = ret.Get(0).(SubjectSchema)
	}
	var r1 error

	if r1f, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = r1f(subject, schema)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockRegistry) Config() (Config, error) {
	ret := _m.Called()

	var r0 Config

	if r0f, ok := ret.Get(0).(func() Config); ok {
		r0 = r0f()
	} else {
		r0 = ret.Get(0).(Config)
	}
	var r1 error

	if r1f, ok := ret.Get(1).(func() error); ok {
		r1 = r1f()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockRegistry) RegisterSubjectSchema(subject string, schema string) (int, error) {
	ret := _m.Called(subject, schema)

	var r0 int

	if r0f, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = r0f(subject, schema)
	} else {
		r0 = ret.Get(0).(int)
	}
	var r1 error

	if r1f, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = r1f(subject, schema)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockRegistry) Schema(id int) (string, error) {
	ret := _m.Called(id)

	var r0 string

	if r0f, ok := ret.Get(0).(func(int) string); ok {
		r0 = r0f(id)
	} else {
		r0 = ret.Get(0).(string)
	}
	var r1 error

	if r1f, ok := ret.Get(1).(func(int) error); ok {
		r1 = r1f(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockRegistry) SetConfig(config Config) (Config, error) {
	ret := _m.Called(config)

	var r0 Config

	if r0f, ok := ret.Get(0).(func(Config) Config); ok {
		r0 = r0f(config)
	} else {
		r0 = ret.Get(0).(Config)
	}
	var r1 error

	if r1f, ok := ret.Get(1).(func(Config) error); ok {
		r1 = r1f(config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockRegistry) SetSubjectConfig(subject string, config Config) (Config, error) {
	ret := _m.Called(subject, config)

	var r0 Config

	if r0f, ok := ret.Get(0).(func(string, Config) Config); ok {
		r0 = r0f(subject, config)
	} else {
		r0 = ret.Get(0).(Config)
	}
	var r1 error

	if r1f, ok := ret.Get(1).(func(string, Config) error); ok {
		r1 = r1f(subject, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockRegistry) SubjectConfig(subject string) (Config, error) {
	ret := _m.Called(subject)

	var r0 Config

	if r0f, ok := ret.Get(0).(func(string) Config); ok {
		r0 = r0f(subject)
	} else {
		r0 = ret.Get(0).(Config)
	}
	var r1 error

	if r1f, ok := ret.Get(1).(func(string) error); ok {
		r1 = r1f(subject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockRegistry) SubjectVersion(subject string, version int) (string, error) {
	ret := _m.Called(subject, version)

	var r0 string

	if r0f, ok := ret.Get(0).(func(string, int) string); ok {
		r0 = r0f(subject, version)
	} else {
		r0 = ret.Get(0).(string)
	}
	var r1 error

	if r1f, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = r1f(subject, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockRegistry) SubjectVersions(subject string) ([]int, error) {
	ret := _m.Called(subject)

	var r0 []int

	if r0f, ok := ret.Get(0).(func(string) []int); ok {
		r0 = r0f(subject)
	} else {
		r0 = ret.Get(0).([]int)
	}
	var r1 error

	if r1f, ok := ret.Get(1).(func(string) error); ok {
		r1 = r1f(subject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockRegistry) Subjects() ([]string, error) {
	ret := _m.Called()

	var r0 []string

	if r0f, ok := ret.Get(0).(func() []string); ok {
		r0 = r0f()
	} else {
		r0 = ret.Get(0).([]string)
	}
	var r1 error

	if r1f, ok := ret.Get(1).(func() error); ok {
		r1 = r1f()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockRegistry) TestCompatibility(subject string, version int, schema string) (bool, error) {
	ret := _m.Called(subject, version, schema)

	var r0 bool

	if r0f, ok := ret.Get(0).(func(string, int, string) bool); ok {
		r0 = r0f(subject, version, schema)
	} else {
		r0 = ret.Get(0).(bool)
	}
	var r1 error

	if r1f, ok := ret.Get(1).(func(string, int, string) error); ok {
		r1 = r1f(subject, version, schema)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
