# d
[![Build Status](https://travis-ci.org/soider/d.svg?branch=develop)](https://travis-ci.org/soider/d)
[![GoDoc](https://godoc.org/github.com/soider/d?status.svg)](https://godoc.org/github.com/soider/d)
[![Go Report Card](https://goreportcard.com/badge/github.com/soider/d)](https://goreportcard.com/report/github.com/soider/d)

d is a better way to do print statement debugging.

Type `d.D` instead of `fmt.Printf` and your variables will be printed like this:

![d output examples](https://i.imgur.com/OFmm7pb.png)

## Why is this better than `fmt.Printf`?

* Faster to type
* Pretty-printed vars and expressions
* Easier to see inside structs
* Pretty colors!

d is fork of https://github.com/y0ssar1an/q with some differences:
* Added glide support
* Do not include vendor into repository
* Changed default behaviour - for me is more comfortable to log on stdout or have possibility to change file
* Settings for removing colors


## Basic Usage

```go
import "github.com/soider/d"
...
d.D(a, b, c)

// Alternatively, use the . import and you can omit the package name.
// q only exports the Q function.
import . "github.com/soider/d"
...
Q(a, b, c)
```


For best results, dedicate a terminal to tailing `$TMPDIR/q` while you work.

## Install

```sh
go get -u github.com/soider/d
```

## Editor Integration

#### Sublime Text
```
cp $GOPATH/src/github.com/soider/d/dd.sublime-snippet Packages/User/dd.sublime-snippet
```

#### Atom
Navigate to your `snippets.cson` file by either opening `~/.atom/snippets.cson`
directly or by selecting the `Atom > Open Your Snippets` menu. You can then add
this code snippet to the bottom and save the file:
```
'.source.go':
  'd log':
    'prefix': 'dd'
    'body': 'd.D($1)'
```

#### VS Code
In the VS Code menu go to `Preferences` and choose `User Snippets`. When the
language dropdown menu appears select `GO`. Add the following snippet to the
array of snippets.
```
"d.D": {
	"prefix": "dd",
	"body": [
		"d.D($1)"
	],
	"description": "Quick and dirty debugging output for tired Go programmers"
}
```

#### vim/Emacs
TBD Send me a PR, please :)

## Haven't I seen this somewhere before?

Python programmers will recognize this as a Golang port of the
[`q` module by zestyping](https://github.com/zestyping/q).

Ping does a great job of explaining `q` in his awesome lightning talk from
PyCon 2013. Watch it! It's funny :)

[![ping's PyCon 2013 lightning talk](https://i.imgur.com/7KmWvtG.jpg)](https://youtu.be/OL3De8BAhME?t=25m14s)

## FAQ

### Why `d.D`?
It's quick to type and unlikely to cause naming collisions.

### Is `d.D()` safe for concurrent use?
Yes.
