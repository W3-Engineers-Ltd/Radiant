// Copyright 2021 radiant Author. All Rights Reserved.
//

package context

import (
	"net/http"
	"testing"
)

func Test_ExtractEncoding(t *testing.T) {
	if parseEncoding(&http.Request{Header: map[string][]string{"Accept-Encoding": {"gzip,deflate"}}}) != "gzip" {
		t.Fail()
	}
	if parseEncoding(&http.Request{Header: map[string][]string{"Accept-Encoding": {"deflate,gzip"}}}) != "deflate" {
		t.Fail()
	}
	if parseEncoding(&http.Request{Header: map[string][]string{"Accept-Encoding": {"gzip;q=.5,deflate"}}}) != "deflate" {
		t.Fail()
	}
	if parseEncoding(&http.Request{Header: map[string][]string{"Accept-Encoding": {"gzip;q=.5,deflate;q=0.3"}}}) != "gzip" {
		t.Fail()
	}
	if parseEncoding(&http.Request{Header: map[string][]string{"Accept-Encoding": {"gzip;q=0,deflate"}}}) != "deflate" {
		t.Fail()
	}
	if parseEncoding(&http.Request{Header: map[string][]string{"Accept-Encoding": {"deflate;q=0.5,gzip;q=0.5,identity"}}}) != "" {
		t.Fail()
	}
	if parseEncoding(&http.Request{Header: map[string][]string{"Accept-Encoding": {"*"}}}) != "gzip" {
		t.Fail()
	}
	if parseEncoding(&http.Request{Header: map[string][]string{"Accept-Encoding": {"x,gzip,deflate"}}}) != "gzip" {
		t.Fail()
	}
	if parseEncoding(&http.Request{Header: map[string][]string{"Accept-Encoding": {"gzip,x,deflate"}}}) != "gzip" {
		t.Fail()
	}
	if parseEncoding(&http.Request{Header: map[string][]string{"Accept-Encoding": {"gzip;q=0.5,x,deflate"}}}) != "deflate" {
		t.Fail()
	}
	if parseEncoding(&http.Request{Header: map[string][]string{"Accept-Encoding": {"x"}}}) != "" {
		t.Fail()
	}
	if parseEncoding(&http.Request{Header: map[string][]string{"Accept-Encoding": {"gzip;q=0.5,x;q=0.8"}}}) != "gzip" {
		t.Fail()
	}
}
