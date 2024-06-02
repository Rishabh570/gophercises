package secrets

import "fmt"

type SecretStore struct {
	encodingKey string
}

func SetSecret() {
	fmt.Println("Setting secret: aldsjfsdlidfwse23423")
}

func GetSecret() {
	fmt.Println("returning secret: aldsjfsdlidfwse23423")
}
