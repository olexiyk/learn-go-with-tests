package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}

type Bitcoin uint64

func (self *Wallet) Deposit(b Bitcoin) {
	self.balance += b
}

func (self *Wallet) Withdraw(b Bitcoin) error {
	if self.balance < b {
		return errors.New("balance is not enough")
	}
	self.balance -= b
	return nil
}

func (self Wallet) Balance() Bitcoin {
	return self.balance
}

func (self Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", self)
}
