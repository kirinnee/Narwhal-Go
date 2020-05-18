package main

import (
	"fmt"
	"os"
)

const BASH = `#! /bin/bash

: ${PROG:=$(basename ${BASH_SOURCE})}

_cli_bash_autocomplete() {
  if [[ "${COMP_WORDS[0]}" != "source" ]]; then
    local cur opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    if [[ "$cur" == "-"* ]]; then
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} ${cur} --generate-bash-completion )
    else
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
    fi
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
  fi
}

complete -o bashdefault -o default -o nospace -F _cli_bash_autocomplete $PROG
unset PROG`

func setupBash() error {
	err := write("/etc/bash_completion.d/nw", BASH, false)
	if err != nil {
		return err
	}
	err = write("/etc/bash_completion.d/narwhal", BASH, false)
	if err != nil {
		return err
	}
	fmt.Println("Please either restart your shell or run:")
	fmt.Println("\tsource ~/.bashrc")
	return nil
}

func tearDownBash() error {

	err := os.Remove("/etc/bash_completion.d/nw")
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove("/etc/bash_completion.d/narwhal")
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
