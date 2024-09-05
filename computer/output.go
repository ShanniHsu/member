package computer

// https://refactoring.guru/design-patterns/adapter/go/example#
// https://golangbyexample.com/adapter-design-pattern-go/
func Output() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	// 如果要使用Windows就需要透過轉換才可使用
	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	// 目的是使客戶端可以使用computer中的Mac以及Windows(他的方法不同需使用適配器)
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}

/*
Client inserts Lightning connector into computer.
Lightning connector is plugged into mac machine.
Client inserts Lightning connector into computer.
Adapter converts Lightning signal to USB.
USB connector is plugged into windows machine.

*/
