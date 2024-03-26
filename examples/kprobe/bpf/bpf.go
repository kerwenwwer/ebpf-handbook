package bpf

import (
	"errors"
	"fmt"
	"os"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang -cflags "-O2 -Wall" bpf ./bpf.c

// Only one entries in kprobe map, so the map key is fixed to 0.
const mapKey uint32 = 0

func LoadObjects() (*bpfObjects, error) {
	var objs bpfObjects
	// Load pre-compiled programs and maps into the kernel.
	if err := loadBpfObjects(&objs, nil); err != nil {
		var ve *ebpf.VerifierError
		if errors.As(err, &ve) {
			fmt.Fprintf(os.Stderr, "Verifier errors:\n%s\n", ve.Error())
		}
		return nil, err
	}

	return &objs, nil
}

func AttachKprobe(objs *bpfObjects, fn string) (link.Link, error) {
	// Open a Kprobe at the entry point of the kernel function and attach the
	// pre-compiled program. Each time the kernel function enters, the program
	// will increment the execution counter by 1. The read loop below polls this
	// map value once per second.
	kp, err := link.Kprobe(fn, objs.KprobeExecve, nil)

	return kp, err
}

func LookupMap(objs *bpfObjects) (uint64, error) {
	var value uint64
	err := objs.KprobeMap.Lookup(mapKey, &value)

	return value, err
}
