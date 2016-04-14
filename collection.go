package cantal

import "bytes"
import "fmt"
import "os"
import "log"
import "encoding/json"
import "syscall"
import "unsafe"
import "reflect"

var active_counters = []*Value{}
var meta_path = ""
var data_path = ""
var mmap []byte = nil

func add_value(value Value) {
    active_counters = append(active_counters, &value)
}

func Start() {
    offset := 0
    scheme := bytes.Buffer{}
    for _, cnt := range active_counters {
        offset += (*cnt).GetSize()
        typ := (*cnt).GetType()
        json, err := json.Marshal((*cnt).GetName())
        if err != nil { log.Panicf("Can't marshal counter name: %s", err); }
        fmt.Fprintf(&scheme, "%s: %s\n", typ, json)
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

    file, err := os.Create(tmppath)
    slice := make([]byte, offset)
    _, err = file.Write(slice)
    if err != nil {
        log.Panicf("Can't write cantal data: %s", err);
    }
    mbytes, err := syscall.Mmap(int(file.Fd()),
        0, offset, syscall.PROT_WRITE, syscall.MAP_SHARED)
    if err != nil {
        log.Panicf("Can't memory map data: %s", err);
    }
    mmap = mbytes
    file.Close();
    err = os.Rename(tmppath, path)
    if err != nil { log.Panicf("Can't rename cantal data: %s", err); }
    data_path = path

    // TODO(tailhook) close file on panic?
    file, err = os.Create(tmppath)
    if err != nil { log.Panicf("Can't open cantal file: %s", err); }
    _, err = scheme.WriteTo(file)
    if err != nil { log.Panicf("Can't write cantal metadata: %s", err); }
    file.Close()

    err = os.Rename(tmppath, metapath)
    if err != nil { log.Panicf("Can't rename cantal metadata: %s", err); }
    meta_path = metapath

    head := (*reflect.SliceHeader)(unsafe.Pointer(&mmap))
    ptr := head.Data
    for _, cnt := range active_counters {
        (*cnt).set_pointer(unsafe.Pointer(ptr))
        ptr = ptr + uintptr((*cnt).GetSize())
    }
}

func Clean() {
    if data_path != "" {
        os.Remove(data_path)
        data_path = ""
    }
    if meta_path != "" {
        os.Remove(meta_path)
        meta_path = ""
    }
    if mmap != nil {
        head := (*reflect.SliceHeader)(unsafe.Pointer(&mmap))
        syscall.Syscall(syscall.SYS_MUNMAP,
            uintptr(head.Data), uintptr(len(mmap)), 0)
        mmap = nil
    }
}
