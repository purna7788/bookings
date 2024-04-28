package main

import (
	"net/http"
	"testing"
)

func TestNosurf(t *testing.T){

	surf := NoSurf(&myHandler{})

	switch v:= surf.(type){
		case http.Handler: 
	 default: 
		t.Errorf("type is not http.handler but is %T", v)

	}
}

func TestSessionLoad(t *testing.T){

	session := SessionLoad(&myHandler{})

	switch v:= session.(type){
		case http.Handler: 
	 default: 
		t.Errorf("type is not http.handler but is %T", v)

	}
}