package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"git.nspcc.ru/nspcc/tzhash/tz2"
	"git.nspcc.ru/nspcc/tzhash/tzC"
	"git.nspcc.ru/nspcc/tzhash/tzbits"
	"git.nspcc.ru/nspcc/tzhash/tzbitspure"
	"git.nspcc.ru/nspcc/tzhash/tzpure"
	"github.com/nspcc-dev/tzhash/tz"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var filename = flag.String("name", "", "file to use")
var hashfunc = flag.String("hash", "tzhash", "hash function to use")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	var f io.Reader
	var err error
	if *filename != "" {
		f, err = os.Open(*filename)
		if err != nil {
			exit("cant open file: %v\n", err)
		}
	} else {
		*filename = "-"
	}

	var h hash.Hash
	switch *hashfunc {
	case "sha256":
		h = sha256.New()
	case "tzC":
		h = tzC.New()
	case "tzCLoop":
		h = tzC.NewWithGoLoop()
	case "tzbits":
		h = tzbits.New()
	case "tz2":
		h = tz2.New()
	case "tzpure":
		h = tzpure.New()
	case "tzpurebits":
		h = tzbitspure.New()
	case "tz":
		h = tz.New()
	default:
		panic("kek")
	}

	// h.Write([]byte{0x63})
	// return
	if _, err := io.Copy(h, f); err != nil {
		exit("error while reading file: %v\n", err)
	}
	fmt.Printf("%x\t%s\n", h.Sum(nil), *filename)

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}

func exit(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
	os.Exit(1)
}
