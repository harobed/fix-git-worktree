package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func createTestData() string {
	current_path, _ := os.Getwd()

	path := filepath.Join(current_path, "test_datas")
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}

	tarpath, err := exec.LookPath("tar")
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command(tarpath, "xfz", "test_datas.tar.gz").Run()
	if err != nil {
		log.Fatal(err)
	}

	return path
}

func TestMain(t *testing.T) {
	path := createTestData()

	ConvertWorktreeToRelativePath(filepath.Join(path, "master"), false)
	gitdir_value, err := ioutil.ReadFile(filepath.Join(path, "master", ".git", "worktrees", "develop", "gitdir"))
	if err != nil {
		log.Fatal(err)
	}
	if string(gitdir_value) != "../develop/.git\n" {
		t.Error("Expected \"../develop/.git\", get ", string(gitdir_value))
	}

	worktree_gitdir_value, err := ioutil.ReadFile(filepath.Join(path, "develop", ".git"))
	if err != nil {
		log.Fatal(err)
	}
	if string(worktree_gitdir_value) != "gitdir: ../master/.git/worktrees/develop\n" {
		t.Error("Expected \"gitdir: ../master/.git/worktrees/develop\", get ", string(worktree_gitdir_value))
	}
}
