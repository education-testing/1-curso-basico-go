package classes

import "fmt"

type Pc struct {
	Memory int
	Cpu    int
	Brand  string
}

func (pc *Pc) SetMemory(memory int) {
	pc.Memory = memory
}

func (pc *Pc) GetMemory() int {
	return pc.Memory
}

func (pc *Pc) SetBrand(brand string) {
	pc.Brand = brand
}

func (pc *Pc) GetBrand() string {
	return pc.Brand
}

func (pc *Pc) SetCpu(cpu int) {
	pc.Cpu = cpu
}

func (pc *Pc) GetCpu() int {
	return pc.Cpu
}

func (pc Pc) String() string {
	return fmt.Sprintf("Memory: %d, CPU: %d, brand: %s", pc.Memory, pc.Cpu, pc.Brand)
}
