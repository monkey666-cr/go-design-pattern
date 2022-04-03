package bridge

func ExampleCommonMessage_SendMessage() {
	m := NewCommonMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via SMS
}

func ExampleUrgencyMessage_SendMessage() {
	m2 := NewUrgencyMessage(ViaEmail())
	m2.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via Email
}
