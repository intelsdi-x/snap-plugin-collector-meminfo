# Snap collector plugin - meminfo
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
* [golang 1.6+](https://golang.org/dl/)

### Operating systems
All OSs currently supported by snap:
* Linux/amd64

### Installation
#### Download meminfo plugin binary:
You can get the pre-built binaries for your OS and architecture at plugin's [GitHub Releases](https://github.com/intelsdi-x/snap-plugin-collector-meminfo/releases) page.

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
This builds the plugin in `./build`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

* If /proc resides in a different directory, say for example by mounting host /proc inside a container at /hostproc, a proc_path configuration item can be added to snapd global config or as part of the task manifest for the metrics to be collected.

As part of snapd global config

```yaml
---
control:
  plugins:
    collector:
      meminfo:
        all:
          proc_path: /hostproc
```

Or as part of the task manifest

```json
{
...
    "workflow": {
        "collect": {
            "metrics": {
	      "/intel/procfs/meminfo/mem_total": {}
	    },
	    "config": {
	      "/intel/procfs": {
                "proc_path": "/hostproc"
	      }
	    },
	    ...
       },
    },
...
```

* Load the plugin and create a task, see example in [Examples](https://github.com/intelsdi-x/snap-plugin-collector-meminfo/blob/master/README.md#examples).

## Documentation

### Collected Metrics
List of collected metrics is described in [METRICS.md](https://github.com/intelsdi-x/snap-plugin-collector-meminfo/blob/master/METRICS.md).

### Examples
Example running meminfo plugin and writing data to a file using [snap-plugin-publisher-file](https://github.com/intelsdi-x/snap-plugin-publisher-file).

Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

Ensure [Snap daemon is running](https://github.com/intelsdi-x/snap#running-snap):
* initd: `service snap-telemetry start`
* systemd: `systemctl start snap-telemetry`
* command line: `sudo snapd -l 1 -t 0 &`


Download and load Snap plugins:
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-meminfo/latest/linux/x86_64/snap-plugin-collector-meminfo
$ snapctl plugin load snap-plugin-publisher-file
$ snapctl plugin load snap-plugin-collector-meminfo
```

See available metrics for your system
```
$ snapctl metric list
```

Create a [task manifest](https://github.com/intelsdi-x/snap/blob/master/docs/TASKS.md) (see [exemplary tasks](examples/tasks/)),
for example `meminfo-file.json` with following content:
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
        "/intel/procfs/meminfo/mem_free": {},
        "/intel/procfs/meminfo/mem_available": {},
        "/intel/procfs/meminfo/mem_total": {},
        "/intel/procfs/meminfo/mem_used": {}
      },
      "config": {
        "/intel/procfs": {
          "proc_path": "/proc"
        }
      },
      "publish": [
        {
          "plugin_name": "file",
          "config": {
            "file": "/tmp/published_meminfo.log"
          }
        }
      ]
    }
  }
}
```

Create a task:
```
$ snapctl task create -t meminfo-file.json
```

Watch created task:
```
$ snapctl task watch <task_id>
```

To stop previously created task:
```
$ snapctl task stop <task_id>
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-meminfo/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-meminfo/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support) or visit [Slack channel](http://slack.snap-telemetry.io).

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

And **thank you!** Your contribution, through code and participation, is incredibly important to us.

## License
[Snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [Marcin Krolik](https://github.com/marcin-krolik/)
