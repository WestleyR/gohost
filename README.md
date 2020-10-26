# GoHost

Simple Go program to host files on your local network

## Install

```
go get github.com/WestleyR/gohost
```

Or download the repository and run `go build` and copy the
`gohost` binary to one of your `PATH` directories.

## Example

Host the current directory:

```
gohost
```

![Alt text](/gohost-screenshot.png?raw=true "gohost screenshot")

By default, this will serve on port `8080` (only accessible
to your machine). To change the port, use the `-p <port>` flag:

```
gohost -p 80
```

Now your current directory should be accessible to other
machines.

