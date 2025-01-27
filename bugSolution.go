```go
func main() {
    var m map[string]int
    if m == nil {
        m = make(map[string]int)
    }
    m["foo"] = 1 
    fmt.Println(m["foo"])
}
```