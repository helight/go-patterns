package command

import "fmt"

type Command interface {
	Do_command()
}

type Sell struct {
}
func (s *Sell) Do_command() {
	fmt.Println("do Sell command")
}

type Buy struct {
}
func (b *Buy) Do_command() {
	fmt.Println("do buy command")
}

type Broker struct {
	Sell Command
	Buy Command
}

func (b *Broker) AddCommand(ctype string, command Command) {
	if "sell" == ctype {
		b.Sell = command
	} else if "buy" == ctype {
		b.Buy = command
	} else {
		fmt.Println("add wrong command")
	}
}

func (b *Broker) DoSell() {
	b.Sell.Do_command()
}

func (b *Broker) DoBuy() {
	b.Buy.Do_command()
}

func main() {
}
