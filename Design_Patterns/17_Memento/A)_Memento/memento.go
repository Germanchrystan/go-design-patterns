package main

import "fmt"

type BankAccount struct {
	balance int
}

// Instead of returning nothing, we will return a memento
// The memento will take a snapshot of out system
func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	return &Memento{b.balance} // Preserves this new balance
}

func (b *BankAccount) Restore(m *Memento) {
	b.balance = m.Balance
}

type Memento struct {
	Balance int
}

func main() {
	ba := BankAccount{100}
	m1 := ba.Deposit(50)
	m2 := ba.Deposit(25)

	fmt.Println(ba)
	ba.Restore(m1)
	ba.Restore(m2)
	fmt.Println(ba)
}

/*
	The problem is that we do not have a memento for the initial state of the bank account.
	We could have an initializer for a bank account that also returns a memento.
*/
func NewBankAccount(balance int) (*BankAccount, *Memento) {
	return &BankAccount{balance: balance}, &Memento{balance}
}

func main2() {
	ba, m0 := NewBankAccount(100)
	m1 := ba.Deposit(50)
	m2 := ba.Deposit(25)

	fmt.Println(ba)
	ba.Restore(m2)
	ba.Restore(m1)
	fmt.Println(m0)
}
