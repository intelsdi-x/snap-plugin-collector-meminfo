# Example tasks

[This](task-mem.json) example task will collect metrics from **meminfo** and publish
them to file.  

## Running the example

### Requirements 
 * `docker` and `docker-compose` are **installed** and **configured** 

Running the sample is as *easy* as running the script `./run-file-meminfo.sh`.

## Files
- [file-meminfo.sh](file-meminfo.sh)
    - Downloads `snapteld`, `snaptel`, `snap-plugin-collector-meminfo`,
        `snap-plugin-publisher-file-file` and starts the task
- [run-file-meminfo.sh](run-file-meminfo.sh)
    - The example is launched with this script     
- [tasks/task-mem.json](tasks/task-mem.json)
    - Snap task definition
- [.setup.sh](.setup.sh)
    - Verifies dependencies and starts the containers.  It's called 
    by [run-file-meminfo.sh](run-file-meminfo.sh).
- [docker-compose.yml](docker-compose.yml)
    - A docker compose file which defines "runner" container where snapteld
     is run from. You will be dumped into a shell in this container
     after running.    
