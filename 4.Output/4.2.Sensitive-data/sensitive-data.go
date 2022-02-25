package main

import (
	
	"context"
	

	"golang.org/x/crypto/bcrypt"
)

func main() {
	ctx := context.Background()
	email := []byte("john.doe@somedomain.com")
	password := []byte("47;u5:B(95m72;Xq")

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	stmt, err := stmt.PrepareContext(ctx, "INSERT INTO accounts SET hash=?, email=?")
	if err != nil {
		panic(err)
	}
	result, err := stmt.ExecContext(ctx, hashedPassword, email)
	if err != nil {
		panic(err)
	}
}