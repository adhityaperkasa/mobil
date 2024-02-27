package main

import (
	"fmt"
)

// Wheel interface untuk merepresentasikan roda pada mobil
type Wheel interface {
	GetType() string
}

// BanKaret struct untuk roda dengan ban karet
type BanKaret struct{}

func (bk BanKaret) GetType() string {
	return "Ban Karet"
}

// BanKayu struct untuk roda dengan ban kayu1234
type BanKayu struct{}

func (bk BanKayu) GetType() string {
	return "Ban Kayu"
}

// BanBesi struct untuk roda dengan ban besi
type BanBesi struct{}

func (bb BanBesi) GetType() string {
	return "Ban Besi"
}

// Pintu struct untuk merepresentasikan pintu pada mobil
type Pintu struct {
	SoundOnKnock string
	SoundOnOpen  string
}

// Mobil struct untuk merepresentasikan mobil
type Mobil struct {
	RodaDepan    Wheel
	RodaBelakang Wheel
	PintuKanan   Pintu
	PintuKiri    Pintu
}

func main() {
	// Membuat objek roda dengan berbagai jenis ban
	rodaDepan := BanKaret{}
	rodaBelakang := BanBesi{}

	// Membuat objek pintu dengan bunyi yang berbeda
	pintuKanan := Pintu{SoundOnKnock: "Knock", SoundOnOpen: "beep"}
	pintuKiri := Pintu{SoundOnKnock: "beep", SoundOnOpen: "Knock"}

	// Membuat objek mobil
	mobil := Mobil{
		RodaDepan:    rodaDepan,
		RodaBelakang: rodaBelakang,
		PintuKanan:   pintuKanan,
		PintuKiri:    pintuKiri,
	}

	// Penggantian roda
	mobil.RodaDepan = BanKayu{}
	mobil.RodaBelakang = BanKaret{}

	// Output informasi roda
	fmt.Println("Roda Depan:", mobil.RodaDepan.GetType())
	fmt.Println("Roda Belakang:", mobil.RodaBelakang.GetType())

	// Menggunakan pintu kanan
	fmt.Println("Ketuk Pintu Kanan:", mobil.PintuKanan.SoundOnKnock)
	fmt.Println("Buka Pintu Kanan:", mobil.PintuKanan.SoundOnOpen)

	// Menggunakan pintu kiri
	fmt.Println("Ketuk Pintu Kiri:", mobil.PintuKiri.SoundOnKnock)
	fmt.Println("Buka Pintu Kiri:", mobil.PintuKiri.SoundOnOpen)
}
