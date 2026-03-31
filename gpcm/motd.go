package gpcm

import (
	"errors"
	"os"
)

var motdFilepath = "./motd.txt"
var motd string = ""

var (
	ErrEmptyMotd   = errors.New("motd cannot be empty")
	ErrMotdTooLong = errors.New("motd is too long, max motd is 255 characters")
)

func GetMessageOfTheDay() (string, error) {
	if motd == "" {
		contents, err := os.ReadFile(motdFilepath)
		if err != nil {
			return "", err
		}

		motd = string(contents)
	}

	return motd, nil
}

func SetMessageOfTheDay(nmotd string) error {
	if nmotd == "" {
		return ErrEmptyMotd
	}

	if len(nmotd) > 255 {
		return ErrMotdTooLong
	}

	err := os.WriteFile(motdFilepath, []byte(nmotd), 0644)
	motd = nmotd

	return err
}
