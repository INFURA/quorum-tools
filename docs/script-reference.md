# Script Reference

Details the contents of the devops directory, and gives a short reference on
what each script does

## 00-environment

Prepares the DEVOPS server and the NODES server. The former builds the
needed programs in the same OS as the NODES.

* bootstrap_devops_machine
* bootstrap_node_machine

Copy these files in their respective machines, and just run them.

## 01-build

Will build into this directory's `bin` directory the `geth` (quorum)
executable and the `constellation-*` executables.

Need to be run in a machine with the adequate building libraries.

* build_constellation
* build_geth

## 02-deploy

* Prepare assets (genesis file, voter and blockmaker addresses and priv/pub node keys)
* Upload the files into the right servers
* Launch the executables

### 00-assets

Contains the prepared assets (ignored by `git`).

### 01-prepare-assets

* address: Needs to have `geth` installed in the machine. Gives you a `geth`
address file.
* node: Creates a 512-bit pseudo random key, and then calls a `golang` executable
to generate its `ECSDA` public key (useful for `enode` composition).
* genesis: Needs to have `node.js` and `quorum-genesis` utility. Takes a json
spec file and outputs the genesis file for the chain.

### 02-upload

* assets: Takes all assets in the directory `00-assets` and uploads them to
the right server, based on its name.
* constellation: Uploads the constellation binaries found in `01-build/bin`.
* geth: Uploads the geth binary found in `01-build/bin`

### 03-launch

* constellation: Launches constellation in the indicated server.
+ geth: Launches geth in the indicated server.

## 03-management

Utilities to access the functionalities of the nodes.

* bash_devops: `SSH` into the `DEVOPS` server.
* bash_node: `SSH` into the node server number N.
* console_node: Sends a geth console command to the node server.
* logs_node: Performs a `tail -f` in the specified node.
* tail_all_nodes: Perform a `tail` over all nodes.

## 04-cluster-layout

Utilities to setup clusters under certain specs.

* 7_5_3_NPr_NPe: Deploys a 7-node, 5-voters, 3-blockmakers, no privacy
features, no permissioned nodes cluster.

## 05-test-benchmark

* WIP
