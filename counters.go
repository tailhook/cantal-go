package cantal

type Counter struct {
    name map[string]string
    value *uint64
}

func NewCounter(name map[string]string) Counter {
    return Counter {
        name: name,
        value: new(uint64), // will be replaced later
    }
}
