package main

import "fmt"

/*
	Let's go back to the previous example.
	So one of the features we can implement under the current command paradigm is the
	undo functionality.

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

// First of all, we need to add the undo method to the command interface
type Command interface {
	Call()
	Undo()
}

/*
	In addition, one problem we will have is that some of the commands are going to fail.
	for instance, when we withdraw a certain amount, the action can be performed only if the
	decreased amount is greater than or equal to the overdraft. Otherwise, the command doesn't apply.
	If the command doen't apply, that means we shouldn't be able to undo it either, because that will
	leave the system in an unpredictable state.
*/
func (b *BankAccount) WithdrawCorrected(amount int) bool { // Now returning bool
	if b.balance-amount >= overdraftlimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount, "\b, balance is now", b.balance)
		return true
	}
	return false
}

// We can use this bool whenever we actually perform the action.

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account  *BankAccount
	action   Action
	amount   int
	succeded bool // This will store the return value of the Withdraw method
}

func (b BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeded = true // In this example, Deposit will always succeed.
	case Withdraw:
		b.succeded = b.account.WithdrawCorrected(b.amount)
	}
}

// Now we can implement the undo method
func (b *BankAccountCommand) Undo() {
	// we can only undo the command if it succeeded
	if !b.succeded {
		return
	}
	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount) /* This isn't strictly correct.
		We cannot consider these two operations symmetrically, but this is just a simple example.
		*/
	case Withdraw:
		b.account.Deposit(b.amount)
	}
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Call()
	fmt.Println(ba)
	cmd2 := NewBankAccountCommand(&ba, Withdraw, 25)
	cmd2.Call()
	fmt.Println(ba)
	cmd2.Undo()
	fmt.Println(ba)
}
