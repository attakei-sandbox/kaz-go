# kaz

## Overview

**kaz** is simle Appimage management tool.

## Description

kaz has been created to manage installed appimage.
It can manage appimage by like brew, yum or apt.

## Installation

```sh
$ go install github.com/attakei/kaz
```

## Usage

```sh
# Initialize einvironment
$ kaz init

# Add app as target into local repository
$ kaz add hyper --vendor=github --args zeit/hyper

# Install appimage
$ kaz install hyper
```

## License

BSD License. See [it](./LICENSE)

