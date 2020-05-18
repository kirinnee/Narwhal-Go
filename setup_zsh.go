package main

import (
	"fmt"
	"os"
)

const ZSH = `#compdef $PROG

_cli_zsh_autocomplete() {

  local -a opts
  local cur
  cur=${words[-1]}
  if [[ "$cur" == "-"* ]]; then
    opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} ${cur} --generate-bash-completion)}")
  else
    opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --generate-bash-completion)}")
  fi

  if [[ "${opts[1]}" != "" ]]; then
    _describe 'values' opts
  else
    _files
  fi

  return
}

compdef _cli_zsh_autocomplete $PROG`

const ZSH_RC = `
PROG=narwhal
_CLI_ZSH_AUTOCOMPLETE_HACK=1
source  ~/.narwhalrc
PROG=nw
source  ~/.narwhalrc
`

func setupZSH() error {
	err := write(".narwhalrc", BASH, true)
	if err != nil {
		return err
	}

	err = clearExistingZsh()
	if err != nil {
		return err
	}
	err = appendTo(".zshrc", ZSH_RC, true)
	if err != nil {
		return err
	}
	fmt.Println("Please either restart your shell or run:")
	fmt.Println("\tsource ~/.zshrc")
	return nil
}

func clearExistingZsh() error {
	profile, err := homeFile(".zshrc")
	if err != nil {
		return err
	}
	return removeFromFile(profile, ZSH_RC)

}

func tearDownZsh() error {

	file, err := homeFile(".narwhalrc")
	if err != nil {
		return err
	}
	err = os.Remove(file)
	if err != nil {
		fmt.Println(err)
	}
	err = clearExistingZsh()
	if err != nil {
		return err
	}
	return nil
}
