package accounts

import (
	"errors"
	"fmt"
)

// Account struct
// private - 소문자
// public - 대문자
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw! you have no money")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit : receiver 를 이용해 Account 의 Method 생성 가능
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// GetBalance of your account
func (a Account) GetBalance() int {
	return a.balance
}

// Withdraw your account
// error type 에는 error 와 nil 이 있다
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// GetOwner of the account
func (a *Account) GetOwner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.GetOwner(), "'s account.\nHas: ", a.GetBalance())
}
