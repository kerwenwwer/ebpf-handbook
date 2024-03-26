package main

import (
	"log"
	"time"

	"github.com/cilium/ebpf/rlimit"
	bpf "github.com/kerwenwwer/ebpf-handbook/examples/kprobe/bpf"
)

func main() {

	// Kprobe trace function name
	fn := "sys_exec"

	// Allow the current process to lock memory for eBPF resources.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal()
	}

	// Load Object with helper function
	objs, err := bpf.LoadObjects()
	if err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()

	// Attach kprobe
	link, err := bpf.AttachKprobe(objs, fn)
	if err != nil {
		log.Fatalf("attaching kprobe: %v", err)
	}
	defer link.Close()

	// Read loop reporting the total amount of times the kernel
	// function was entered, once per second.
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	log.Println("Waiting for events..")

	for range ticker.C {
		value, err := bpf.LookupMap(objs)
		if err != nil {
			log.Fatalf("reading map: %v", err)
		}
		log.Printf("%s called %d times\n", fn, value)
	}

}
