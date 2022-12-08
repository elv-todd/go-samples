// https://flowscan.org/account/0xcbd420284fd5e19b
// 


package main

import (
	"context"
	"fmt"

	"github.com/onflow/flow-go-sdk/client"

	"google.golang.org/grpc"
)

func scriptGetToken(addrOwner string) string {
	return `
import FlowServiceAccount from 0xe467b9dd11fa00df

pub fun main(): UFix64 {
    let account = getAccount(` + addrOwner + `)

    let balance = FlowServiceAccount.defaultTokenBalance(account)

    return balance
}
`
}

func scriptGetNrNFT(addrOwner string, addrNFT string, nameNFT string) string {
	return `
import ` + nameNFT + ` from ` + addrNFT + `

pub fun main(): Int {
    let nftOwner = getAccount(` + addrOwner + `)

    // Find the public Receiver capability for their Collection
    let capability = nftOwner.getCapability<&{` + nameNFT + `.` + nameNFT + `CollectionPublic}>(` + nameNFT + `.CollectionPublicPath)

    // borrow a reference from the capability
    let receiverRef = capability.borrow()
            ?? panic("Does not have some NFT of this collection")

    return receiverRef.getIDs().length
}
`
}

func ExecuteScript(node string, script []byte) {
	ctx := context.Background()
	c, err := client.New(node, grpc.WithInsecure())
	if err != nil {
		panic("failed to connect to node")
	}

	result, err := c.ExecuteScriptAtLatestBlock(ctx, script, nil)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("result:", result)
}

func main() {
	node := "access.mainnet.nodes.onflow.org:9000"

	fmt.Println("compare https://flowscan.org/account/0xcbd420284fd5e19b")
	fmt.Println("get account->balance:")
	ExecuteScript(node, []byte(scriptGetToken("0xcbd420284fd5e19b")))
	fmt.Println("get account->capability->IDs.len:")
	ExecuteScript(node, []byte(scriptGetNrNFT("0xcbd420284fd5e19b", "0x329feb3ab062d289", "CNN_NFT")))
}