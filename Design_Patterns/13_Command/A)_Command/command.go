package main

import "fmt"

/*
	We are going to construct a bank account struct, with withdrawal and deposit operations.
*/
var overdraftlimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount, "\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) {
	if b.balance-amount >= overdraftlimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount, "\b, balance is now", b.balance)
	}
}

/*
	Now we can set up the command pattern. There are
	different ways of handling commands.

	One way is where the bank account itself handles the command.
	The other is when the command handles itself. We are going to consider this scenario.
*/

type Command interface {
	Call()
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account *BankAccount
	action  Action
	amount  int
}

func (b BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
	case Withdraw:
		b.account.Withdraw(b.amount)
	}
}

// We can also make a factory function for this command struct
func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Call()
	fmt.Println(ba)
	cmd2 := NewBankAccountCommand(&ba, Withdraw, 50)
	cmd2.Call()
	fmt.Println(ba)
}
