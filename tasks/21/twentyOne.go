// Задача: Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// некоторый уникальный интерфейс
type DeviceConnector interface {
	Connect()
}

// один из наших классов, реализующий подключение флешки
type FlashDrive struct{}

func (fd *FlashDrive) ConnectWithUsb() {
	fmt.Println("connect flash drive with usb")
}

// еще один из наших классов, реализующий подключение карты памяти
type MemoryCard struct{}

func (mc *MemoryCard) insert() {
	fmt.Println("insert memory card")
}

// класс адаптера для флешки, который будет содержать метод Connect()
type FlashDriveAdapter struct {
	*FlashDrive
}

func (adapter *FlashDriveAdapter) Connect() {
	adapter.ConnectWithUsb()
}

// и конструктор для каста адаптера в DeviceConnector
func NewFlashDriveAdapter(fd *FlashDrive) DeviceConnector {
	return &FlashDriveAdapter{fd}
}

// адаптер для карты памяти
type MemoryCardAdapter struct {
	*MemoryCard
}

func (adapter *MemoryCardAdapter) Connect() {
	adapter.insert()
}

// конструктор для адаптера
func NewMemoryCardAdapter(mc *MemoryCard) DeviceConnector {
	return &MemoryCardAdapter{mc}
}

func main() {

	devices := []DeviceConnector{NewFlashDriveAdapter(&FlashDrive{}), NewMemoryCardAdapter(&MemoryCard{})}

	for _, device := range devices {
		device.Connect()
	}

}
