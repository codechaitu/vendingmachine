# TDD Vending Machine

Practice pairing and test-driven development skills as you create a vending machine.

## Objectives

In this task you will:

- Use test-driven development to create a state machine
- Practice writing tests before code
- Practice using the SEAT pattern for tests
- Practice effective pairing techniques

## Setup

Clone this repository to `$GOPATH/src/github.com/codechrysalis/go.vending-machine`

We will be using [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](https://onsi.github.io/gomega/#making-assertions) in order to write our tests in a BDD style.

To start off you will have to `go get` the two:

```bash
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega/...
```

this will install the `ginkgo` executable under `$GOPATH/bin`, so make sure that folder is in your `$PATH`.

Ginkgo hooks into `go test`, so you can use `go test ./vm` to run the tests.
But realistically you will want to use `ginkgo vm` instead to execute your tests, as it provides a nice output.

After you have installed ginkgo and gomega, go ahead and run `ginkgo vm`, you should be getting output similar to this:

```
Running Suite: Vending Machine Suite
====================================
Random Seed: 1548590317
Will run 2 of 2 specs

++
Ran 2 of 2 Specs in 0.007 seconds
SUCCESS! -- 2 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 1.8919994s
Test Suite Passed
```

As you can see, two tests (and the implementation for them) have been pre-written for you. Look at them and use them as a guideline for how to write good tests.

Note that if you don't like a design decision made in the initial tests, you are free to change the test and implementations.

## User story

- As a shopper,
- I would like to have a vending machine
- So that I can buy goods efficiently.

## State

This is a suggestion for what the state could look like, as a Software Engineer you will have to decide if this is enough to fullfill the User Story. For example you could decide that the vending machine should have a secondary till for coin storage. Decide as a pair what you need.

- balance - the amount of money currently inserted but not yet used
- till - a map of coins and counts in the machine
- selectedRow - stores the row selected, if any
- selectedColumn - stores the selected column, if any
- inventory - stores all products that are in the vending machine

## Behavior

**Please note that these are suggestions for methods, you will probably have an easier time adding more than these methods to make testing easier. Remember: Single Responsibility Principle, a method should only do 1 thing.**

- `InsertCoin(coin)` - to put a coin in the machine
- `ReturnChange()` - resets the balance to 0 (emptying the till) and returns a `map[int]int` that represents all the coins dispensed
- `PressButton('A'-'D')` - select a row
- `PressButton('1'-'4')` - select a column

## Instructions

Model the products as a 4x4 array of arrays containing your favorite products, with a name and price, eg:

```go
// Note again that this is a suggestion, you might also want to put this into it's own file
type Product struct {
  Name string
  Price int
  Count int
}

plumbus := &Product{Name: "Plumbus", Price: 350, Count: 5};
coffee := &Product{Name: "Tully's", Price: 250, Count: 7};

vm.Inventory = [4][4]*Product{
  {plumbus, coffee, ..., ...},
  {..., ..., ..., ...},
  {..., ..., ..., ...},
  {..., ..., ..., ...},
}
```

The happy path should proceed as follows:

1.  Insert coins
1.  Select a row
1.  Select a column
1.  Dispense the product (just return the product name)
1.  Update the inventory
1.  Dispense change

For 4. and 6., depending on your structure, this could be done by returning multiple values (product name, change) to make your testing easier.

## Acceptance criteria

1.  _Given_ that the balance is zero, _when_ a coin is inserted, _then_ the balance should rise _and_ types of coins should be stored
1.  _Given_ that no row is selected, _when_ a row is selected the letter should be saved and printed to the console
1.  _Given_ that a row is selected, _when_ there is sufficiant balance and inventory and a column is selected
    1.  _then_ the row and column should be logged to the console
    1.  _and_ a message should be logged stating "Here is your [item name]"
    1.  _and_ the item inventory should decrease by 1
    1.  _and_ the item name and correct change should be returned from the method
1.  _Given_ that a row and column are selected, _when_ there is no inventory at that column, _then_ a panic should be triggered.
1.  _Given_ that a row and column are selected, _when_ the balance is insufficant to purchace the selected item, _then_ a panic should be triggered
1.  _Given_ that the program has just started, _when_ the balance is read, _then_ it should read zero

**_Please note: you must track both the types of coins and number of coins to compute the balance and return change_**

**_Hint: do not try to test what is being logged to console, instead think about how to structure your methods and have them return values for things such as returned change or dispensed item, this then allows you to test the return value_**

## Unhappy paths

At minimum, you must implement tests for all the acceptance criteria, however you should not limit yourself to only these tests. Write as many tests as needed.

As a good TDD developer, when you find an edge case outside the specs, you should ask your PM (instructor) what the desired functionality should be.

## Pairing

Take this opportunity to practice _ping-pong pairing_:

1.  The pair on the left writes a failing test
1.  The pair on the right writes the minimum code to make it pass
1.  Righty then writes the next test
1.  Lefty gets the test to pass

Don't forget to refactor, and only do so when green!
