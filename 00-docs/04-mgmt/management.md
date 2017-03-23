# Management Tools

You should be able to perform all your operations from your console.

The file `deploy.conf` needs to be updated!

## Bash Node

Has two uses: Logging into a node via `ssh` and performing a remote command.

```
./03-mgmt/bash-node 0
```

Log into the node `0` (specified in `01-servers/IPs`.

```
./03-mgmt/bash-node 0 "ps aux | grep geth"
```

Will run the command `ps aux | grep geth` in the node `0` and return its result.
