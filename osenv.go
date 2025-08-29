package stdutil

import "os"

type Getenv func(string) string

type OSEnv struct{}

func New() *OSEnv {
	return &OSEnv{}
}

func (o *OSEnv) Getenv(k string) string {
	return os.Getenv(k)
}

type MockOSEnv struct {
	m map[string]string
}

func NewMock(m map[string]string) *MockOSEnv {
	return &MockOSEnv{
		m: m,
	}
}

func (o *MockOSEnv) Getenv(k string) string {
	return o.m[k]
}
