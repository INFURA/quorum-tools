# Servers

The servers directory contains instructions (in bash) to prepare two kind
of machines:

* Compiler Machine
* Node Machine
* Assets Machine
* Cluster IPs

## Compiler machine

The compiler machine is useful to prepare for you executables in `RHEL 7.2.` of
both geth (quorum) and constellation.

In the future, the idea is having a service that compiles the lastest version
of this program, and make those available (See [This issue](https://github.com/ConsenSys/quorum-tools/issues/5)).

To prepare your compiler machine, make sure you initialize an `RHEL 7.2.` box,
copy and run the script

```
./01-servers/bootstrap_compiler_machine
```

The building scripts in the directory `02-build` are designed to perform
all their activities via `ssh`, to ensure that you don't need to copy and paste more
scripts inside this machine.

### Why not a container?

Currently (2017.MAR.09), there are not Red Hat docker images available for the public.

## Node machine

The node machine is where geth (quorum) and constellation will be running.

To prepare your node machine, make sure you initialize an `RHEL 7.2.` box,
copy and run the script

```
./01-servers/bootstrap_node_machine $USER
```

Where `$USER` is your `sudoer` of course. Or, put in another way, the user you
created the server.

This script will configure the ports in your machine, install dependencies to run
constellation and prepare the directories for executables and ethereum.

## Assets Machine

You can find the latest assets machine built at

```
https://hub.docker.com/r/infura/quorum-tools-assets
```

This image enables you to build

* Geth (quorum) assets: ethereum addresses and nodekeys (priv/pub ECDSA keys)
* Constellation keys
* Quorum genesis files

If you want to build it yourself, just use the `Dockerfile` located in

```
./01-servers/assets-machine/Dockerfile
```

and use the command (provided you have `docker` installed):

```
docker build --squash -t <my-image-name> -f ./01-servers/assets-machine/Dockerfile .
```

If for some reason you want to log into the container, you can do

```
docker run -ti infura/quorum-tools-assets /bin/bash
```
