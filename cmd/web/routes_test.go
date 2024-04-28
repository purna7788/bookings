package main

import (
	"net/http"
	"testing"

	"github.com/purna7788/bookings/cmd/internal/config"
)

func TestMux(t *testing.T){
	var app *config.AppConfig

	h := routes(app)

	switch v:= h.(type) {
		case http.Handler:
	default:
		t.Errorf("type is not of type http.Handler. It is of type %T",v)
	}


}