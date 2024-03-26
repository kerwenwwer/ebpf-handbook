all: build
clean: 
	rm ./bin/*
	rm ./bpf/bpf_*

.PHONY: bpf/*.o 
bpf/*.o: bpf/*.c
	go generate ./bpf/

.PHONY: build
build: bpf/*.o
	go build -o ./bin/kprobe ./main.go
