package vm_test

import (
	. "github.com/codechrysalis/go.vending-machine/vm"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vending Machine", func() {
	var machine *VendingMachine

	BeforeEach(func() {
		// Setup
		machine = New()
		machine.Factory()
	})

	Describe("Inserting Coins", func() {
		Context("After inserting 1 500 Yen coin", func() {
			BeforeEach(func() {
				// Execute
				machine.InsertCoin(500)
			})

			It("Should have a balance of 500", func() {
				// Assert
				Expect(machine.Balance()).To(Equal(500))
			})

			Specify("The till should have 1 500 Yen Coin", func() {
				// Assert
				Expect(machine.Till).To(Equal(map[int]int{
					10:  0,
					50:  0,
					100: 0,
					500: 1,
				}))
			})
		})

	})
	Describe("Inventory", func() {
		Context("After removing coke", func() {
			BeforeEach(func() {
				// Excute
				machine.InsertCoin(500)
				machine.SelectProduct(1, 1)
			})
			It("Should have one less coke", func() {
				// Assert
				Expect(machine.CurrentInventory[1][1].Count).To(Equal(4))

			})

		})

	})

	Describe("Calculate remaining balance", func() {
		Context("After dispensing a coke", func() {
			BeforeEach(func() {
				// Execute
				machine.InsertCoin(500)
				machine.SelectProduct(1, 1)
			})
			It("Change must be returned of 400Yen", func() {
				// Assert
				Expect(machine.ReturnRemainingBalance(1, 1)).To(Equal(400))

			})

		})
	})

	Describe("Calculate denomination ", func() {
		Context("After dispensing a coke, should return denomination", func() {
			BeforeEach(func() {
				// Execute
				machine.InsertCoin(260)
			})
		})

		It("Should return highest coin denoimations possible", func() {
			// Assert
			Expect(machine.ReturnChangeDenomination(160)).To(Equal(map[int]int{

				100: 1,
				50:  1,
				10:  1,
			}))

		})

	})

	Describe("Situation when money is not sufficient for any product ", func() {
		Context("User inserted one 10 yen coin, all products are more than 10 yen", func() {
			BeforeEach(func() {
				// Execute
				machine.InsertCoin(10)
				machine.SelectProduct(1, 1)
			})
			It("10 yen must be returned back to user", func() {
				//machine.InsertCoin(10)
				Expect(machine.ReturnChangeDenomination(10)).To(Equal(map[int]int{
					10: 1,
				}))
			})
		})
	})

	Describe("Happy Path ", func() {
		Context("Insert, Select Product, Return Change", func() {

			It("Should return highest coin denoimations possible", func() {
				// Assert
				machine.InsertCoin(500)
				machine.SelectProduct(1, 1)
				Expect(machine.ReturnChangeDenomination(400)).To(Equal(map[int]int{
					100: 4,
				}))

			})

		})
	})

})
