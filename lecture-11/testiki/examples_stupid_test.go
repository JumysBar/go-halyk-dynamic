package main

import "testing"

func TestMaxStupid(t *testing.T) {
	// Тестовые данные
	numbers := []int{1, 2, 3, 4, 15, 6, 7, 8}
	answer := 15

	// настройка тестовых данных
	result := Max(numbers)

	// проверка результатов
	if result != answer {
		t.Errorf("Unexpected result: want: %d, got: %d",
			answer, result)
	}
}

func TestMaxIndexStupid(t *testing.T) {
	// Тестовые данные
	numbers := []int{1, 2, 3, 4, 15, 6, 7, 8}
	answer := 4

	// настройка тестовых данных
	result := MaxIndex(numbers)

	// проверка результатов
	if result != answer {
		t.Errorf("Unexpected result: want: %d, got: %d",
			answer, result)
	}
}
