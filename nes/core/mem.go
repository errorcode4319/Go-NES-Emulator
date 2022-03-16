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

func (p_mem *Memory) Read(addr uint16, readOnly bool) (uint8, error) {
    
    return 0, nil
}

func (p_mem *Memory) Write(addr uint16, data uint8) error {


    return nil
}
