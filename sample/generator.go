package sample

import (
	"pcbook/pb"

	"github.com/golang/protobuf/ptypes"
)

// NewKEyboard creates a new sample keyboard
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

func NEWCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)
	min_Ghz := randomFloat64(2.0, 3.5)
	max_Ghx := randomFloat64(min_Ghz, 5.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        min_Ghz,
		MaxGhz:        max_Ghx,
	}

	return cpu
}

func NEWGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	min_Ghz := randomFloat64(1.0, 2.0)
	max_Ghz := randomFloat64(min_Ghz, 3.0)

	memory := &pb.Memory{
		Value: uint64(randomInt(2, 16)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: min_Ghz,
		MaxGhz: max_Ghz,
		Memory: memory,
	}

	return gpu
}

// NewRAM creates a new sample RAM
func NewRAM() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}

	return ram
}

func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_SDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return hdd
}

func NewScreen() *pb.Screen {

	screen := &pb.Screen{
		SizeInch:   randomFloat32(13.0, 17.0),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

func NewLaptop() *pb.Laptop {

	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id:    randomID(),
		Brand: brand,
		Name:  name,
		Cpu:   NEWCPU(),
		//RAM:      NewRAM(),
		Gpus:     []*pb.GPU{NEWGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		Price:       randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2023)),
		UpdatedAt:   ptypes.TimestampNow(),
	}

	return laptop
}
