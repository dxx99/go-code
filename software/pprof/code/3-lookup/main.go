package main

import (
	"io"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
)

type LookupType int8

const (
	LookupGoroutine LookupType = iota
	LookupThreadCreate
	LookupHeap
	LookupAllocs
	LookupBlock
	LookupMutex
)

func pprofLookup(lType LookupType, w io.Writer) error {
	var err error
	switch lType {
	case LookupGoroutine:
		p := pprof.Lookup("goroutine")
		err = p.WriteTo(w, 2)
	case LookupThreadCreate:
		p := pprof.Lookup("threadcreate")
		err = p.WriteTo(w, 2)
	case LookupHeap:
		p := pprof.Lookup("heap")
		err = p.WriteTo(w, 2)
	case LookupAllocs:
		p := pprof.Lookup("allocs")
		err = p.WriteTo(w, 2)
	case LookupBlock:
		p := pprof.Lookup("block")
		err = p.WriteTo(w, 2)
	case LookupMutex:
		p := pprof.Lookup("mutex")
		err = p.WriteTo(w, 2)
	}

	return err
}

func init() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)
}

func main() {

	http.HandleFunc("/lookup/heap", func(w http.ResponseWriter, r *http.Request) {
		_ = pprofLookup(LookupHeap, os.Stdout)
	})

	http.HandleFunc("/lookup/goroutine", func(w http.ResponseWriter, r *http.Request) {
		_ = pprofLookup(LookupGoroutine, os.Stdout)
	})

	http.HandleFunc("/lookup/block", func(w http.ResponseWriter, r *http.Request) {
		_ = pprofLookup(LookupBlock, os.Stdout)
	})

	http.HandleFunc("/lookup/threadcreate", func(w http.ResponseWriter, r *http.Request) {
		_ = pprofLookup(LookupThreadCreate, os.Stdout)
	})

	_ = http.ListenAndServe(":6009", nil)
}
