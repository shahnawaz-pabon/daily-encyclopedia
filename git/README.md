<div align="center">
  <a href="https://git-scm.com/">
    <img alt="git" src="../logos/git.png"/>
  </a>
  <h1>Git</h1>
</div>

# Table of Contents

- [Awesome github profiles](https://zzetao.github.io/awesome-github-profile/)
- [Revert back to the original branch](#revert-back-to-the-original-branch)
- [Delete a local branch](#delete-a-local-branch)
- [Delete a remote branch](#delete-a-remote-branch)
- [Set git editor to vim](#set-git-editor-to-vim)
- [Save username and password](#save-username-and-password)
- [Removes latest commit from the stash, KEEPS changes](#removes-latest-commit-from-the-stash-keeps-changes)
- [Removes latest commit from the stash, DELETES changes](#removes-latest-commit-from-the-stash-deletes-changes)
- [Set remote url with access token](#set-remote-url-with-access-token)
- [Changing git commit message after push](#changing-git-commit-message-after-push)
- [Change the url of a remote git repository](#change-the-url-of-a-remote-git-repository)
- [Remove a big file from any commit](#remove-a-big-file-from-any-commit)

  <br>

## Revert back to the original branch

If you're in one branch but pulling from another remote branch. You want to revert back to the original branch you were in.

```sh
git fetch origin
git reset --hard origin/master # if you are in master branch
git clean -f -d
```

<br>

## Delete a local branch

```sh
git branch -d branch_name
```

<br>

## Delete a remote branch

```sh
git push origin --delete remoteBranchName
```

<br>

## Reset speficific file

Revert back to the last commit you were in for a specific file

```sh
git restore `filepath`
```

<br>

## Set git editor to vim

```sh
git config --global core.editor "vim"
```

<br>

## Replace all occurences of `old` words by a `new` word in the file

```sh
:%s/old/new/g
```

## Save username and password

```sh
git config --global credential.helper store
```

<br>

## Removes latest commit from the stash, KEEPS changes

```sh
git reset --soft HEAD~
```

<br>

## Removes latest commit from the stash, DELETES changes

```sh
git reset --hard HEAD~
```

<br>

## Set remote url with access token

Git has removed their password authentication on August 13, 2021. So to push to your remote repository, you can do the following:

```sh
# see the origin's url
git remote -v
# remove origin's url from the git config
git remote remove origin
git remote add origin https://<githubtoken>@github.com/<username>/<repositoryname>.git
```

<br>

## Changing git commit message after push

```sh
git commit --amend -m "New commit message"
# You should probably refrain from erasing any modifications that someone else may have submitted to the same branch. The option —force-with-lease is the safest since it will terminate if any upstream modifications are made.
git push --force-with-lease <repository> <branch>
```

<br>

## Change the url of a remote git repository

```sh
git remote set-url origin <new_remote_repository_url>
```

<br>


## Remove a big file from any commit

```sh
git filter-branch --force --index-filter "git rm --cached --ignore-unmatch yourfilepath/bigfile.csv" -- --all
```