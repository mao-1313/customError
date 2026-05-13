package main

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("id %d not found", e.ID)
}

type InvalidError struct {
	ID int
}

func (e *InvalidError) Error() string {
	return fmt.Sprintf("id %d is invalid", e.ID)
}

var ErrUnauthorized = errors.New("unauthorized")

func getUserByID(id int) error {
	// 形式エラーのケース
	if id == 0 {
		return fmt.Errorf("getUserByID: %w", &InvalidError{ID: id})
	}
	// 存在しないIDの場合(1~100がある前提)
	if id > 100 {
		return fmt.Errorf("getUserByID: %w", &NotFoundError{ID: id})
	}

	if id < 0 {
		return fmt.Errorf("getUserByID: %w", ErrUnauthorized)
	}
	return nil
}
