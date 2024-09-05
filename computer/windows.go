package computer

import "fmt"

// Unknown service

type Windows struct{}

// 與Mac差異在於InsertIntoLightningPort與insertIntoLightningPort區別
// 即是不同的方法，透過windowsAdapter轉換後即可使用
func (w *Windows) insertIntoLightningPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
