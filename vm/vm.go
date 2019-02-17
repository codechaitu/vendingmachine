package vm

import (
	"fmt"
)

/*
 * Please note that the provided structure is not a must and can be changed
 * As Software Engineers we often have to make design decisions, if you don't agree
 * with this initial structure, feel free to change it, just don't forget to fix your tests, too!
 */

// VendingMachine offers fields and methods that represent a typical vending machine
type VendingMachine struct {
	Till             map[int]int
	Row              int
	Col              int
	CurrentInventory [4][4]Product
}

// Product structure
type Product struct {
	Name  string
	Price int
	Count int
}

func (vm *VendingMachine) Factory() {
	vm.SelectRow()
	vm.SelectCol()
	coke1 := Product{"coke1", 100, 5}
	water1 := Product{"water1", 100, 5}
	tea1 := Product{"tea1", 100, 5}
	coffee1 := Product{"coffee1", 100, 5}
	coke2 := Product{"coke2", 100, 5}
	water2 := Product{"water2", 100, 5}
	tea2 := Product{"tea2", 100, 5}
	coffee2 := Product{"coffee2", 100, 5}
	coke3 := Product{"coke3", 100, 5}
	water3 := Product{"water3", 100, 5}
	tea3 := Product{"tea3", 100, 5}
	coffee3 := Product{"coffee3", 100, 5}
	coke4 := Product{"coke4", 100, 5}
	water4 := Product{"water4", 100, 5}
	tea4 := Product{"tea4", 100, 5}
	coffee4 := Product{"coffee4", 100, 5}
	vm.CurrentInventory = [4][4]Product{
		{coke1, water1, tea1, coffee1},
		{coke2, water2, tea2, coffee2},
		{coke3, water3, tea3, coffee3},
		{coke4, water4, tea4, coffee4},
	}

}

// For selecting a product
func (vm *VendingMachine) SelectProduct(row, col int) {
	// Check if inserted amount is less than the item opted for
	if vm.Balance() < vm.CurrentInventory[row][col].Price {
		//panic("Insufficient funds") , changed to print statment because, it is easy to test.
		fmt.Println("Insufficient funds")
		// If user has inserted some money, which is not sufficient for the product, in this case we would like to return back his money.
		vm.ReturnUserInsufficientMoney()
	}

	// Check if the selected items are available in Inventory
	if vm.CurrentInventory[row][col].Count == 0 {
		//panic("there is no item available"), to test, instead of panic, print statements are used.
		fmt.Println("there is no item available")
		vm.ReturnUserInsufficientMoney()
	}
	// Return the remaining balance of user
	vm.ReturnRemainingBalance(row, col)
	// Update the inventory
	vm.CurrentInventory[row][col].Count = vm.CurrentInventory[row][col].Count - 1
}

func (vm *VendingMachine) ReturnUserInsufficientMoney() {
	// 1) Give back his money
	// Get the balance first, and give that as input to ReturnChangeDenomination, which gives back money to user
	vm.ReturnChangeDenomination(vm.Balance())

}

func (vm *VendingMachine) SelectRow() {
	fmt.Println("Enter Row: ")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "A":
		fmt.Println("You entered Row: " + input)
		vm.Row = 1
	case "B":
		fmt.Println("You entered Row: " + input)
		vm.Row = 2
	case "C":
		fmt.Println("You entered Row: " + input)
		vm.Row = 3
	case "D":
		fmt.Println("You entered Row: " + input)
		vm.Row = 4
	default:
		fmt.Println("You entered Incorrect Row")
	}

}

func (vm *VendingMachine) SelectCol() {
	fmt.Println("Enter Column: ")
	var input2 string
	fmt.Scanln(&input2)
	switch input2 {
	case "1":
		fmt.Println("You entered col: " + input2)
		vm.Col = 1
	case "2":
		fmt.Println("You entered col: " + input2)
		vm.Col = 2
	case "3":
		fmt.Println("You entered col: " + input2)
		vm.Col = 3
	case "4":
		fmt.Println("You entered Row: " + input2)
		vm.Col = 4
	default:
		fmt.Println("You entered Incorrect Column")
	}

}

// To calculate remaining amount, after product disposal
func (vm *VendingMachine) ReturnRemainingBalance(row, col int) int {
	productPrice := vm.CurrentInventory[row][col].Price
	InsertedMoney := vm.Balance()
	bal := InsertedMoney - productPrice

	// check if remaining balance is available to give back
	if bal > 0 {
		fmt.Println(vm.ReturnChangeDenomination(bal))
		return bal
	} else {
		//panic("Insufficient funds")
		fmt.Println("Insufficient funds")
		return 0
	}
}

// Return change in proper denomination,
func (vm *VendingMachine) ReturnChangeDenomination(bal int) map[int]int {
	coins := [4]int{500, 100, 50, 10}
	denomination := make(map[int]int)
	i := 0
	for bal > 0 {
		currentCoin := coins[i]
		howManyCoins := bal / currentCoin // Number of coins to take, to given denomination
		bal = bal % currentCoin
		if howManyCoins > 0 {
			denomination[currentCoin] = howManyCoins
		}
		i = i + 1
	}
	return denomination
}

// InsertCoin inserts a new coin into the Vending Machine
func (vm *VendingMachine) InsertCoin(coin int) {
	if _, ok := vm.Till[coin]; !ok {
		panic("Not a valid coin!")
	}

	vm.Till[coin]++
}

// Balance returns the currently available balance in the Vending Machine that can be used by the customer
func (vm *VendingMachine) Balance() (sum int) {
	for coin, amount := range vm.Till {
		sum += coin * amount
	}
	return sum
}

// New constructs a new Vending Machine
func New() *VendingMachine {
	return &VendingMachine{
		Till: map[int]int{
			10:  0,
			50:  0,
			100: 0,
			500: 0,
		},
	}
}
