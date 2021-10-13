package main

import "testing"

func TestMaxBoleeMenee(t *testing.T) {
	// Тестовые данные
	testTable := []struct {
		numbers []int
		answer  int
	}{
		{
			[]int{1, 2, 33, 4, 5, 6},
			33,
		},
		{
			[]int{0, -2, -33, -4, -5, -6},
			0,
		},
		{
			[]int{},
			0,
		},
	}

	// настройка тестовых данных
	for _, testCase := range testTable {
		result := Max(testCase.numbers)

		t.Logf("Testing Max(%v), want %d,  result %d\n",
			testCase.numbers, testCase.answer, result)

		// проверка результатов
		if result != testCase.answer {
			t.Errorf("Unexpected result: want: %d, got: %d",
				testCase.answer, result)
		}
	}
}
