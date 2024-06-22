# godoclite

`little "godoc" emulator that extend "go doc" with "go/doc" pkg usage`

Help to check if documentation to [pkg.go.dev](https://pkg.go.dev) is ok on the fly.

Follow the comments rules at [tip.golang.org/doc/comment](https://tip.golang.org/doc/comment)

## why godoclite?

- "go doc" only show documentation on a single mode at one time

- "godoc" only show documentation found in gopath, is necessary use "go get" to download module with packages every time documentation comments change

- godoclite approach can be a isolated documentation server for projects

- godoclite may help to check if documentation to hold in pkg.go.dev is ok before publish

## installation

```sh
go install github.com/zeroidentidad/godoclite@latest
```

## usage

```sh
# show usage help:
godoclite -h

#output:
Usage of godoclite:
  -pkg string
        Path to the Go package (default ".")
  -port string
        Port to serve documentation (default "8080")
```

```sh
# run binary with flags:
godoclite -pkg=./path/to/your/package -port=8080

# run binary without flags on package dir:
godoclite
```

## sample result

![gif](https://raw.githubusercontent.com/zeroidentidad/zeroidentidad/master/img/sample.gif)