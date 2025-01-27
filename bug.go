```go
func main() {
    var m map[string]int
    m["foo"] = 1 // This will panic if m is nil
    fmt.Println(m["foo"])
}
```