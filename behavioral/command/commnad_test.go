package command

func ExampleCommand() {
	mb := &MotherBoard{}
	startCommand := NewStartCommand(mb)
	rebootCommand := NewRebootCommand(mb)

	box1 := NewBox(startCommand, rebootCommand)
	box1.PresButton1()
	box1.PresButton2()

	box2 := NewBox(rebootCommand, startCommand)
	box2.PresButton1()
	box2.PresButton2()

	// Output:
	// system starting
	// system rebooting
	// system rebooting
	// system starting
}
