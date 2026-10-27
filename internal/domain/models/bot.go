package models

import (
	"time"
	"errors"
)



type BotConfig struct {
	AssetPair string `json:"asset_pair"`
	StrategyType string `json:"strategy_type"`
	TargetBuyPrice int64 `json:"target_buy_price"`
	TradeAmout int64 `json:"trade_amout"`
}

type Bot struct {
	ID string `json:"id"`
	UserId string `json:"user_id"`
	IsActive bool `json:"is_active"`
	Config BotConfig `json:"config"`
	CreatedAt time.Time `json:"created_at"`
}



func (b *Bot) validate() error {
	if b.ID == "" || b.UserID == "" || b.Config.AssetPair == "" || b.Config.TradeAmount <= 0 {
		return errors.New("ошибка регистрации бота: неверно заполнены параметры")
	}
	return nil
}