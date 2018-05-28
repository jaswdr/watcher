# Watcher
> Watch files and directories and run command when changes happen

[![Build Status](https://travis-ci.org/jaswdr/watcher.svg?branch=master)](https://travis-ci.org/jaswdr/watcher)

### Instalation

Download on the [releases](https://github.com/jaswdr/watch/releases) page.

### Usage

```sh
$ watch "ls" .
```
Each time a file is updated the command will be run.

### Development

Clone the repository.

```sh
$ git clone https://github.com/jaswdr/watch
```

Install dependencies.

```sh
$ make dep
```

(Optional) Make your own snapshot

```sh
$ make snapshot
```
