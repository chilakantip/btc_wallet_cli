package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func ShaHash(b []byte, out []byte) {
	s := sha256.New()
	s.Write(b[:])
	tmp := s.Sum(nil)
	s.Reset()
	s.Write(tmp)
	copy(out[:], s.Sum(nil))
}

func getSeedFromCli() (seed string, err error) {
	fmt.Print("Enter seed: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return
	}
	password := string(bytePassword)

	fmt.Print("\nRe-enter seed: ")
	bytePasswordAgain, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return
	}
	passwordAgain := string(bytePasswordAgain)

	if password != passwordAgain {
		fmt.Println("\npassword not match")
		return "", fmt.Errorf("password not match")
	}

	return strings.TrimSpace(password), nil
}

func exportKeys(key string) error {
	return ioutil.WriteFile(exportFileName, []byte(key), 0400)
}

func getLineFromCmdLine() (line string, err error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}
