package main

import (
	"context"
	"fmt"
	"log"

	"gocloud.dev/secrets/localsecrets"

	_ "gocloud.dev/secrets/gcpkms"
)

func main() {
	ctx := context.Background()

	// Construct a *secrets.Keeper from one of the secrets subpackages.
	// This example uses localsecrets.
	sk, err := localsecrets.NewRandomKey()
	if err != nil {
		log.Fatal(err)
	}
	keeper := localsecrets.NewKeeper(sk)
	defer keeper.Close()

	// Now we can use keeper to Encrypt.
	plaintext := []byte("Go CDK Secrets")
	ciphertext, err := keeper.Encrypt(ctx, plaintext)
	if err != nil {
		log.Fatal(err)
	}

	// And/or Decrypt.
	decrypted, err := keeper.Decrypt(ctx, ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(decrypted))

}
