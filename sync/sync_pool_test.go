package sync

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

// `sync.Pool` 是 Go 语言提供的一个对象池，用于存储临时对象，可以有效地提高代码的性能。在使用时，可以通过 `Get()` 方法获取一个对象，使用完成后通过 `Put()` 方法将对象放回池中，以便下次使用。
// 下面是一个使用 `sync.Pool` 的示例，假设需要频繁地创建一个长度为 1024 的 byte 数组，并在使用完成后将其丢弃。使用 `sync.Pool` 可以避免频繁的内存分配和垃圾回收，从而提高性能。
func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}

	// 从对象池中获取一个 byte 数组
	buf := pool.Get().([]byte)
	//输出buf地址
	fmt.Printf("%p\n", buf)

	// 使用 byte 数组
	copy(buf, []byte("Hello, world!"))
	fmt.Println(string(buf))

	// 将 byte 数组放回对象池中
	pool.Put(buf)
	pool.Put(buf)
	pool.Put(buf)
	pool.Put(buf)
	pool.Put(buf)

	// 再次从对象池中获取 byte 数组，此时应该是从对象池中获取的
	for i := 0; i < 50; i++ {
		buf := pool.Get().([]byte)
		fmt.Println(string(buf))
		fmt.Printf("%p\n", buf)
	}
}

// 在上面的示例中，我们首先创建了一个 `sync.Pool` 对象，并通过 `New` 字段指定了一个用于创建新对象的函数，该函数创建一个长度为 1024 的 byte 数组。
// 然后我们通过 `Get()` 方法从对象池中获取一个 byte 数组，使用完成后将其放回对象池中，再次调用 `Get()` 方法时应该会从对象池中获取到上一次放回去的 byte 数组。

// `sync.Pool` 的设计目的是为了减少 GC 压力，因此在使用时需要注意以下几点：
//1. 对象池中的对象不应该有任何外部依赖，以免影响其复用。例如，对象池中的对象应该是无状态的。
//2. 对象池中的对象应该是可以重复使用的。换句话说，对象池中的对象应该能够被多次获取和放回。
//3. 对象池中的对象应该是可变的。如果对象池中的对象是不可变的，则在每次使用之前都需要进行深拷贝，这样会增加 CPU 的负载和内存的占用。
//4. 对象池中的对象应该是相对较小的，否则会增加内存占用。在大多数情况下，对象池中的对象应该是固定大小的。
//5. 对象池应该被设计成并发安全的。多个 Goroutine 可以同时访问对象池，因此需要保证对象池的线程安全性。
//需要注意的是，由于对象池中的对象可以被多次获取和放回，因此在使用完之后需要清理对象中的敏感信息，以免泄露敏感信息。此外，由于对象池会减少垃圾回收的次数，因此在使用完之后需要将对象放回池中，以便下一次复用。

// ----------------------------------------------------------------------------------------------------------------------
// 由于 sync.Pool 是一个并发安全的对象池，多个 goroutine 可以同时从中获取和放回对象，因此我们需要在处理 Request 对象时进行适当的同步，以避免出现竞态条件。
// 一种解决方法是为每个 goroutine 分配一个 Request 对象，这样就不需要进行同步操作。
// 另一种解决方法是使用一个 sync.Mutex 对象对 Request 对象进行保护，这样可以确保同一时间只有一个 goroutine 在处理 Request 对象。下面是使用 Mutex 的示例代码：
func TestSyncPoolConcurrent(t *testing.T) {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

var reqPool = sync.Pool{
	New: func() interface{} {
		return &http.Request{}
	},
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 从对象池中获取 Request 对象
	req := reqPool.Get().(*http.Request)

	// 对 Request 对象加锁，保护它的状态
	reqLock := sync.Mutex{}
	reqLock.Lock()

	// 处理请求
	// ...

	// 请求处理完成后释放锁并将 Request 对象放回对象池
	reqLock.Unlock()
	reqPool.Put(req)
}

//在上面的示例中，我们为每个 Request 对象分配了一个 Mutex 对象来进行同步，以避免出现竞态条件。在处理 Request 对象时，我们首先从对象池中获取一个 Request 对象，然后为它分配一个 Mutex 对象进行同步，处理完请求后再将 Request 对象放回对象池。
//需要注意的是，使用 Mutex 的方法会影响性能，因为在处理 Request 对象时需要进行加锁和解锁操作。如果您的应用程序在高并发情况下需要处理大量的请求，建议使用第一种方法为每个 goroutine 分配一个 Request 对象。
