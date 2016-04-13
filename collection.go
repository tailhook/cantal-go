package cantal

import "bytes"
import "fmt"
import "os"
import "log"
import "encoding/json"

var active_counters = []*Value{}
var meta_path = ""
var data_path = ""

func add_value(value Value) {
    active_counters = append(active_counters, &value)
}

func Start() {
    offset := 0
    scheme := bytes.Buffer{}
    for _, cnt := range active_counters {
        size := (*cnt).GetSize()
        offset += size
        typ := (*cnt).GetType()
        json, err := json.Marshal((*cnt).GetName())
        if err != nil { log.Panicf("Can't marshal counter name: %s", err); }
        fmt.Fprintf(&scheme, "%s %d: %s\n", typ, size, json)
    }
    basepath := os.Getenv("CANTAL_PATH")
    if basepath == "" {
        runtime_dir := os.Getenv("XDG_RUNTIME_DIR")

        if runtime_dir != "" {
            basepath = fmt.Sprintf("%s/cantal.%d", runtime_dir, os.Getpid())
        } else {
            basepath = fmt.Sprintf("/tmp/cantal.%d.%d",
                os.Getuid(), os.Getpid())
        }

        log.Printf("Warning: No CANTAL_PATH is set in the environment, "+
            "using %s. The cantal-agent will be unable to discover it.",
            basepath)
    }

    tmppath := fmt.Sprintf("%s.tmp", basepath)
    metapath := fmt.Sprintf("%s.meta", basepath)
    path := fmt.Sprintf("%s.values", basepath)

    err := os.Remove(tmppath)
    if err != nil && !os.IsNotExist(err) {
        log.Panicf("Can't delete file %s: %s", tmppath, err);
    }
    err = os.Remove(path)
    if err != nil && !os.IsNotExist(err) {
        log.Panicf("Can't delete file %s: %s", path, err);
    }
    err = os.Remove(metapath)
    if err != nil && !os.IsNotExist(err) {
        log.Panicf("Can't delete file %s: %s", metapath, err);
    }

    // TODO(tailhook) write data

    file, err := os.Create(tmppath)
    if err != nil { log.Panicf("Can't open cantal file: %s", err); }
    _, err = scheme.WriteTo(file)
    if err != nil { log.Panicf("Can't write cantal metadata: %s", err); }
    file.Close()

    err = os.Rename(tmppath, metapath)
    if err != nil { log.Panicf("Can't rename cantal metadata: %s", err); }

    meta_path = metapath
    data_path = path
}

func Clean() {
    os.Remove(data_path)
    os.Remove(meta_path)
}
