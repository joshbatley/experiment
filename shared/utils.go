package utils

import "math/rand"

func FindIndex[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func Find[T any](slice []*T, comparable func(*T) bool) *T {
	for _, v := range slice {
		if comparable(v) {
			return v
		}
	}
	return nil
}

func GetRandomItem[T any](slice []T) T {
	randomIdx := rand.Intn(len(slice))
	return slice[randomIdx]
}
