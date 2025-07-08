package main

import "testing"

func TestAbs(t *testing.T) {

	tests := []struct { // добавляем слайс тестов
		name  string
		value float64
		want  float64
	}{
		{
			name:  "simple test #1", // описываем каждый тест:
			value: -1,               // значения, которые будет принимать функция,
			want:  1,                // и ожидаемый результат
		},
		{
			name:  "zero",
			value: -0,
			want:  0,
		},
	}
	for _, test := range tests { // цикл по всем тестам
		t.Run(test.name,
			func(t *testing.T) {
				if got := Abs(test.value); got != test.want {
					t.Errorf("Abs(%v) = %v, want %v", test.value, got, test.want)
				}
			})
	}
}
