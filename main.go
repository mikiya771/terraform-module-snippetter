package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	cs "github.com/mikiya771/terraform-module-snippetter/internal/create_snip"
	cra "github.com/mikiya771/terraform-module-snippetter/internal/module_repo_analyzer"
)

func main() {
	homeDir, _ := os.UserHomeDir()
	cacheDir := path.Join(homeDir, ".cache")
	orgName := "terraform-aws-modules"
	repoName := "terraform-aws-vpc"
	url := fmt.Sprintf("https://github.com/%s/%s", orgName, repoName)
	modDir := path.Join(cacheDir, orgName, repoName)
	if _, err := os.Stat(modDir); os.IsNotExist(err) {
		fmt.Println("file doesn't exist")
		fmt.Println(cacheDir)
		cmd1 := exec.Command("mkdir", "-p", path.Join(cacheDir, orgName))
		cmd1.Stdout = os.Stdout
		cmd1.Stderr = os.Stderr
		err = cmd1.Run()
		if err != nil {
			log.Fatal(err)
		}
		cmd2 := exec.Command("git", "clone", url, modDir)
		cmd2.Stdout = os.Stdout
		cmd2.Stderr = os.Stderr

		err = cmd2.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	mis := cra.ModuleRepoAnalyzer(cacheDir, orgName, repoName)
	var b bytes.Buffer
	for _, mi := range mis {
		cs.CreateSnip(&b, url, modDir, mi.ModulePath, mi.GetTitle(), mi.GetDescription())
	}

	fmt.Println(b.String())
}
