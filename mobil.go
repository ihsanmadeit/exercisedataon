package main

import "fmt"

type Roda interface {
	JenisBan() string
}

type BanKaret struct{}

func (b BanKaret) JenisBan() string {
	return "Ban Karet"
}

type BanKayu struct{}

func (b BanKayu) JenisBan() string {
	return "Ban Kayu"
}

type BanBesi struct{}

func (b BanBesi) JenisBan() string {
	return "Ban Besi"
}

type Pintu interface {
	Ketuk() string
	Buka() string
}

type PintuKanan struct{}

func (p PintuKanan) Ketuk() string {
	return "Knock"
}

func (p PintuKanan) Buka() string {
	return "Beep"
}

type PintuKiri struct{}

func (p PintuKiri) Ketuk() string {
	return "Beep"
}

func (p PintuKiri) Buka() string {
	return "Knock"
}

type Mobil struct {
	Roda       [4]Roda
	PintuKanan Pintu
	PintuKiri  Pintu
}

func (mobil *Mobil) GantiRoda(rodaDepanKanan, rodaDepanKiri, rodaBelakangKanan, rodaBelakangKiri Roda) {
	mobil.Roda[0] = rodaDepanKanan
	mobil.Roda[1] = rodaDepanKiri
	mobil.Roda[2] = rodaBelakangKanan
	mobil.Roda[3] = rodaBelakangKiri
}

func main() {
	mobilSaya := Mobil{}
	pintuKanan := PintuKanan{}
	pintuKiri := PintuKiri{}

	//

	mobilSaya.GantiRoda(BanKaret{}, BanKayu{}, BanBesi{}, BanKaret{})

	fmt.Println("Ketuk pintu kanan:", pintuKanan.Ketuk()) // Output: "Knock"
	fmt.Println("Buka pintu kanan:", pintuKanan.Buka())   // Output: "Beep"

	fmt.Println("Ketuk pintu kiri:", pintuKiri.Ketuk()) // Output: "Beep"
	fmt.Println("Buka pintu kiri:", pintuKiri.Buka())   // Output: "Knock"
}
