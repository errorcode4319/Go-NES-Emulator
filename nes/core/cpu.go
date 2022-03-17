package core 

// MOS 6502
// Instruction Set: https://www.masswerk.at/6502/6502_instruction_set.html
// http://archive.6502.org/datasheets/rockwell_r650x_r651x.pdf

type CPU struct {
    A      uint8       //Accumulator
    X      uint8       //X Register
    Y      uint8       //Y Register 
    Sp     uint8       //Stack Pointer 
    Pc     uint16      //Program Counter

    Fetchd      uint8

    AddrAbs     uint16  //Absolute Address
    AddrRel     uint16  //Relative Address 
    Opcode      uint8 
    Cycles      uint8

    _bus    *NESBus     //NES Bus
}

const ( // 6502 Flags
    FLG_C uint8 = 1 << iota // Carry Bit
    FLG_Z   // zero
    FLG_I   // disable interrupts
    FLG_D   // decimal Mode
    FLG_B   // break
    FLG_U   // unused
    FLG_V   // overflow
    FLG_N   // negative 
)

func NewCPU(bus *NESBus) *CPU {
    cpu := CPU{}
    cpu._bus = bus
    return &cpu
}

func(cpu *CPU) SIG_clock() {
    
}

func(cpu *CPU) SIG_reset() {

}

func(cpu *CPU) SIG_irq() {

}

func(cpu *CPU) SIG_nmi() {

}

// private, read memory
func(cpu *CPU) read(addr uint16) (uint8, error) {
    return cpu._bus.Read(addr, false)
}
// private, write memory
func(cpu *CPU) write(addr uint16, data uint8) error {
    return cpu._bus.Write(addr, data) 
}
