package main
import cantal "github.com/tailhook/cantal-go"
import "fmt"
import "time"

var counter = cantal.NewCounter(map[string]string{
    "metric": "dots_printed",
    })

func main() {
    cantal.Start()
    defer cantal.Clean()

    for true {
        counter.Incr()
        time.Sleep(100*time.Millisecond)
        fmt.Print(".")
    }
}
