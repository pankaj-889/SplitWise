package main

import (
	"fmt"
	"time"
)

type User struct {
	Name     string
	Email    string
	Password string
}

type Group struct {
	Id        int64
	Name      string
	CreatedBy string
	CreatedAt time.Time
}

type Expense struct {
	Id            int64
	GroupId       int64
	ExpenseAmount int64
	PaidBy        []User
	SplitBetween  []User
	Description   string
}

type Payment struct {
	Id          int64
	Amount      int64
	PaymentDate time.Time
	PayerId     int64
	PayeeId     int64
	GroupId     int64
}

type Balance struct {
	UserId     int64
	GroupId    int64
	AmountOwed int64
	AmountDue  int64
}

type Settlement struct {
	Id       int64
	Amount   int64
	FromUser int64
	ToUser   int64
	Date     time.Time
}

func main() {
	fmt.Println("Welcome to the SplitWise")
}
