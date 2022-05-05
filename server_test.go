package main

import (
	"net/http"
	"strings"
	"testing"
)

func Test_getPort(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPort(); got != tt.want {
				t.Errorf("getPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_echo(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			echo(tt.args.w, tt.args.r)
		})
	}
}
func Test_echoEur(t *testing.T) {
	got := echoEur()
	want := "Actual price of one € is "
	if !strings.HasPrefix(got, want) {
		t.Errorf("echoEur() = %v, want %v", got, want)
	}
}

func Test_echoName(t *testing.T) {
	got := echoName()
	want := "Chatbots name is "
	if !strings.HasPrefix(got, want) {
		t.Errorf("echoName() = %v, want %v", got, want)
	}
}

func Test_echoTime(t *testing.T) {
	got := echoTime()
	want := "Actual server time is "
	if !strings.HasPrefix(got, want) {
		t.Errorf("echoTime() = %v, want %v", got, want)
	}
}

func Test_home(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			home(tt.args.w, tt.args.r)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
