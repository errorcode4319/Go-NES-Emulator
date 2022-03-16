package core 

import (
    "errors"
)

type NESBus struct {
    _cpu	*CPU	// cpu.go
    _mem	*Memory // mem.go
}

func NewNESBus() *NESBus {
    bus := NESBus{} 
    return &bus
}

func (bus *NESBus) SetCPU(cpu *CPU) error {
    if cpu == nil {
        return errors.New("Bus SetCPU: cpu is nullptr")
    }
    bus._cpu = cpu
    return nil
}

func (bus *NESBus) SetMemory(mem *Memory) error {
    if mem == nil {
        return errors.New("Bus SetMemory: mem is nullptr") 
    }
    bus._mem = mem
    return nil 
}

func (bus *NESBus) Read(addr uint16, readOnly bool) (uint8, error) {
    return bus._mem.Read(addr, readOnly)
} 

func (bus *NESBus) Write(addr uint16, data uint8) error {
    return bus._mem.Write(addr, data)
}
