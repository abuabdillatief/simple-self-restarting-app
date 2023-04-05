package main

import "fmt"

type Product struct {
	Name       string
	Price      int
	PriceFloat float32
}

type User struct {
	Name        string
	Wallet      int
	WalletFloat float32
}

func (u *User) Purchase(product Product, disc int) {
	fmt.Println("\n==========int==========")
	u.Wallet -= (product.Price * (100 - disc) / 100)
	fmt.Printf("%s just purchased %s for %d", u.Name, product.Name, product.Price/100*(100-disc)/100)
	fmt.Printf("Remaining balance: %d", u.Wallet/100)
	fmt.Println("\n=======================")
}

func (u *User) PurchaseWithFloat(product Product, disc int) {
	fmt.Println("\n==========float==========")
	u.WalletFloat -= (product.PriceFloat * float32((100-disc)/100))
	fmt.Printf("%s just purchased %s for %f", u.Name, product.Name, product.PriceFloat*float32((100-disc)/100))
	fmt.Printf("Remaining balance: %f", u.WalletFloat)
	fmt.Println("\n========================")
}

func HandlePriceWithInt() {
	user := User{
		Name:        "User 1",
		Wallet:      100000000,
		WalletFloat: 100000000,
	}

	keyboard := Product{
		Name:       "Keyboard 1",
		Price:      75000000,
		PriceFloat: 75000000,
	}

	user.Purchase(keyboard, 15)
}

func HandlePriceWithFloat() {
	user := User{
		Name:   "User 1",
		Wallet: 100000000,
	}

	keyboard := Product{
		Name:  "Keyboard 1",
		Price: 75000000,
	}
	user.PurchaseWithFloat(keyboard, 15)
}
