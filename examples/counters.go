package main
import cantal "github.com/tailhook/cantal-go"
import "fmt"
import "time"
import "math/rand"

var counter = cantal.NewCounter(map[string]string{
    "metric": "dots_printed",
    })
var integer = cantal.NewInteger(map[string]string{
    "metric": "random_value",
    })

func main() {
    cantal.Start()
    defer cantal.Clean()

    for true {
        counter.Incr()
        integer.Set(10 + rand.Int63() % 100)
        time.Sleep(100*time.Millisecond)
        fmt.Print(".")
    }
}
