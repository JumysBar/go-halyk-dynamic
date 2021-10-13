package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
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

		assert.Equal(t, testCase.answer, result, fmt.Sprintf("Unexpected result: want: %d, got: %d",
			testCase.answer, result))
	}
}

func TestMaxWithError(t *testing.T) {
	testTable := []struct {
		numbers []int
		answer  int
		isError bool
		err     error
	}{
		{
			[]int{1, 2, 33, 4, 5, 6},
			33,
			false,
			nil,
		},
		{
			[]int{0, -2, -33, -4, -5, -6},
			0,
			true,
			fmt.Errorf("got negative number"),
		},
		{
			[]int{},
			0,
			true,
			fmt.Errorf("slice must be non empty"),
		},
	}
	for _, testCase := range testTable {
		result, err := MaxWithError(testCase.numbers)
		if testCase.isError && err == nil {
			t.Errorf("[%+v] ждали ошибку, но не пулучили :/", testCase)
		}
		if !testCase.isError && err != nil {
			t.Errorf("[%+v] НЕ ждали ошибку, но пулучили :/ %v", testCase, err)
		}
		if !testCase.isError && err == nil {
			t.Logf("Testing MaxWithError(%v), want %d,  result %d\n",
				testCase.numbers, testCase.answer, result)

			assert.Equal(t, testCase.answer, result, fmt.Sprintf("Unexpected result: want: %d, got: %d",
				testCase.answer, result))
		}
	}
}
