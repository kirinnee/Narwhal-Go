package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

const POWER_SHELL = `$fn = $($MyInvocation.MyCommand.Name)
$name = $fn -replace "(.*)\.ps1$", '$1'
Register-ArgumentCompleter -Native -CommandName $name -ScriptBlock {
     param($commandName, $wordToComplete, $cursorPosition)
     $other = "$wordToComplete --generate-bash-completion"
         Invoke-Expression $other | ForEach-Object {
            [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
         }
 }`

const PS_PROFILE_1 = "\r\n& $(\"$(Split-Path -Path $profile)/AutoComplete/narwhal.ps1\")\r\n"
const PS_PROFILE_2 = "\r\n& $(\"$(Split-Path -Path $profile)/AutoComplete/nw.ps1\")\r\n"

func setupPowerShell(profile string) error {
	// Create AutoComplete Folder
	dir := filepath.Dir(profile)
	folder := path.Join(dir, "AutoComplete")
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return err
	}

	err = write(path.Join(folder, "narwhal.ps1"), POWER_SHELL, false)
	if err != nil {
		return err
	}

	err = write(path.Join(folder, "nw.ps1"), POWER_SHELL, false)
	if err != nil {
		return err
	}

	err = clearExisting(profile)
	if err != nil {
		return err
	}

	err = appendTo(profile, PS_PROFILE_1, false)
	if err != nil {
		return err
	}
	err = appendTo(profile, PS_PROFILE_2, false)
	if err != nil {
		return err
	}

	fmt.Println("Please either restart your shell or run:")
	fmt.Println("\t& $profile")
	return nil
}

func clearExisting(profile string) error {
	err := removeFromFile(profile, PS_PROFILE_1)
	if err != nil {
		return err
	}
	return removeFromFile(profile, PS_PROFILE_2)
}

func tearDownPowerShell(profile string) error {
	dir := filepath.Dir(profile)
	folder := path.Join(dir, "AutoComplete")
	err := os.Remove(path.Join(folder, "narwhal.ps1"))
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove(path.Join(folder, "nw.ps1"))
	if err != nil {
		fmt.Println(err)
	}
	return clearExisting(profile)
}
