package usecase

import (
	"context"
	"fmt"
	"trading-platform/internal/domain/models"
	"trading-platform/internal/domain/ports"
)

//оркестрация сценариев работы с ордерами
type OrderProcessor struct {
	//зависимость только от интерфейсов
	//нам всё равно, какая будет бд
	orderRepo ports.OrderRepository
}

//конструктор для сборки usecase
func newOrderProcessor (repo ports.OrderREpository) *OrderProcessor {
	return &OrderProcessor{
		orderRepo: repo,
	}
}

func (p *orderProcessor) createOrder