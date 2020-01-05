package singleton

import (
	"fmt"
	"testing"
)

func Test_singleton_AddOne(t *testing.T) {
	type fields struct {
		count int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &singleton{
				count: tt.fields.count,
			}
			if got := s.AddOne(); got != tt.want {
				t.Errorf("singleton.AddOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInstance(t *testing.T) {
	tests := []struct {
		name string
		want Singleton
	}{
		{name: "one"},
		{name: "two"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetInstance()
			count := got.AddOne()
			fmt.Printf("Current: %s ,count:%d \n", tt.name, count)
		})
	}
}
