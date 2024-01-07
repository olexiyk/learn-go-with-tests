package main

import "fmt"

type Wallet struct {
	balance Bitcoin
}

type Bitcoin uint64

func (self *Wallet) Deposit(b Bitcoin) {
	self.balance += b
}

func (self *Wallet) Withdraw(b Bitcoin) {
	self.balance -= b
}

func (self Wallet) Balance() Bitcoin {
	return self.balance
}

func (self Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", self)
}
