package main

import (
    "fmt"
    "crypto/ed25519"
    //"encoding/hex"
    "github.com/mr-tron/base58"
)

func main() {
  valid := verify("Ag3m1p1B6FMWKunTQwDW98fLEpcPaobmuthx1u9xLP9q", 
   "3PffhqzduwkmHaS7JCUntPRS8kS25JBZEr1ksp9rWafaVuCt9dd4Y5DuUGXN7caBvaJzTVXP5VRUjdFqF8Fd54Yy", 
   "4PXXepLUr7vyB8ReRgC3qhsDxXTMbAAmh2MBb6DP6FdZ")
  fmt.Printf("valid %v\n", valid)
}

// verify returns true if verified
// ref: https://pkg.go.dev/crypto/ed25519#Verify
func verify(message string, signedMessage string, pubkey string) bool {
    bytes, err := base58.Decode(pubkey)
    if err != nil {
        return false
    }
    messageAsBytes := []byte(message)

    //signedMessageAsBytes, err := hex.DecodeString(signedMessage)
    signedMessageAsBytes, err := base58.Decode(signedMessage)
    if err != nil {
        return false
    }

    return ed25519.Verify(bytes, messageAsBytes, signedMessageAsBytes)
}