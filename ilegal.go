
type Buffer struct {
	count int
	mu    sync.Mutex
}
var buffer Buffer
var cond = sync.NewCond(&buffer.mu)
var size = 3
func producer(buffer *Buffer) {
	for {

		for buffer.count == size {
			cond.Wait()
		}
		buffer.mu.Lock()
		buffer.count++
		fmt.Print("(")

		cond.Broadcast()
		buffer.mu.Unlock()
	}
}
func consumer(buffer *Buffer) {
	for {
		for buffer.count == 0 {
			cond.Wait()
		}
		buffer.mu.Lock()
		buffer.count--
		fmt.Print(")")

		cond.Broadcast()
		buffer.mu.Unlock()
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(&buffer)
	go producer(&buffer)
	go consumer(&buffer)
	go consumer(&buffer)
	wg.Wait()
}
