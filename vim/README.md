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

### Delete a range of lines

`:[Starting Line Number],[Ending Line Number]d`

```sh
# Line number 5,6,7 will be removed
:5,7d
```

### Delete the line with a custom pattern

`:g /<WORD>/d`

```sh
# Lines containing the "linux" word will be deleted
:g /linux/d
```

### Delete all blank lines

```sh
:g /^$/d
```

### Delete lines that start with a specific letter

```sh
# Delete all lines that start with letter "S"
:g /^S/d
```

### Delete all lines after line number

```sh
# If your cursor is on line number 5, then this command will delete all lines below line number 5
:+1,$d
```

### Delete all lines before line number

```sh
# If your cursor is on line number 5, then this command will delete all lines before line number 5
:1,-1d
```

## Copy the current line

```sh
yy
```

## Copy selected text to the clipboard

```sh
y
```

## Cut current line

```sh
dd
```

## Copy to end of line

```sh
y$
```

## Cut to end of line

```sh
D
```

## Paste clipboard contents

```sh
p
```

## Search word from the top to bottom

```sh
/word
```

## Search word from the bottom to top

```sh
?word
```

## Search the word under the cursor

```sh
*
```

## Search a string, case insensitive

```sh
/\cstring
```

## Adding quotes at the start and end of each line

```sh
# this command will add quotes at the start and end of each line with comma at the end
:%s/\(.*\)/'\1',
```

## Show difference between two files

```sh
vimdiff file1.txt file2.txt
```
