# snap plugin collector - meminfo

## Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Description
----------|-----------------------
/intel/procfs/meminfo/Active | The total amount of buffer or page cache memory, in bytes, that is in active use; this memory has been used more recently and usually not reclaimed unless absolutely necessary.
/intel/procfs/meminfo/Active_anon | The amount of anonymous memory, in bytes, that has been used more recently and usually not swapped out
/intel/procfs/meminfo/Active_file | The amount of pagecache memory, in bytes,  that has been used more recently and usually not reclaimed until needed
/intel/procfs/meminfo/AnonHugePages | The size of non-file backed huge pages mapped into user-space page tables, in bytes
/intel/procfs/meminfo/AnonPages | The size of non-file backed pages mapped into user-space page tables, in bytes
/intel/procfs/meminfo/Bounce | The amount of memory used for block device "bounce buffers", in bytes
/intel/procfs/meminfo/Buffers | The amount of physical RAM, in bytes, used for file buffers
/intel/procfs/meminfo/Cached | The amount of physical RAM, in bytes, used as cache memory
/intel/procfs/meminfo/CmaFree | The size of Contiguous Memory Allocator pages, in bytes, which are not used
/intel/procfs/meminfo/CmaTotal | The total size of Contiguous Memory Allocator pages, in bytes
/intel/procfs/meminfo/CommitLimit | The  amount of  memory, in bytes, currently available to be allocated on the system based on the overcommit ratio
/intel/procfs/meminfo/Committed_AS | The amount of memory, in bytes, estimated to complete the workload; this value represents the worst case scenario value, and also includes swap memory
/intel/procfs/meminfo/DirectMap1G | The amount of memory, in bytes, being mapped to 1 G pages
/intel/procfs/meminfo/DirectMap2M | The amount of memory, in bytes, being mapped to 2 MB pages
/intel/procfs/meminfo/DirectMap4k | The amount of memory, in bytes, being mapped to standard 4k pages
/intel/procfs/meminfo/Dirty | The total amount of memory, in bytes, waiting to be written back to the disk.
/intel/procfs/meminfo/HardwareCorrupted | The amount of failed memory in bytes (can only be detected when using ECC RAM).
/intel/procfs/meminfo/HugePages_Free | The total number of hugepages available for the system
/intel/procfs/meminfo/HugePages_Rsvd | The number of huge pages for which a commitment to allocate from the pool has been made, but no allocation has yet been made.
/intel/procfs/meminfo/HugePages_Surp |The number of huge pages in the pool above the value in /proc/sys/vm/nr_hugepages
/intel/procfs/meminfo/HugePages_Total | The total number of hugepages for the system
/intel/procfs/meminfo/Hugepagesize | The size for each hugepages unit, in bytes
/intel/procfs/meminfo/Inactive | The total amount of buffer or page cache memory, in bytes, that are free and available; this memory has not been recently used and can be reclaimed for other purposes
/intel/procfs/meminfo/Inactive_anon | The amount of anonymous memory, in bytes, that has not been used recently and can be swapped out
/intel/procfs/meminfo/Inactive_file | The amount of  pagecache memory, in bytes, that can be reclaimed without huge performance impact
/intel/procfs/meminfo/KernelStack | The amount of memory allocated to kernel stacks in bytes
/intel/procfs/meminfo/Mapped | The total amount of memory, in bytes, which have been used to map devices, files, or libraries using the mmap command
/intel/procfs/meminfo/MemAvailable | The estimated amount of memory, in bytes, which is available for starting new applications without swapping
/intel/procfs/meminfo/MemFree | The amount of physical RAM, in bytes, left unused by the system (the sum of LowFree+HighFree)
/intel/procfs/meminfo/MemTotal | Total amount of physical RAM, in bytes
/intel/procfs/meminfo/MemUsed | The amount of physical Ram, in bytes which is used; it equals: MemTotal-(MemFree+Buffers+Cached+Slab)
/intel/procfs/meminfo/Mlocked | The total amount of memory, in bytes, which is locked from userspace.
/intel/procfs/meminfo/NFS_Unstable | The size of NFS pages, in bytes, which are sent to the server, but not yet committed to stable storage
/intel/procfs/meminfo/PageTables | The total amount of memory, in bytes, dedicated to the lowest page table level.
/intel/procfs/meminfo/SReclaimable | The part of Slab, in bytes, that might be reclaimed, such as caches
/intel/procfs/meminfo/SUnreclaim | The part of Slab, in bytes, that cannot be reclaimed on memory pressure
/intel/procfs/meminfo/Shmem | The total amount of memory, in bytes, which is shared
/intel/procfs/meminfo/Slab | The total amount of memory, in bytes, used by the kernel to cache data structures for its own use.
/intel/procfs/meminfo/SwapCached | The amount of swap, in bytes, used as cache memory
/intel/procfs/meminfo/SwapFree | The total amount of swap free, in bytes 
/intel/procfs/meminfo/SwapTotal | The total amount of swap available, in bytes
/intel/procfs/meminfo/Unevictable | The amount of memory, in bytes, that cannot be reclaimed (for example, because it is Mlocked or used as a RAM disk). 
/intel/procfs/meminfo/VmallocChunk | The largest contiguous block of vmalloc area, in bytes, which is free
/intel/procfs/meminfo/VmallocTotal | The total size of vmalloc memory area in bytes
/intel/procfs/meminfo/VmallocUsed | The amount of vmalloc area, in bytes, which is used
/intel/procfs/meminfo/Writeback | The total amount of memory, in bytes, actively being written back to the disk
/intel/procfs/meminfo/WritebackTmp | The amount of memory, in bytes, used by FUSE for temporary writeback buffers

All above metrics are additionally presented as percentage of total available memory. Those metrics has added suffix "_perc".