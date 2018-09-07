package command

import (
	"testing"
)

func TestChain(t *testing.T) {
	sell := new(Sell)
	buy := new(Buy)
	broker := new(Broker)
	broker.AddCommand("sell", sell)
	broker.AddCommand("buy", buy)

	broker.DoSell()
	broker.DoBuy()
}

func init() {
}
