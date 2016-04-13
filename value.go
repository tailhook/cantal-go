package cantal

import "unsafe"

type Value interface {
    GetSize() int
    GetType() string
    GetName() *map[string]string
    set_pointer(unsafe.Pointer)
}
