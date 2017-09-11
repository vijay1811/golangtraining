package main

import (
	"errors"
	"io"
	"net/http"
	"testing"
)

type readCloser struct {
	data []byte
	pos  int
}

// func urlfuncMock(url string) (*http.Response, error) {
// 	return &http.Response{
// 		Body: &readCloser{
// 			data: []byte(url),
// 		},
// 	}, nil
// }

func (rc *readCloser) Read(b []byte) (int, error) {
	if rc.pos == len(rc.data) {
		return 0, io.EOF
	}
	b[0] = rc.data[rc.pos]
	rc.pos++
	return 1, nil
}
func (rc *readCloser) Close() error {
	rc.pos = len(rc.data) - 1
	return nil
}

func Test_getLenTime(t *testing.T) {
	urlfunc = func(url string) (*http.Response, error) {
		if url != "hit" {
			t.Errorf("expectd : %s, got : %s", "hit", url)
		}
		return &http.Response{
			Body: &readCloser{
				data: []byte(url),
			},
		}, nil
	}
	data, length, _, _ := getDataLenTime("hihit")
	if string(data) != "hit" {
		t.Errorf("expected data :%s, got : %s", "hit", data)
	}
	if length != 3 {
		t.Errorf("expected length of data : %v, got : %v", 3, length)
	}
}

func Test_getLenTimeFailure(t *testing.T) {
	urlfunc = func(url string) (*http.Response, error) {
		if url != "hit" {
			t.Fatalf("expectd : %s, got : %s", "hit", url)
		}
		return nil, errors.New("I am a test error")
	}
	data, length, _, _ := getDataLenTime("len")
	if length != 0 {
		t.Fatalf("expected :%v, got : %v", 3, length)
	}
	if data != nil {
		t.Fatalf("expected :%v, got : %v", nil, data)
	}
}
