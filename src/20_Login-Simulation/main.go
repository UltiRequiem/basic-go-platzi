package main

import "fmt"

func main() {
	fmt.Println("El estado es:", login("UltiRequiem", "contraseña muy segura"))
}

func login(username, password string) bool {
	const usernameStorage string = "UltiRequiem"
	const passwordStorage string = "contraseña muy segura"

	if username == usernameStorage && password == passwordStorage {
		return true
	} else {
		return false
	}
}
