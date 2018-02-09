# 14.2.8 Future/Promise

複数のタスクがFutureから値を取得できるようキャッシュ機能を持ったFuture/Promiseパターン


# Result

```bash
$ go run main.go
func NewStringFuture() (*StringFuture, chan string) {
func (f *StringFuture) Get() string {
func (f *StringFuture) Close() {
func readFile(path string) *StringFuture {
func printFunc(futureSource *StringFuture) chan []string {
func countLines(futureSource *StringFuture) chan int {
func main() {
85
```
