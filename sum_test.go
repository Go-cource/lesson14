package main

import "testing"

func TestSum(t *testing.T) {
	if sum := Sum(1, 2); sum != 3 {
		t.Errorf("Я ожидал %v,  а получил %v", 3, sum)
	}
}
