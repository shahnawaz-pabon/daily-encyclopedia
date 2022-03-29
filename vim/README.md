<div align="center">
  <a href="https://www.vim.org/">
    <img alt="vim" src="../logos/vim.png"/>
  </a>
  <h1>Vim</h1>
</div>

## Commands

### Jumps to the end of the file

```sh
:$
```

<br>

### Enters insert mode at end of current line

```sh
Shift + A
```

### Customize your vim with .vimrc

Create a file `.vimrc` in your home directory, write the following code and save it:

```sh
" Disable compatibility with vi which can cause unexpected issues.
set nocompatible
" Enable type file detection. Vim will be able to try to detect the type of file in use.
filetype on
" Enable plugins and load plugin for the detected file type.
filetype plugin on
" Load an indent file for the detected file type.
filetype indent on
" Turn syntax highlighting on.
syntax on
" Add numbers to each line on the left-hand side.
set number
" Highlight cursor line underneath the cursor horizontally.
set cursorline
" Highlight cursor line underneath the cursor vertically.
set cursorcolumn
" Set shift width to 4 spaces.
set shiftwidth=4
" Set tab width to 4 columns.
set tabstop=4
" Use space characters instead of tabs.
set expandtab
" Do not save backup files.
set nobackup
" Do not let cursor scroll below or above N number of lines when scrolling.
set scrolloff=10
" Do not wrap lines. Allow long lines to extend as far as the line goes.
set nowrap
" While searching though a file incrementally highlight matching characters as you type.
set incsearch
" Ignore capital letters during search.
set ignorecase
" Override the ignorecase option if searching for capital letters.
" This will allow you to search specifically for capital letters.
set smartcase
" Show partial command you type in the last line of the screen.
set showcmd
" Show the mode you are on the last line.
set showmode
" Show matching words during a search.
set showmatch
" Use highlighting when doing a search.
set hlsearch
" Set the commands to save in history default number is 20.
set history=1000
```

Then add the following line in your `~/.bashrc` file (if not found, create the file in your home directory)

```sh
alias vim="vim -S ~/.vimrc"
```

### Delete a single line

Bring your cursor to the line you want to delete, then write `:d` and press `Enter`. Or simply you can quickly press `dd` to delete the current line.

### Delete all the lines

```sh
:%d
```
