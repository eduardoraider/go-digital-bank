package main

import (
	"bank/accounts"
	"bank/holders"
	"fmt"
)

func PayBanlSlip(account verifyAccount, value float64) {
	account.WithdrawValue(value)
}

type verifyAccount interface {
	WithdrawValue(value float64) string
}

func main() {

	holderEduardo := holders.Holder{
		Name:       "Eduardo",
		CPF:        "123.456.789.01",
		Occupation: "Developer",
	}

	accountEduardo := accounts.CheckingAccount{
		AccountHolder: holderEduardo,
		BranchNumber:  123,
		AccountNumber: 123456,
	}

	accountEduardo.DepositValue(700)

	fmt.Println("Eduardo Account Data: ", accountEduardo, "Balance", accountEduardo.GetBalance())

	accountSavingTaisa := accounts.SavingAccount{}
	accountSavingTaisa.DepositValue(500)
	accountSavingTaisa.WithdrawValue(300)

	fmt.Println("Taisa Saving Account - Balance", accountSavingTaisa.GetBalance())

	PayBanlSlip(&accountSavingTaisa, 60)

	fmt.Println("Taisa Saving Account - Balance", accountSavingTaisa.GetBalance())

	status, value := accountEduardo.DepositValue(200)
	fmt.Println(status, "- Eduardo Account Balance", value)

	accountTaisa := accounts.CheckingAccount{}

	statusOperation := accountEduardo.Transfer(500, &accountTaisa)
	fmt.Println("Eduardo transfers 500 to Taisa")

	fmt.Println("Operation status:", statusOperation)
	fmt.Println("Eduardo Account Balance", accountEduardo.GetBalance())
	fmt.Println("Taisa Account Balance", accountTaisa.GetBalance())
}
