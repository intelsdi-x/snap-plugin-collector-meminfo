# snap collector plugin - meminfo
This plugin collects metrics from /proc/meminfo kernel interface about distribution and utilization of memory.  

It's used in the [snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements
* [golang 1.4+](https://golang.org/dl/)

### Operating systems
All OSs currently supported by snap:
* Linux/amd64

### Installation
#### Download meminfo plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [GitHub Releases](https://github.com/intelsdi-x/snap/releases) page.

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-collector-meminfo  
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-meminfo.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `/build/rootfs/`

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported  
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/build`

## Documentation

### Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Description (optional)
----------|-----------------------
/intel/procfs/meminfo/Active | The total amount of buffer or page cache memory, in bytes, that is in active use; this memory has been used more recently and usually not reclaimed unless absolutely necessary.
/intel/procfs/meminfo/Active_anon | 
/intel/procfs/meminfo/Active_file | 
/intel/procfs/meminfo/AnonHugePages | The size of non-file backed huge pages mapped into user-space page tables, in bytes
/intel/procfs/meminfo/AnonPages | The size of non-file backed pages mapped into user-space page tables, in bytes
/intel/procfs/meminfo/Bounce | The amount of memory used for block device "bounce buffers", in bytes
/intel/procfs/meminfo/Buffers | The amount of physical RAM, in bytes, used for file buffers
/intel/procfs/meminfo/Cached | The amount of physical RAM, in bytes, used as cache memory
/intel/procfs/meminfo/CmaFree | The size of Contiguous Memory Allocator pages, in bytes, which are not used
/intel/procfs/meminfo/CmaTotal | The total size of Contiguous Memory Allocator pages, in bytes
/intel/procfs/meminfo/CommitLimit | The  amount of  memory, in bytes, currently available to be allocated on the system based on the overcommit ratio
/intel/procfs/meminfo/Committed_AS | The amount of memory, in bytes, estimated to complete the workload; this value represents the worst case scenario value, and also includes swap memory
/intel/procfs/meminfo/DirectMap1G | 
/intel/procfs/meminfo/DirectMap2M | 
/intel/procfs/meminfo/DirectMap4k |
/intel/procfs/meminfo/Dirty | The total amount of memory, in bytes, waiting to be written back to the disk.
/intel/procfs/meminfo/HardwareCorrupted | 
/intel/procfs/meminfo/HugePages_Free | The total number of hugepages available for the system
/intel/procfs/meminfo/HugePages_Rsvd | The number of huge pages for which a commitment to allocate from the pool has been made, but no allocation has yet been made.
/intel/procfs/meminfo/HugePages_Surp |The number of huge pages in the pool above the value in /proc/sys/vm/nr_hugepages
/intel/procfs/meminfo/HugePages_Total | The total number of hugepages for the system
/intel/procfs/meminfo/Hugepagesize | The size for each hugepages unit, in bytes
/intel/procfs/meminfo/Inactive | The total amount of buffer or page cache memory, in bytes, that are free and available; this memory has not been recently used and can be reclaimed for other purposes
/intel/procfs/meminfo/Inactive_anon |  
/intel/procfs/meminfo/Inactive_file | 
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
/intel/procfs/meminfo/Unevictable | 
/intel/procfs/meminfo/VmallocChunk | The largest contiguous block of vmalloc area, in bytes, which is free
/intel/procfs/meminfo/VmallocTotal | The total size of vmalloc memory area in bytes
/intel/procfs/meminfo/VmallocUsed | The amount of vmalloc area, in bytes, which is used
/intel/procfs/meminfo/Writeback | The total amount of memory, in bytes, actively being written back to the disk
/intel/procfs/meminfo/WritebackTmp | The amount of memory, in bytes, used by FUSE for temporary writeback buffers

All above metrics are additionally presented as percentage of total available memory. Those metrics has added suffix "_perc".

### Examples
Example running meminfo, passthru processor, and writing data to a file.

This is done from the snap directory.

In one terminal window, open the snap daemon (in this case with logging set to 1 and trust disabled):
```
$ $SNAP_PATH/bin/snapd -l 1 -t 0
```

In another terminal window:
Load meminfo plugin
```
$ $SNAP_PATH/bin/snapctl plugin load snap-plugin-collector-meminfo
```
See available metrics for your system
```
$ $SNAP_PATH/bin/snapctl metric list
```

Create a task manifest file (e.g. `mem-file.json`):    
```json
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "1s"
    },
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/procfs/meminfo/MemFree": {},
                "/intel/procfs/meminfo/MemAvailable": {},
                "/intel/procfs/meminfo/MemTotal": {},
                "/intel/procfs/meminfo/MemUsed": {}
            },
            "config": {},
            "process": null,
            "publish": [
                {
                    "plugin_name": "file",
                    "config": {
                        "file": "/tmp/published_meminfo"
                    }
                }
            ]
        }
    }
}
```
Load file plugin for publishing:
```
$ $SNAP_PATH/bin/snapctl plugin load build/plugin/snap-publisher-file
Plugin loaded
Name: file
Version: 3
Type: publisher
Signed: false
Loaded Time: Fri, 20 Nov 2015 11:41:39 PST
```

Create task:
```
$ $SNAP_PATH/bin/snapctl task create -t examples/tasks/mem-file.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```

Stop task:
```
$ $SNAP_PATH/bin/snapctl task stop 02dd7ff4-8106-47e9-8b86-70067cd0a850
Task stopped:
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-meminfo/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-meminfo/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@MarcinKrolik](https://github.com/marcin-krolik/)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.