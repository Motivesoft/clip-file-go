# clip-file-go
Go implementation of clip-file, to read a text file and place its contents onto the clipboard

Currently, the implementation is pretty rudimentary and expects everything to be fine.

## Versioning
Version information can be injected into Windows builds of the executable using [winres](https://github.com/tc-hib/winres).

### Installing winres
Install with this command and then check the website
```
go install github.com/tc-hib/go-winres@latest
```

### Using winres
As the base winres files are already in this repo, simply update ```winres/winres.json``` as required and then run this command before building the executable
```
go-winres make
```

## Building the executable
This can be built as an executable with the following command:
```
go build
```

For Windows, build with special load flags to make the program run without opening a console window, which is useful if you want to use this in the Send To context menu:

```
go build -ldflags -H=windowsgui
```

## Packages
The project uses the following packages:
 * fmt
 * golang.design/x/clipboard
 * io/ioutil
 * os

### Clipboard package - ```golang.design/x/clipboard```
To use the clipboard package, you may need to download it first:
```
go get golang.design/x/clipboard
```