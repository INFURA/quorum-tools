# Cheat Sheet

```
# Building

## Build in yours devops machine
./devops/03-management/bash_devops "./banketh/devops/01-build/build_constellation"
./devops/03-management/bash_devops "./banketh/devops/01-build/build_geth"

## Bring the files home
## In case you just want to use the DEVOPS machine as a mere RHEL7.2 compiler
## and perform all your operations from a local machine
scp banketh@<DEVOPS MACHINE IP>:/home/banketh/banketh/devops/01-build/bin/* devops/01-build/bin/.

## Preparing Assets
## You can use the DEVOPS machine, or your local machine, provided you have quorum/geth
## and nodejs + the genesis builder application installed
./devops/02-deploy/01-prepare-assets/address
./devops/02-deploy/01-prepare-assets/node
./devops/02-deploy/01-prepare-assets/genesis <path of your genesis-specs.json>

# Deploy - Upload
## You should be able to perform these steps from your local machine
## Of course, you need to have the files to upload there, if not,
## call all these script using `./devops/03-management/bash_devops "<./banketh/script>"`
./devops/02-deploy/02-upload/assets
./devops/02-deploy/02-upload/geth <node number>

# Launch Geth
# Geth is already uploaded, so issue the order
# You can add as many command options for geth (<params>) as you want
./devops/02-deploy/03-launch/geth <node number> <params>

# Management

## Bash to devops
./devops/03-management/bash_devops

## Bash to a particular node
./devops/03-management/bash_node <node number>

## tail -f to machine
./devops/03-management/logs_node 0

## Pipe to the console itself
./devops/03-management/console_nodes 0

## Useful console commands
### Node Info
./devops/03-management/console_node 0 admin.nodeInfo

### New Account
./devops/03-management/console_node 0 "personal.newAccount('ola')"

## Tail the log of all the nodes at once
./devops/03-management/tail_all_nodes

## Stop a particular geth (example node 1)
i=1; ./devops/03-management/bash_node $i "pid=\$(ps aux | grep 'geth' | awk '{print \$2}' | head -1); echo \$pid |xargs kill"

## Remove a particular chaindata
i=1; ./devops/03-management/bash_node $i "rm -rf /home/banketh/ethereum/geth/chaindata"

## Stop all geth processes in the nodes
for i in `seq 0 6`; do ./devops/03-management/bash_node $i "pid=\$(ps aux | grep 'geth' | awk '{print \$2}' | head -1); echo \$pid |xargs kill"; done
```
