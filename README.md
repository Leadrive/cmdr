# cmdr

[![Build Status](https://travis-ci.org/hedzr/cmdr.svg?branch=master)](https://travis-ci.org/hedzr/cmdr)
[![Go Report Card](https://goreportcard.com/badge/github.com/hedzr/cmdr)](https://goreportcard.com/report/github.com/hedzr/cmdr)
![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/hedzr/cmdr.svg?label=release)

getopt/getopt_long like command-line UI golang library.

A getopt-like parser of command-line options, compatible with the [getopt_long](http://www.gnu.org/s/libc/manual/html_node/Argument-Syntax.html#Argument-Syntax) syntax, which is an extension of the syntax recommended by POSIX.

`cmdr` is a UNIX command-line UI library written by golang.


## Features

- Unix [*getopt*(3)](http://man7.org/linux/man-pages/man3/getopt.3.html) representation but without its programmatic interface.

- Automatic help screen generation

- Support for unlimited multiple sub-commands.

- Support for command short and long name, and aliases names.

- Support for both short and long options (`-o` and `--opt`). Support for multiple aliases

- Automatically allows those formats (applied to long flags too):
  - `-I file`, `-Ifile`, and `-I=files`
  - `-I 'file'`, `-I'file'`, and `-I='files'`
  - `-I "file"`, `-I"file"`, and `-I="files"`

- Support for `-D+`, `-D-` to enable/disable a bool flag.

- Support for circuit-break by `--`.

- Support for options being specified multiple times, with different values

- Support for optional arguments.

- Groupable commands and options/flags.

- Sortable commands and options/flags. Or sorted by alphabetic order.

- Bash and Zsh (not yet) completion.

  ```bash
  bin/wget-demo generate shell --bash
  ```

- Predefined yaml config file locations:
  - `/etc/<appname>/<appname>.yml` and `conf.d` sub-directory.
  - `/usr/local/etc/<appname>/<appname>.yml` and `conf.d` sub-directory.
  - `$HOME/<appname>/<appname>,yml` and `conf.d` sub-directory.
  - Watch `conf.d` directory:
    - `AddOnConfigLoadedListener(c)`
    - `RemoveOnConfigLoadedListener(c)`
    - `SetOnConfigLoadedListener(c, enabled)`
  - As a feature, do NOT watch the changes on `<appname>.yml`.

- Overrides by environment variables.

  not yet: prior:

- `cmdr.Getkey)`, `cmdr.GetBool(key)`, `cmdr.GetInt(key)`, `cmdr.GetString(key)`, `cmdr.GetStringSlice(key)` for Option value extraction.

  - bool
  - int, int64, uint, uint64
  - string
  - string slice

- More...




## Examples

1. [**short**](./examples/short/README.md)  
   simple codes.
2. [demo](./examples/demo/README.md)  
   normal demo with external config files.
3. [wget-demo](./examples/wget-demo/README.md)  
   partial-impl wget demo.
   ![](./docs/wget-demo.png)



## LICENSE

MIT.





