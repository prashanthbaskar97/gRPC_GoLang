package sample

import (
	"math/rand"
	"pcbook/pb"
	"time"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY

	case 2:
		return pb.Keyboard_AZERTY
	}

	return pb.Keyboard_QWERTZ
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {

	return randomStringFromSet("Intel", "AMD")
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo", "Asus", "HP")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet("i3", "i5", "i7", "i9")
	}

	return randomStringFromSet("Ryzen 3", "Ryzen 5", "Ryzen 7", "Ryzen 9")

}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet("GTX 1050", "GTX 1060", "GTX 1070", "GTX 1080")
	}
	return randomStringFromSet("RX 550", "RX 560", "RX 570", "RX 580")
}

func randomLaptopName(brand string) string {
	if brand == "Apple" {
		return randomStringFromSet("MacBook Air", "MacBook Pro")
	}
	if brand == "Dell" {
		return randomStringFromSet("XPS 13", "XPS 15")
	}
	if brand == "Lenovo" {
		return randomStringFromSet("ThinkPad X1 Carbon", "ThinkPad T14")
	}
	if brand == "Asus" {
		return randomStringFromSet("ZenBook 13", "ROG Zephyrus G14")
	}

	return randomStringFromSet("Spectre x360", "Pavilion 15")

}
func randomInt(Min int, Max int) int {
	return rand.Intn(Max-Min+1) + Min
}

func randomFloat64(Min float64, Max float64) float64 {
	return Min + rand.Float64()*(Max-Min)
}

func randomFloat32(Min float32, Max float32) float32 {
	return Min + rand.Float32()*(Max-Min)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9

	return &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
}

func randomID() string {
	return uuid.New().String()
}
