package models

import (
	"errors"
	"time"
)

type OrderSide struct string {
	const (
		Buy OrderSide = "BUY"
		Sell OrderSide = "SELL"
	)
}

type Order struct {
	ID string `json:"id"`
	UserId string `json:"user_id"`
	AssetPair string `json:"asset_pair"`
	Price int64 `json:"price"`
	Amount int64 `json:"amount"`
	Side OrderSide `json:"side"`
	CreatedAt time.Time `json:"created_at"`
}

func (o *Order) validate() error {
	if o.ID == "" || o.UserID == "" || o.AssetPair == "" || o.Price <= 0 || o.Amount <= 0 {
		return errors.New("ошибка создания ордера: неверно заполнены параметры")
	}
	if o.Side != Buy && o.Side != Sell {
		return errors.New("ошибка создания ордера: неверное направление сделки")
	}
	return nil
}