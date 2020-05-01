package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer_LoadDependencies_Success(t *testing.T) {
	testServer := StartServer("../../.env")
	assert.NotPanics(t, func() {
		testServer.LoadDependencies()
		testServer.LoadRoutes()
	})
}

func TestServer_LoadDependencies_Fail(t *testing.T) {
	assert.Panics(t, func() {
		testServer := StartServer("../notexists.env")
		testServer.LoadDependencies()
		testServer.LoadRoutes()
	})
}
