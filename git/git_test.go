package git_test

import (
	"os"
	"testing"

	"github.com/getantibody/antibody/git"
	"github.com/getantibody/antibody/internal"
	"github.com/stretchr/testify/assert"
)

func TestClonesRepo(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGitRepo("caarlos0/env", home)
	assert.NoError(t, repo.Download())
	internal.AssertFileCount(t, 1, home)
}

func TestUpdatesRepo(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGitRepo("caarlos0/zsh-pg", home)
	assert.NoError(t, repo.Download())
	assert.NoError(t, repo.Update())
	internal.AssertFileCount(t, 1, home)
}

func TestCloneDoesNothingIfFolderAlreadyExists(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGitRepo("caarlos0/zsh-add-upstream", home)
	assert.NoError(t, repo.Download())
	assert.NoError(t, repo.Download())
	internal.AssertFileCount(t, 1, home)
}

func TestClonesUnexistentRepo(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGitRepo("doesn-not-exist-really", home)
	assert.Error(t, repo.Download())
	internal.AssertFileCount(t, 0, home)
}

func TestUpdatesUnexistentRepo(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGitRepo("doesn-not-exist-really", home)
	assert.Error(t, repo.Update())
	internal.AssertFileCount(t, 0, home)
}

func TestGetsRepoInfo(t *testing.T) {
	home := internal.TempHome()
	defer os.RemoveAll(home)
	repo := git.NewGitRepo("caarlos0/zsh-pg", home)
	assert.Equal(t, "caarlos0/zsh-pg", repo.Name())
	assert.Equal(t, home+"https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-zsh-pg", repo.Folder())
}
