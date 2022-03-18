package core

import (
    "errors"
)

// MOS 6502
// Instruction Set: https://www.masswerk.at/6502/6502_instruction_set.html
// http://archive.6502.org/datasheets/rockwell_r650x_r651x.pdf

type CPU struct {
    A       uint8       //Accumulator
    X       uint8       //X Register
    Y       uint8       //Y Register 
    Stat    uint8       //Status Register
    Sp      uint8       //Stack Pointer 
    Pc      uint16      //Program Counter

    Fetchd      uint8

    AddrAbs     uint16  //Absolute Address
    AddrRel     uint16  //Relative Address 
    Opcode      uint8 
    Cycles      uint8

    _bus    *NESBus     //NES Bus
}

const ( // 6502 Flags
    FLG_C uint8 = iota // Carry Bit
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

func(cpu *CPU) SetFlag(flag uint8, value bool) error {
    if flag >= 8 {
        return errors.New("Invalid Flag")
    }
    if value == true {
        cpu.Stat |= (1 << flag) 
    } else {
        cpu.Stat &= 255 - (1 << flag)
    }
    return nil 
} 

func(cpu *CPU) GetFlag(flag uint8) (bool, error) {
    if flag >= 8 {
        return false, errors.New("Invalid Flag")
    }
    if (cpu.Stat & (1 << flag)) != 0 {
        return true, nil
    }
    return false, nil 
}

func(cpu *CPU) SIG_clock() error {
    if cpu.Cycles == 0 {
        opcode, err := cpu.read(cpu.Pc)
        if err != nil {
            return errors.New("Cannot Read OPCODE")
        }

        /* Address Mode + OPCODE */

        /* Add Cycles */

        cpu.Opcode = opcode 
    }
    cpu.Cycles--
    return nil 
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
