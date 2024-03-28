# BPF Project Examples

This repository contains example projects demonstrating the use of eBPF (Extended Berkeley Packet Filter) in various contexts, including `kprobe`, `tc` (Traffic Control), and `xdp` (eXpress Data Path). These examples are intended to provide a starting point for developing BPF applications and to showcase the standard project structure for such developments.

## Project Structure

The project is organized into separate directories for each example, following a standard BPF project layout. Below is an outline of the project structure, using the `kprobe` example for illustration:
```bash
.
├── kprobe
│ ├── bin
│ │ └── kprobe # Compiled binary
│ ├── bpf
│ │ ├── bpf_bpfeb.go # BPF bytecode for BE (Big Endian) systems
│ │ ├── bpf_bpfeb.o # Compiled BPF object file for BE systems
│ │ ├── bpf_bpfel.go # BPF bytecode for LE (Little Endian) systems
│ │ ├── bpf_bpfel.o # Compiled BPF object file for LE systems
│ │ ├── bpf.c # BPF program in C
│ │ └── bpf.go # Go bindings for the BPF program
│ ├── go.mod # Go module file
│ ├── go.sum # Go module checksum file
│ ├── main.go # Main application file
│ └── Makefile # Makefile for building the project
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