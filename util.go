package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func appendFile(file, text string) error {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		return err
	}
	return nil
}

func homeFile(s string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(home, s), nil
}

func removeFromFile(file, rm string) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		_, err = os.Create(file)
		if err != nil {
			return err
		}
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	text := string(b)
	if strings.Contains(text, rm) {
		text = strings.Replace(text, rm, "", -1)
	}
	return write(file, text, false)
}

func appendTo(file, text string, home bool) error {
	if home {
		var err error
		file, err = homeFile(file)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		_, err = os.Create(file)
		if err != nil {
			return err
		}
	}
	return appendFile(file, text)
}

func write(file, text string, home bool) error {
	if home {
		var err error
		file, err = homeFile(file)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(file, []byte(text), 0644)
}
