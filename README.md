# gomap
A very simple way to wrap sync.map, avoiding type assertions.


---
## Example:


```go
func main() {
    mp := core.Make[int, int]()
    wg := sync.WaitGroup{}
    for i := 0; i < 10; i++ {
        i := i
        wg.Add(1)
        go func() {
            defer wg.Done()
            mp.Put(i, i)
        }()
    }
    wg.Wait()
    mp.Range(func(key, val int) bool {
        fmt.Println(key, val)
        return true
    })
}
// 9 9
// 1 1
// 2 2
// 4 4
// 5 5
// 7 7
// 0 0
// 3 3
// 6 6
// 8 8
```
