# nslookupper

## Install

```
go get github.com/akm/nslookupper
```

Or download from https://github.com/akm/nslookupper/releases

## Usage

```
NAME:
   nslookupper - github.com/akm/nslookupper

USAGE:
   nslookupper [GLOBAL OPTIONS] HOST_NAME

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --name-server value  Name server to look up (default: "8.8.8.8")
   --help, -h           show help
   --version, -v        print the version
```

## Example

Fetch source ranges of Google App Engine.

```
$ nslookupper _cloud-netblocks.googleusercontent.com
35.190.224.0/20
35.232.0.0/15
35.234.0.0/16
35.235.0.0/17
(snip)
```

See `Static IP Addresses and App Engine apps` in https://cloud.google.com/appengine/kb/ for more detail
