package core 

// MOS 6502
// Instruction Set: https://www.masswerk.at/6502/6502_instruction_set.html
// http://archive.6502.org/datasheets/rockwell_r650x_r651x.pdf

type CPU struct {
	_a	uint8	//Accumulator
	_x	uint8	//X Register
	_y	uint8	//Y Register 
	_sp uint8	//Stack Pointer 
	_pc	uint16	//Program Counter
	_bus *NESBus //NES Bus
}


func NewCPU(bus *NESBus) *CPU {
	cpu := CPU{}
	cpu._bus = bus
	return &cpu
}
