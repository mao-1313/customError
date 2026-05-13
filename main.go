package main

import (
	"errors"
	"fmt"
)

func main() {
	err := getUserByID(101)
	var notFound *NotFoundError
	var invalidError *InvalidError

	if errors.As(err, &notFound) {
		fmt.Println("ユーザーが見つかりませんでした")
		fmt.Println(err)
	}

	err = getUserByID(0)
	if errors.As(err, &invalidError) {
		fmt.Println("値が不正です")
		fmt.Println(err)
	}

	err = getUserByID(-1)
	if errors.Is(err, ErrUnauthorized) {
		fmt.Println("権限がありません")
		fmt.Println(err)
	}
}
