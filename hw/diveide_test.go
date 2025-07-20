package main

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr bool
	}{
		{"div by zero", 10, 0, 0, true},
		{"noraml test", 10, 2, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Got error %v", err)
				return
			}
			if res != tt.want {
				t.Errorf("Got fail value %v", res)
			}
		})
	}
}
