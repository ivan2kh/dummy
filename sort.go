package main
//sort and compress 64MB 32b ints

import "fmt"
import "sort"
import "math/rand"
import "os"
import "bufio"
import "encoding/gob"
import "compress/zlib"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

var inp [16777216]int

func write(name string) {
    f, err := os.Create(name)
    check(err)
    defer f.Close()
    w := bufio.NewWriter(f)
    encoder := gob.NewEncoder(w)
    err = encoder.Encode(inp)
    check(err)
}

func writeCompress(name string) {
    f, err := os.Create(name)
    check(err)
    defer f.Close()
    w := bufio.NewWriter(f)
    compr := zlib.NewWriter(w)
    encoder := gob.NewEncoder(compr)
    err = encoder.Encode(inp)
    check(err)
}

func main() {
    for i := range inp {
        inp[i]=rand.Intn(16777216)
    }
    write("input.dat")

    sort.Ints(inp[:])

    write("sorted.dat")

    fmt.Println("inp sorted:   ", inp[16777206:16777215])
    
    for i := 16777216-1; i > 0; i-- {
        inp[i] = inp[i] - inp[i-1]
    }

    fmt.Println("inp diffs:   ", inp[1:10])

    write("diffs.dat")

    writeCompress("compr.dat")

    for i := 1; i < 16777216; i++ {
        inp[i] = inp[i] + inp[i-1]
    }

    write("recover.dat")
}