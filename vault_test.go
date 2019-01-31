package vault

import (
	"fmt"
	"testing"
)

func TestNewVaultCreation(t *testing.T) {
	vault := CreateVault("./sample.vault", "sample")
	fmt.Println(vault)
	vault.Close()
}

func TestExistingVaultAccess(t *testing.T) {
	vault, err := InitVault("./sample.vault", "sample")
	if err != nil {
		fmt.Println("error = ", fmt.Sprintf("%s", err))
	}
	fmt.Println(vault)
	vault.Close()
}

func TestDataAdditionAndFetchingDataFromVault(t *testing.T) {
	vault, err := InitVault("./sample.vault", "sample")
	if err != nil {
		fmt.Println("error = ", fmt.Sprintf("%s", err))
	}
	err = vault.Upsert("demo-key", "demo-demo")
	if err != nil {
		fmt.Println("error = ", fmt.Sprintf("%s", err))
	}
	data, err := vault.Get("demo-key")
	if err != nil {
		fmt.Println("error = ", fmt.Sprintf("%s", err))
	} else {
		fmt.Println("data = ", string(data))
	}
	vault.Close()
}

func TestDeleteDataFromVault(t *testing.T) {
	vault, err := InitVault("./testing.vault", "demo")
	if err != nil {
		fmt.Println("error = ", fmt.Sprintf("%s", err))
	}
	data, err := vault.Get("demo-key")
	if err != nil {
		fmt.Println("error = ", fmt.Sprintf("%s", err))
	} else {
		fmt.Println("data = ", string(data))
	}
	err = vault.Delete("demo-key")
	if err != nil {
		fmt.Println("error = ", fmt.Sprintf("%s", err))
	}
	data, err = vault.Get("demo-key")
	if err != nil {
		fmt.Println("error = ", fmt.Sprintf("%s", err))
	} else {
		fmt.Println("data = ", string(data))
	}
	vault.Close()
}

func TestChangePasswordOfVault(t *testing.T) {
	vault, err := InitVault("./testing.vault", "sample")
	if err != nil {
		fmt.Println("error = ", fmt.Sprintf("%s", err))
	}
	err = vault.Upsert("demo-key", "demo-demo")
	if err != nil {
		fmt.Println("error = ", fmt.Sprintf("%s", err))
	}
	err = vault.ChangePassword("demo")
	if err != nil {
		fmt.Println("error = ", fmt.Sprintf("%s", err))
	}
	vault.Close()
	vault, err = InitVault("./testing.vault", "demo")
	if err != nil {
		fmt.Println("error = ", fmt.Sprintf("%s", err))
	}
	data, err := vault.Get("demo-key")
	if err != nil {
		fmt.Println("error = ", fmt.Sprintf("%s", err))
	} else {
		fmt.Println("data = ", string(data))
	}
	vault.Close()
}
