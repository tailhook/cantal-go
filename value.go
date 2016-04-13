package cantal

type Value interface {
    GetSize() int
    GetType() string
    GetName() *map[string]string
}
