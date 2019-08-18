# GiftBog ![Language](https://img.shields.io/badge/language-Go-blue?style=plastic)
A simple command line interface for the r/FashionReps repfams to scrape Superbuy's giftbags.

Installation
------------
``` 
$ go get -u -v github.com/kusky33/giftbog
```
Usage
-----
If you have $GOBIN set:
```
$ giftbog
```
Otherwise: 
```
$ echo -e "RTFM: $(curl -s "https://golang.org/doc/code.html#GOPATH"|grep "export P"|head -c-5|tail -c-38)"
```
