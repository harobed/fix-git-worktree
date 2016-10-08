package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ConvertWorktreeToRelativePath(path string, verbose bool) {
	files, err := ioutil.ReadDir(filepath.Join(path, ".git", "worktrees"))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		gitdir_file_path := filepath.Join(path, ".git", "worktrees", file.Name(), "gitdir")
		gitdir_value, err := ioutil.ReadFile(gitdir_file_path)
		if err != nil {
			log.Fatal(err)
		}
		if strings.HasPrefix(string(gitdir_value), "/") {
			rel, err := filepath.Rel(path, string(gitdir_value))
			err = ioutil.WriteFile(gitdir_file_path, []byte(rel), 0644)
			if verbose {
				fmt.Printf("Fix \"%s\"\n", gitdir_file_path)
			}
			if err != nil {
				log.Fatal(err)
			}
		}

		worktree_gitdir_path := strings.TrimSpace(string(gitdir_value))
		if !strings.HasPrefix(worktree_gitdir_path, "/") {
			worktree_gitdir_path = filepath.Join(path, worktree_gitdir_path)
		}
		worktree_gitdir_value, err := ioutil.ReadFile(worktree_gitdir_path)
		if err != nil {
			log.Fatal(err)
		}
		var worktree_gitdir_field_value string

		infos := strings.Split(string(worktree_gitdir_value), "\n")
		for _, info := range infos {
			if strings.HasPrefix(info, "gitdir: ") {
				worktree_gitdir_field_value = info[8:]
				break
			}
		}
		if strings.HasPrefix(string(worktree_gitdir_field_value), "/") {
			rel, err := filepath.Rel(filepath.Dir(worktree_gitdir_path), string(worktree_gitdir_field_value))
			worktree_gitdir_value = []byte(strings.Replace(
				string(worktree_gitdir_value),
				fmt.Sprintf("gitdir: %s", worktree_gitdir_field_value),
				fmt.Sprintf("gitdir: %s", rel),
				1,
			))
			err = ioutil.WriteFile(worktree_gitdir_path, worktree_gitdir_value, 0644)
			if verbose {
				fmt.Printf("Fix \"%s\"\n", worktree_gitdir_path)
			}
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: fix-git-worktree <git_path>\n\n")
		flag.PrintDefaults()
	}
	verbose := flag.Bool("verbose", false, "verbose mode")

	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	path, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("Error: %s not found\n", path)
		os.Exit(1)
	}

	_, err = os.Stat(filepath.Join(path, ".git"))
	if os.IsNotExist(err) {
		fmt.Printf("Error: %s isn't git folder\n", path)
		os.Exit(1)
	}
	fmt.Printf("Convert worktree to relative path in \"%s\" git repository\n", path)
	ConvertWorktreeToRelativePath(path, *verbose)
	fmt.Printf("Done\n")
}
