//+build ignore

/* You need to add build ingnore in the first line of the file.*/

#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>

struct {
    __uint(type, BPF_MAP_TYPE_ARRAY);
    __type(key, __u32);
    __type(value, __u64);
    __uint(max_entries, 1);
} kprobe_map SEC(".maps");

SEC("kprobe/sys_execve")
int kprobe_execve() {
    __u32 key = 0;
    __u64 initval = 1, *valp;

    valp = bpf_map_lookup_elem(&kprobe_map, &key);
	if (!valp) {
		bpf_map_update_elem(&kprobe_map, &key, &initval, BPF_ANY);
		return 0;
	}
	__sync_fetch_and_add(valp, 1);

	return 0;
}

char __license[] SEC("license") = "Dual MIT/GPL";
