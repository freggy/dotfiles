Dotfiles
========

This repository contains all configuration files and scripts needed to bootstrap my workspace and keep it up-to-date.
The layout is simple: It's basically the standard linux directory structure. If I need a file at `/usr/local/bin` I 
put it in `usr/local/bin` in the repositories root folder.

Command overview
----------------

```
$ ./bootstrap.sh # inital bootstrap
$ ./bootstrap.sh --update # fetch most recent files and update them
```
