package cantal

import "io"
import "fmt"

func Summary(out io.Writer) {
    for _, item := range active_counters {
        name := (*item).GetName()
        // Groups and state are sent from cantal to carbon (graphite).
        // So we add them to the summary
        group := (*name)["group"]
        state := (*name)["state"]
        metric := (*name)["metric"]
        if metric != "" {
            if group != "" {
                fmt.Printf("%s.%s: %v\n", group, metric, *item)
            } else if state != "" {
                fmt.Printf("%s.%s: %v\n", state, metric, *item)
            }
        }
    }
}
