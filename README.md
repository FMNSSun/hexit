# hexit

Hexit provides utility functions to convert `uint*` to hex. It's faster than using `fmt.Sprintf`, `strconv.AppendUint`
or `strconv.FormatUint`.

## Examples

```go
buf := make([]byte, 8)
hexit.HexUint32To(num, buf)
```

```go
buf := hexit.HexUint32(num)
```

```go
str := hexit.HexUint16(num)
```