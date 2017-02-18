package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/p2p/discover"
)

/*
	Usage

	go run nodekey.go --privkey $(perl -e '@c=("a".."f",0..9);$p.=$c[rand(scalar @c)] for 1..64; print "$p\n"')

	Output sample

	d2374f946df93cde864505aee99122e4ad0040061cca668a19ec2b369448c137
	844c9c9e926a96e67fae7124bf6fcb6ecf37d121e2d3031db4b7d7bdd1388d9fc33c96c70535c65fcf34d8a4258fd40a9a7e2c24ac92bb152bd0261464b845d1
*/

func main() {
	privKeyString := flag.String("privkey", "", "Private key string in hex")
	flag.Parse()

	if *privKeyString == "" {
		fmt.Println("ERROR: You need to enter --privkey")
		return
	}

	if len(*privKeyString) != 64 {
		fmt.Println("ERROR: You need to enter a 512-bit number (64 hex chars)")
		return
	}

	privHex, err := hex.DecodeString(*privKeyString)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = secp256k1.S256()
	priv.D = common.BigD(privHex)
	priv.PublicKey.X, priv.PublicKey.Y = secp256k1.S256().ScalarBaseMult(privHex)

	fmt.Printf("%v\n%v\n", *privKeyString, discover.PubkeyID(&priv.PublicKey))
}
