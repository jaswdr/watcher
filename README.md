# Watcher
> Watch files and directories and run command when changes happen

[![Build Status](https://travis-ci.org/jaswdr/watcher.svg?branch=master)](https://travis-ci.org/jaswdr/watcher)

### Instalation

Download on the [releases](https://github.com/jaswdr/watcher/releases) page.

or you can use [GoBinaries](https://gobinaries.com/) and install it from command line.

```bash
curl -sf https://gobinaries.com/jaswdr/watcher | sh
```

### Usage

```sh
$ watch "ls" .
```
Each time a file is updated the command will be run.

### Development

Clone the repository.

```sh
$ git clone https://github.com/jaswdr/watcher
```

Install dependencies.

```sh
$ make dep
```

(Optional) Make your own snapshot

```sh
$ make snapshot
```
