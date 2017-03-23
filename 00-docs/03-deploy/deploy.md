# Deploy

Deploy tools are the elements that will create the `assets` (as, for example,
`nodekeys`, `genesis files` or `constellation keys`), as well as the tools
to deploy those assets, along with the `geth` and `constellation`
executables, and their launch commands.

## Assets directory

It is local and `.gitignore`'d. It is placed using the `cluster-name` variable
in the directory

```
./03-deploy/00-assets/<cluster-name>
```

Make sure, after you generate your assets, to backup this directory!

## Assets preparation

### Ethereum address

This script

```
./03-deploy/01-prepare-assets/address
```

Will invoke the image `infura/quorum-tools-assets`, create and dump in the terminal
a new ethereum address.

Example

```
$ ./03-deploy/01-prepare-assets/address
{"address":"af390cb51a5f03d45201e82d670a5ea577fe9604","crypto":{"cipher":"aes-128-ctr","ciphertext":"e68ef0bae29497fa528736d1f71d47d62bd64bc98a961ab2213385cf36b625fd","cipherparams":{"iv":"bd4bdce8987ca48e2e5623c04365e40b"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"fb9f7fd2e1786f4f0e85dede74916b1fd3710a68b10986cbc547314158452300"},"mac":"408ebb120bfc2adffcc3b40eb2cb26f690144ed471c9c65bc42502472ceba703"},"id":"6d7ea228-c21f-4d9f-a6d7-4fccd47a117f","version":3}
```

You can, of course, pipe the output of this command

```
./.03/deploy/01-prepare-assets/address > /tmp/my-address
```

### Constellation keys

This script

```
./03-deploy/01-prepare-assets/constellation
```

Will invoke the image `infura/quorum-tools-assets`, create and dump in the terminal
a new constellation prv/pub pair.

Example

```
$ 03-deploy/01-prepare-assets/constellation
5JIgWVCHBJ3ObqN+rr2/OINxuyaijXVHMn7/Q0MU804=
{"data":{"bytes":"q+5prDHbxo+GWaREhWGLa2lqmJD96Wnn5O5DO28FBLE="},"type":"unlocked"}
```

To pipe the output into files, do, inside of other script, or by separate lines

```
$ RESULT=($(./03-deploy/01-prepare-assets/constellation))
$ echo ${RESULT[0]} > /tmp/1-pub
$ echo ${RESULT[1]} > /tmp/1-prv
```

### Nodekey

This script

```
./03-deploy/01-prepare-assets/node
```

Will invoke the image `infura/quorum-tools-assets`, create and dump in the terminal
an ECDSA public/private pair. The public element being using to define a node's `enode`,
while the private element can be used as a `nodekey` for an ethereum node.

Example

```
 ./03-deploy/01-prepare-assets/node
8d357ba9f20d7b78511b1555ef35e11ec54103066ebd2cc492058d9b55434be4
3e0c6b75e650c8bb640b33736ec6a319765fb97b46ef032d775cfd79f2cf737d0045e50b9bc987d4611d31b960b228f26b9795df3cad12b5c32ed8c027e2ba23
```

*NOTE:* `/dev/random` is used as the random generator.

To pipe the output into files, do, inside of other script, or by separate lines

```
$ RESULT=($(03-deploy/01-prepare-assets/node))
$ echo "${RESULT[0]}" > /tmp/1-nodekey
$ echo "${RESULT[1]}" > /tmp/1-enode
```

### Genesis file

This script

```
./03-deploy/01-prepare-assets/genesis
```


Will invoke the image `infura/quorum-tools-assets`, create and dump in the terminal a
`genesis file` for the use of the blockchain.

*WARNING*: You need to have in your assets directory a `genesis-source.json` file
in order to this script to work. You can build one by reading the sample located at
`03-deploy/01-prepare-assets/genesis-source.sample`.

Example

```
./03-deploy/01-prepare-assets/genesis
{
  "alloc": {
    "0x0000000000000000000000000000000000000020": {
      "code": "60606040
...
  },
  "difficulty": "0x0",
  "extraData": "0x",
  "gasLimit": "0x2FEFD800",
  "mixhash": "0x00000000000000000000000000000000000000647572616c65787365646c6578",
  "nonce": "0x0",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp": "0x00"
}
```

This one is really easy to pipe into a file ;-)

## Uploading

### Assets

(TODO)

### Geth

(TODO)

### Constellation

(TODO)

## Launch

(TODO)
