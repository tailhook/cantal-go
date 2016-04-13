package cantal

import "io"
import "fmt"

func Summary(out io.Writer) error {
    var err error = nil
    for _, item := range active_counters {
        name := (*item).GetName()
        // Groups and state are sent from cantal to carbon (graphite).
        // So we add them to the summary
        group := (*name)["group"]
        state := (*name)["state"]
        metric := (*name)["metric"]
        if metric != "" {
            if group != "" {
                _, err = fmt.Printf("%s.%s: %v\n", group, metric, *item)
            } else if state != "" {
                _, err = fmt.Printf("%s.%s: %v\n", state, metric, *item)
            }
            if err != nil {
                return err
            }
        }
    }
    return nil
}
