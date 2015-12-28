# snap collector plugin - meminfo
This plugin collects metrics from /proc/meminfo kernel interface about distribution and utilization of memory.  

It's used in the [snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
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
/intel/linux/meminfo/Active | Memory that has been used more recently and usually not reclaimed unless absolutely necessary
/intel/linux/meminfo/Active(anon) | 
/intel/linux/meminfo/Active(file) | 
/intel/linux/meminfo/AnonHugePages |
/intel/linux/meminfo/AnonPages |
/intel/linux/meminfo/Bounce | Memory used for block device "bounce buffers"
/intel/linux/meminfo/Buffers | Relatively temporary storage for raw disk blocks
/intel/linux/meminfo/Cached | in-memory cache for files read from the disk
/intel/linux/meminfo/CmaFree |
/intel/linux/meminfo/CmaTotal | 
/intel/linux/meminfo/CommitLimit | total amount of  memory currently available to be allocated on the system based on the overcommit ratio
/intel/linux/meminfo/Committed_AS | The total amount of memory, estimated to complete the workload 
/intel/linux/meminfo/DirectMap1G | 
/intel/linux/meminfo/DirectMap2M | 
/intel/linux/meminfo/DirectMap4k |
/intel/linux/meminfo/Dirty | The total amount of memory, waiting to be written back to the disk
/intel/linux/meminfo/HardwareCorrupted | 
/intel/linux/meminfo/HugePages_Free | The total number of hugepages available for the system
/intel/linux/meminfo/HugePages_Rsvd |
/intel/linux/meminfo/HugePages_Surp |
/intel/linux/meminfo/HugePages_Total | The total number of hugepages for the system
/intel/linux/meminfo/Hugepagesize | The size for each hugepages
/intel/linux/meminfo/Inactive | The total amount of buffer or page cache memory, that are free and available
/intel/linux/meminfo/Inactive(anon) |
/intel/linux/meminfo/Inactive(file) |
/intel/linux/meminfo/KernelStack | 
/intel/linux/meminfo/Mapped | The total amount of memory, which have been used to map devices, files, or libraries using the mmap command.
/intel/linux/meminfo/MemAvailable | Estimate of how much memory is available for starting new applications without swapping
/intel/linux/meminfo/MemFree | The sum of LowFree+HighFree
/intel/linux/meminfo/MemTotal | Total usable ram
/intel/linux/meminfo/MemUsed | MemTotal - (MemFree + Buffers + Cached + Slab
/intel/linux/meminfo/Mlocked | 
/intel/linux/meminfo/NFS_Unstable | NFS pages sent to the server, but not yet committed to stable storage
/intel/linux/meminfo/PageTables | amount of memory dedicated to the lowest level of page tables
/intel/linux/meminfo/SReclaimable | Part of Slab, that might be reclaimed, such as caches
/intel/linux/meminfo/SUnreclaim | Part of Slab, that cannot be reclaimed on memory pressure
/intel/linux/meminfo/Shmem | 
/intel/linux/meminfo/Slab | in-kernel data structures cache
/intel/linux/meminfo/SwapCached | Memory that once was swapped out, is swapped back in but still also is in the swapfile
/intel/linux/meminfo/SwapFree | Memory which has been evicted from RAM, and is temporarily on the disk
/intel/linux/meminfo/SwapTotal | total amount of swap space available
/intel/linux/meminfo/Unevictable | 
/intel/linux/meminfo/VmallocChunk | largest contiguous block of vmalloc area which is free
/intel/linux/meminfo/VmallocTotal | total size of vmalloc memory area
/intel/linux/meminfo/VmallocUsed | amount of vmalloc area which is used
/intel/linux/meminfo/Writeback | Memory which is actively being written back to the disk
/intel/linux/meminfo/WritebackTmp | Memory used by FUSE for temporary writeback buffers

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
                "/intel/linux/meminfo/MemFree": {},
                "/intel/linux/meminfo/MemAvailable": {}, 
                "/intel/linux/meminfo/MemTotal": {},
                "/intel/linux/meminfo/MemUsed": {} 
            },
            "config": {
                "/intel/mock": {
                    "password": "secret",
                    "user": "root"
                }
            },
            "process": [
                {
                    "plugin_name": "passthru",
                    "process": null,
                    "publish": [
                        {                         
                            "plugin_name": "file",
                            "config": {
                                "file": "/tmp/published_meminfo"
                            }
                        }
                    ],
                    "config": null
                }
            ],
            "publish": null
        }
    }
}
```

Load passthru plugin for processing:
```
$ $SNAP_PATH/bin/snapctl plugin load build/plugin/snap-processor-passthru
Plugin loaded
Name: passthru
Version: 1
Type: processor
Signed: false
Loaded Time: Fri, 20 Nov 2015 11:44:03 PST
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