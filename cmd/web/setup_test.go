package main

import (
	"net/http"
	"os"
	"testing"
)

type myHandler struct {
}

func SetUpMain(m *testing.M) {
	os.Exit(m.Run())
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { }

