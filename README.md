# Tool to convert git worktree config files from absolute path to relative

This tool fix this issue: [Feature request: use relative path in worktree config files](http://marc.info/?l=git&m=147591940221454&w=4)


## Usage

```
$ ./fix-git-worktree
Usage: fix-git-worktree <git_path>

  -verbose
    	verbose mode
```


## Build

```
$ go build -v -o fix-git-worktree
```


## Make release files

```
$ make release
```


## Execute tests

```
$ go test -v
```
