# BPF Project Examples

This repository contains example projects demonstrating the use of eBPF (Extended Berkeley Packet Filter) in various contexts, including `kprobe`, `tc` (Traffic Control), and `xdp` (eXpress Data Path). These examples are intended to provide a starting point for developing BPF applications and to showcase the standard project structure for such developments.

## Project Structure

The project is organized into separate directories for each example, following a standard BPF project layout. Below is an outline of the project structure, using the `kprobe` example for illustration:
```bash
.
├── examples
│   ├── kprobe
│   │   ├── bin
│   │   │   └── kprobe
│   │   ├── bpf
│   │   │   ├── bpf_bpfeb.go
│   │   │   ├── bpf_bpfeb.o
│   │   │   ├── bpf_bpfel.go
│   │   │   ├── bpf_bpfel.o
│   │   │   ├── bpf.c
│   │   │   └── bpf.go
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── main.go
│   │   └── Makefile
```

Each example directory (`kprobe`, `tc`, `xdp`) follows the same structural pattern to maintain consistency and ease of understanding across different BPF applications.

## Getting Started

To get started with these examples, you will need a Linux environment with the necessary tools for eBPF development installed, including the `bcc` toolkit, `libbpf`, and a modern version of `clang` and `llvm`.

### Building the Examples

Each example can be built by navigating to its respective directory and running the `make` command. This will compile the BPF C code and the user space application:

```bash
cd kprobe
make
```

Repeat this process for the tc and xdp examples as needed.

### Contributing

Contributions to improve the examples or documentation are welcome. Please feel free to submit issues or pull requests with your suggestions or improvements.