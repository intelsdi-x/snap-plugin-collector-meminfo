# snap plugin collector - meminfo

## Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Description
----------|-----------------------
/intel/procfs/meminfo/active | The total amount of buffer or page cache memory, in bytes, that is in active use; this memory has been used more recently and usually not reclaimed unless absolutely necessary.
/intel/procfs/meminfo/active_anon | The amount of anonymous memory, in bytes, that has been used more recently and usually not swapped out
/intel/procfs/meminfo/active_file | The amount of pagecache memory, in bytes,  that has been used more recently and usually not reclaimed until needed
/intel/procfs/meminfo/anon_huge_pages | The size of non-file backed huge pages mapped into user-space page tables, in bytes
/intel/procfs/meminfo/anon_pages | The size of non-file backed pages mapped into user-space page tables, in bytes
/intel/procfs/meminfo/bounce | The amount of memory used for block device "bounce buffers", in bytes
/intel/procfs/meminfo/buffers | The amount of physical RAM, in bytes, used for file buffers
/intel/procfs/meminfo/cached | The amount of physical RAM, in bytes, used as cache memory
/intel/procfs/meminfo/cma_free | The size of Contiguous Memory Allocator pages, in bytes, which are not used
/intel/procfs/meminfo/cma_total | The total size of Contiguous Memory Allocator pages, in bytes
/intel/procfs/meminfo/commit_limit | The  amount of  memory, in bytes, currently available to be allocated on the system based on the overcommit ratio
/intel/procfs/meminfo/Committed_as | The amount of memory, in bytes, estimated to complete the workload; this value represents the worst case scenario value, and also includes swap memory
/intel/procfs/meminfo/direct_map1g | The amount of memory, in bytes, being mapped to 1 G pages
/intel/procfs/meminfo/direct_map2m | The amount of memory, in bytes, being mapped to 2 MB pages
/intel/procfs/meminfo/direct_map4k | The amount of memory, in bytes, being mapped to standard 4k pages
/intel/procfs/meminfo/dirty | The total amount of memory, in bytes, waiting to be written back to the disk.
/intel/procfs/meminfo/hardware_corrupted | The amount of failed memory in bytes (can only be detected when using ECC RAM).
/intel/procfs/meminfo/huge_pages_free | The total number of hugepages available for the system
/intel/procfs/meminfo/huge_pages_rsvd | The number of huge pages for which a commitment to allocate from the pool has been made, but no allocation has yet been made.
/intel/procfs/meminfo/huge_pages_surp |The number of huge pages in the pool above the value in /proc/sys/vm/nr_hugepages
/intel/procfs/meminfo/huge_pages_total | The total number of hugepages for the system
/intel/procfs/meminfo/hugepagesize | The size for each hugepages unit, in bytes
/intel/procfs/meminfo/inactive | The total amount of buffer or page cache memory, in bytes, that are free and available; this memory has not been recently used and can be reclaimed for other purposes
/intel/procfs/meminfo/inactive_anon | The amount of anonymous memory, in bytes, that has not been used recently and can be swapped out
/intel/procfs/meminfo/inactive_file | The amount of  pagecache memory, in bytes, that can be reclaimed without huge performance impact
/intel/procfs/meminfo/kernel_stack | The amount of memory allocated to kernel stacks in bytes
/intel/procfs/meminfo/mapped | The total amount of memory, in bytes, which have been used to map devices, files, or libraries using the mmap command
/intel/procfs/meminfo/mem_available | The estimated amount of memory, in bytes, which is available for starting new applications without swapping
/intel/procfs/meminfo/mem_free | The amount of physical RAM, in bytes, left unused by the system (the sum of low_free+high_free)
/intel/procfs/meminfo/mem_total | Total amount of physical RAM, in bytes
/intel/procfs/meminfo/mem_used | The amount of physical Ram, in bytes which is used; it equals: mem_total-(mem_free+buffers+cached+slab)
/intel/procfs/meminfo/mlocked | The total amount of memory, in bytes, which is locked from userspace.
/intel/procfs/meminfo/nfs_unstable | The size of NFS pages, in bytes, which are sent to the server, but not yet committed to stable storage
/intel/procfs/meminfo/page_tables | The total amount of memory, in bytes, dedicated to the lowest page table level.
/intel/procfs/meminfo/sreclaimable | The part of Slab, in bytes, that might be reclaimed, such as caches
/intel/procfs/meminfo/sunreclaim | The part of Slab, in bytes, that cannot be reclaimed on memory pressure
/intel/procfs/meminfo/shmem | The total amount of memory, in bytes, which is shared
/intel/procfs/meminfo/slab | The total amount of memory, in bytes, used by the kernel to cache data structures for its own use.
/intel/procfs/meminfo/swap_cached | The amount of swap, in bytes, used as cache memory
/intel/procfs/meminfo/swap_free | The total amount of swap free, in bytes
/intel/procfs/meminfo/swap_total | The total amount of swap available, in bytes
/intel/procfs/meminfo/unevictable | The amount of memory, in bytes, that cannot be reclaimed (for example, because it is Mlocked or used as a RAM disk).
/intel/procfs/meminfo/vmalloc_chunk | The largest contiguous block of vmalloc area, in bytes, which is free
/intel/procfs/meminfo/vmalloc_total | The total size of vmalloc memory area in bytes
/intel/procfs/meminfo/vmalloc_used | The amount of vmalloc area, in bytes, which is used
/intel/procfs/meminfo/writeback | The total amount of memory, in bytes, actively being written back to the disk
/intel/procfs/meminfo/writeback_tmp | The amount of memory, in bytes, used by FUSE for temporary writeback buffers

All above metrics are additionally presented as percentage of total available memory. Those metrics has added suffix "_perc".
