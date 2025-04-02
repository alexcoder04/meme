
# meme

Get memory usage information in Go. This library obtains and parses memory data from `/proc/meminfo`.

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
