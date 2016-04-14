package cantal

import "sync/atomic"
import "unsafe"
import "fmt"


type Counter struct {
    name map[string]string
    value *uint64
}

func NewCounter(name map[string]string) *Counter {
    counter := new(Counter)
    *counter = Counter {
        name: name,
        value: new(uint64), // will be replaced later
    }
    add_value(counter)
    return counter
}

func (*Counter) GetSize() int {
    return 8
}

func (*Counter) GetType() string {
    return "counter 8"
}

func (self*Counter) Get() uint64 {
    return *self.value
}

func (self*Counter) Incr() {
    atomic.AddUint64(self.value, 1)
}

func (self*Counter) Add(amount uint64) {
    atomic.AddUint64(self.value, amount)
}

func (self*Counter) GetName() *map[string]string {
    return &self.name
}

func (self*Counter) set_pointer(ptr unsafe.Pointer) {
    self.value = (*uint64)(ptr)
}

func (self*Counter) Format(f fmt.State, c rune) {
    fmt.Fprintf(f, "%d", *self.value)
}
