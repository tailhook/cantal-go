package main
import cantal "github.com/tailhook/cantal-go"
import "fmt"

var counter = cantal.NewCounter(map[string]string{
    "metric": "Hello",
    })

func main() {
    cantal.Start()
    defer cantal.Clean()
    fmt.Println("hello world")
}
