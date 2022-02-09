<div align="center">
  <a href="https://git-scm.com/">
    <img alt="git" src="../logos/git.png"/>
  </a>
  <h1>Git</h1>
</div>

## Commands

If you're in one branch but pulling from another remote branch. You want to revert back to the original branch you were in.

```sh
git fetch origin
git reset --hard origin/master # if you are in master branch
git clean -f -d
```
