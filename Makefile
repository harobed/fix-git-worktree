.PHONY: build
release:
	mkdir -p releases/linux_amd64/
	mkdir -p releases/darwin_amd64/
	GOOS=linux GOARCH=amd64  go build -v -o releases/linux_amd64/fix-git-worktree
	GOOS=darwin GOARCH=amd64  go build -v -o releases/darwin_amd64/fix-git-worktree
