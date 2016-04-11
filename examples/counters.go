package main
import cantal "github.com/tailhook/cantal-go"
import "fmt"

var counter = cantal.NewCounter(map[string]string{
    "metric": "Hello",
    })

func main() {
    fmt.Println("hello world")
}
