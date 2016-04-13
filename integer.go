package cantal

import "sync/atomic"
import "unsafe"


type Integer struct {
    name map[string]string
    value *int64
}

func NewInteger(name map[string]string) *Integer {
    counter := new(Integer)
    *counter = Integer {
        name: name,
        value: new(int64), // will be replaced later
    }
    add_value(counter)
    return counter
}

func (*Integer) GetSize() int {
    return 8
}

func (*Integer) GetType() string {
    return "counter"
}

func (self*Integer) Get() int64 {
    return *self.value
}

func (self*Integer) Set(value int64) {
    // This isn't strictly needed, at least for x86_64
    // But we think it makes code more clear
    atomic.StoreInt64(self.value, value)
}

func (self*Integer) GetName() *map[string]string {
    return &self.name
}

func (self*Integer) set_pointer(ptr unsafe.Pointer) {
    self.value = (*int64)(ptr)
}
