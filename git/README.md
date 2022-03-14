<div align="center">
  <a href="https://git-scm.com/">
    <img alt="git" src="../logos/git.png"/>
  </a>
  <h1>Git</h1>
</div>

# Table of Contents

- [Revert back to the original branch](#revert-back-to-the-original-branch)
- [Delete a local branch](#delete-a-local-branch)
- [Delete a remote branch](#delete-a-remote-branch)

## Revert back to the original branch

If you're in one branch but pulling from another remote branch. You want to revert back to the original branch you were in.

```sh
git fetch origin
git reset --hard origin/master # if you are in master branch
git clean -f -d
```

## Delete a local branch

```sh
git branch -d branch_name
```

## Delete a remote branch

```sh
git push origin --delete remoteBranchName
```

## Reset speficific file

Revert back to the last commit you were in for a specific file

```sh
git restore `filepath`
```
