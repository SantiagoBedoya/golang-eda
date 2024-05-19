package usecase

import (
	"context"
	"fmt"
	"golang-eda/internal/domain/event"
)

type SendOrderEmailUseCase struct {
}

func NewSendOrderEmailUseCase() *SendOrderEmailUseCase {
	return &SendOrderEmailUseCase{}
}

func (h *SendOrderEmailUseCase) Execute(ctx context.Context, payload *event.OrderCreatedEvent) error {
	fmt.Println("--- SendOrderEmailUseCase ---")
	fmt.Printf("--- MAIL Order Created: R$ %f \n", payload.TotalPrice)
	return nil
}
