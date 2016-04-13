package main
import cantal "github.com/tailhook/cantal-go"
import "fmt"
import "os"
import "os/signal"
import "syscall"
import "time"
import "math/rand"

var counter = cantal.NewCounter(map[string]string{
    "group": "example",
    "metric": "dots_printed",
    })
var integer = cantal.NewInteger(map[string]string{
    "group": "example",
    "metric": "random_value",
    })

func set_interrupt_handler() {
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

    go func() {
        for _ = range ch {
            fmt.Println("") // Print newline after ^C
            cantal.Summary(os.Stdout)
            cantal.Clean()
            os.Exit(0)
        }
    }()
}

func main() {
    set_interrupt_handler()
    cantal.Start()

    for true {
        counter.Incr()
        integer.Set(10 + rand.Int63() % 100)
        time.Sleep(100*time.Millisecond)
        fmt.Print(".")
    }
}
