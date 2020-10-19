package web

import (
	"net/http"
	"testing"
)

type MockRender struct {
	T    *testing.T
	Name string
	Data Data
}

func (m *MockRender) Render(w http.ResponseWriter, name string, data Data) error {
	m.T.Helper()
	m.Name = name
	m.Data = data
	return nil
}
