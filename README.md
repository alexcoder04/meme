
# meme

[![License](https://img.shields.io/github/license/alexcoder04/meme)](https://github.com/alexcoder04/meme/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/alexcoder04/meme)](https://github.com/alexcoder04/meme/blob/main/go.mod)
[![Lines](https://img.shields.io/tokei/lines/github/alexcoder04/meme?label=lines)](https://github.com/alexcoder04/meme/pulse)
[![Release](https://img.shields.io/github/v/release/alexcoder04/meme?display_name=tag&sort=semver)](https://github.com/alexcoder04/meme/releases/latest)
[![Stars](https://img.shields.io/github/stars/alexcoder04/meme)](https://github.com/alexcoder04/meme/stargazers)
[![Contributors](https://img.shields.io/github/contributors-anon/alexcoder04/meme)](https://github.com/alexcoder04/meme/graphs/contributors)

Get memory usage information in Go.

Currently Linux only.

## Install and Use

In your project directory, type

```sh
go get github.com/alexcoder04/meme
```

And then, in your code

```go
memData, err := meme.GetMemInfo()
```

The result comes in following type:

```go
type MemData struct {
	Buffers      float64
	Cached       float64
	MemAvailable float64
	MemFree      float64
	MemTotal     float64
	Shared       float64
	Used         float64

	SwapFree  float64
	SwapTotal float64
	SwapUsed  float64
}
```
