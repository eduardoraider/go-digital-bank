package accounts

import (
	"bank/holders"
)

type CheckingAccount struct {
	AccountHolder               holders.Holder
	BranchNumber, AccountNumber int
	balance                     float64
}

func (c *CheckingAccount) WithdrawValue(value float64) string {
	allowWithdraw := value <= c.balance

	if allowWithdraw {
		c.balance -= value
		return "Successful withdrawal"
	} else {
		return "Insufficient funds"
	}
}

func (c *CheckingAccount) DepositValue(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Deposit made successfully", c.balance
	} else {
		return "The deposit amount is less than zero", c.balance
	}
}

func (c *CheckingAccount) Transfer(value float64, destAccount *CheckingAccount) bool {
	if value < c.balance && value > 0 {
		c.balance -= value
		destAccount.DepositValue(value)
		return true
	} else {
		return false
	}
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}
