package accounts

import "bank/holders"

type SavingAccount struct {
	AccountHolder                          holders.Holder
	BranchNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingAccount) WithdrawValue(value float64) string {
	allowWithdraw := value <= c.balance

	if allowWithdraw {
		c.balance -= value
		return "Successful withdrawal"
	} else {
		return "Insufficient funds"
	}
}

func (c *SavingAccount) DepositValue(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Deposit made successfully", c.balance
	} else {
		return "The deposit amount is less than zero", c.balance
	}
}

func (c *SavingAccount) GetBalance() float64 {
	return c.balance
}
