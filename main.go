package main

import (
    "fmt"
    "go-nes-emulator/nes/core"
)

func main() {
    fmt.Println("Go NES Emulator")
    nes_bus := core.NewNESBus()
    nes_mem := core.NewMemory(2048)
    nes_bus.SetMemory(nes_mem)
    nes_cpu := core.NewCPU(nes_bus)
    nes_bus.SetCPU(nes_cpu)
    nes_bus.Write(0, 0) 
}
