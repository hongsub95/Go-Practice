package accounts

import "errors"

//Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("잔고금액이 부족합니다")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// 포인터(*)를 가져와야 한다. 복사본이 아닌 실제 account를 가져와야 해서
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

// 포인터(*)를 쓴 이유는 위와같음
// nil은 null,none 이라고 생각하면 된다

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
