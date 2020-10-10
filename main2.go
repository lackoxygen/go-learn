package main

import "fmt"

const(
	DB_HIST string = "127.0.0.1"
	DB_PORT string = "3306"
	DB_USERNAME string = "root"
	DB_PASSWORD string = "123456"
	DB_DATABASE = "forget"
)

const X  = 1

const Y  = 2

func main()  {
	gog := 1

	label:

	fmt.Println("DB_HOST="+DB_HIST + "\n" +
		"DB_PORT="+DB_PORT+"\n" +
		"DB_USERNAME="+DB_USERNAME+"\n" +
		"DB_PASSWORD="+DB_PASSWORD+"\n" +
		"DB_DATABASE="+DB_DATABASE)

	gog += 1

	fmt.Println("x * y", X * Y)

	if gog < 2 {
		goto label
	}
}
