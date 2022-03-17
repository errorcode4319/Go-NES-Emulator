package core 


type Memory struct {
    Size    uint16
    arr	    []uint8	
}

func NewMemory(size uint16) *Memory {
    mem := Memory{}
    mem.Size = size 
    mem.arr = make([]uint8, size)
    return &mem
}

func (mem *Memory) Read(addr uint16, readOnly bool) (uint8, error) {
    if 0x0000 <= addr && addr <= 0x0000 {
        return mem.arr[addr], nil 
    }
    return 0, nil
}

func (mem *Memory) Write(addr uint16, data uint8) error {
    if 0x0000 <= addr && addr <= 0x0000 {
        mem.arr[addr] = data
    }
    return nil
}
