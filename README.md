# clip-file-go
Go implementation of clip-file, to read a text file and place its contents onto the clipboard

Currently, the implementation is pretty rudimentary and expects everything to be fine.

For Windows, build with this to make the program run without opening a console window, which is useful if you want to use this in the Send To context menu:

```
go build -ldflags -H=windowsgui hello.go
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