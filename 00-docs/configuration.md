# Configuration

The configuration file is very straight-forward. It enables you
to perform otherwise cumbersome activities involving `ssh` and `scp`
commands.

You just need to take care of two files here

* `deploy.conf`
* `/01-servers/IPs/<cluster-name>-NODES`

## deploy.conf

Copy the file `deploy-sample.conf` to `deploy.conf` in this repository's
root

```
cp deploy-sample.conf deploy.conf
```

### Fields

#### cluster-name

Is an arbitrary name you assign to your cluster. The scripts will look for

* `01-servers/IPs/<cluster-name>-NODES`: Number of node, public IP, private IP
* `03-deploy/00-assets/<cluster-name>/*`: Your cluster assets (keys, genesis file, etc)

#### ssh-username

To access to your machines via `ssh`, the scripts will take this parameter at run
time to know which username employ. Make sure you set up all your nodes with the
same username!

#### ssl-cert-path

The physical route to your `SSL` certificate. It will be invoked everytime we perform
a connection to your machine. To help you avoid inconveniences (i.e. typing that
password a thousand times), add it to your agent with

```
eval `ssh-agent -s` # In case you don't have an agent running (OSX has one by default)

ssh-add <Path to your cert>

ssh-add -l # To verify
```

## Cluster IPs

The scripts in this suite will look for the file

```
./01-servers/IPs/<cluster-name>-NODES
```

Where `<cluster-name>` stands for the `cluster-name` field chosen in the `deploy.conf`
file. Scripts within this suite will look for servers IPs and assets using this name
as prefix.

You can find an example in the `git` commited file `mycluster-NODES`

```
0	11.12.13.14	10.0.0.5
1	11.12.13.20	10.0.0.6
2	11.12.13.32	10.0.0.7
```

The first column represents the `node number`. Scripts will iterate over this number
and management scripts will reference to this number (ex: `node 0`).

The second column stands for the `public IP`, this will be used by the management
scripts to enable the `ssh` connections.

Finally, the third column represents the `private IP`, which is used to create the
inter-node connection (i.e. `enode`), in the bootstraping configuration.
