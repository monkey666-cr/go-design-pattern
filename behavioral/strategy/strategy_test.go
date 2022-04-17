package strategy

func ExampleCash_Pay() {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()

	// Output:
	// Pay $123 to Ada by cash
}

func ExampleBank_Pay() {
	payment := NewPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()

	// Output:
	// Pay $888 to Bob by bank account 0002
}
