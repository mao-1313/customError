package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestNotFoundError(t *testing.T) {
	err := getUserByID(101)
	var notFound *NotFoundError
	if errors.As(err, &notFound) {
		fmt.Println("ユーザーが見つかりませんでした")
		fmt.Println(err)
	}
}

func TestInvalidError(t *testing.T) {

	var invalidError *InvalidError
	err := getUserByID(0)
	if errors.As(err, &invalidError) {
		fmt.Println("値が不正です")
		fmt.Println(err)
	}
}

func TestUnauthorizedError(t *testing.T) {
	err := getUserByID(-1)
	if errors.Is(err, ErrUnauthorized) {
		fmt.Println("権限がありません")
		fmt.Println(err)
	}
}
