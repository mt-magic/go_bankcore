package go_bank

import (
	"errors"
	"fmt"
)

// func Hello() string {
// 	return "Hey! I'm working!"
// }

/// 顧客データ
// Customer ...
type Customer struct {
    Name    string
    Address string
    Phone   string
}

/// アカウントデータ
// Account ...
type Account struct {
    Customer
    Number  int32
    Balance float64
}

/// 預金を加算
/// エラーチェック：預金加算の金額が 0 未満の場合にエラーを返す
// Deposit ...
func (a *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("the amount to deposit should be greater than zero")
    }

    a.Balance += amount
    return nil
}

/// 預金を減算
/// エラーチェック：預金減算の金額が 0 未満の場合にエラーを返す
/// エラーチェック：預金減算の金額が預金より大きい場合にエラーを返す
// Withdraw ...
func (a *Account) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("the amount to withdraw should be greater than zero")
    }

    if a.Balance < amount {
        return errors.New("the amount to withdraw should be less than the account's balance")
    }

    a.Balance -= amount
    return nil
}

/// アカウント（+顧客）情報を表示
// Statement ...
func (a *Account) Statement() string {
    return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}
