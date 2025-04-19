package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	purgeCount int
	mu         sync.Mutex
)

func purgeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	purgeCount++
	mu.Unlock()

	log.Println("🔁 Purge request received")
	fmt.Fprintf(w, "✅ Cache purged! Total purges: %d\n", purgeCount)
}

func main() {
	http.HandleFunc("/purge", purgeHandler)
	log.Println("🚀 Go Cache Controller running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
