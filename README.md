# eBPF Handbook

Welcome to the eBPF Handbook, a comprehensive guide and reference for building and deploying eBPF applications using Go and the Cilium eBPF library. This repository is aimed at developers interested in the innovative eBPF technology, offering examples and patterns for various use cases.

## Introduction

Extended Berkeley Packet Filter (eBPF) is a powerful technology that allows code to run in the Linux kernel without changing the kernel source code or loading kernel modules. Thanks to eBPF, developers can write applications in high-level languages like Go, which are then executed at various points in the kernel. This handbook provides examples and guidance on leveraging eBPF for networking, security, and performance monitoring tasks.

## Prerequisites

Before you dive into the eBPF Handbook, make sure you have the following prerequisites installed on your system:

- Go (version 1.15 or later)
- Linux Kernel (version 4.19 or later, with eBPF support)
- Cilium eBPF library

For detailed installation instructions for each prerequisite, please refer to their respective documentation.

## Installation

Clone this repository to get started with the eBPF Handbook examples:

```sh
git clone https://github.com/<your-username>/ebpf-handbook.git
cd ebpf-handbook/examples/<the example your want>
make
```