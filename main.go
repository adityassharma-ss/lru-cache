package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type CacheItem struct {
	Value      interface{} `json:"value"`
	Expiration int64       `json:"expiration"`
}

type LRUCache struct {
	Capacity int
	Items    map[string]CacheItem
	Mutex    sync.RWMutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Capacity: capacity,
		Items:    make(map[string]CacheItem),
	}
}

func (cache *LRUCache) Set(key string, value interface{}, duration time.Duration) {
	cache.Mutex.Lock()
	defer cache.Mutex.Unlock()

	expiration := time.Now().Add(duration).UnixNano()
	cache.Items[key] = CacheItem{
		Value:      value,
		Expiration: expiration,
	}
}

func (cache *LRUCache) Get(key string) (interface{}, bool) {
	cache.Mutex.RLock()
	defer cache.Mutex.RUnlock()

	item, found := cache.Items[key]
	if !found || time.Now().UnixNano() > item.Expiration {
		return nil, false
	}

	return item.Value, true
}

func (cache *LRUCache) Delete(key string) {
	cache.Mutex.Lock()
	defer cache.Mutex.Unlock()

	delete(cache.Items, key)
}

var cache *LRUCache

type SetRequest struct {
	Key      string      `json:"key"`
	Value    interface{} `json:"value"`
	Duration int64       `json:"duration"`
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	var req SetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	duration := time.Duration(req.Duration)

	cache.Set(req.Key, req.Value, duration)
	w.WriteHeader(http.StatusOK)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing key parameter", http.StatusBadRequest)
		return
	}

	value, found := cache.Get(key)
	if !found {
		http.Error(w, "Key not found or expired", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(value)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing key parameter", http.StatusBadRequest)
		return
	}

	cache.Delete(key)
	w.WriteHeader(http.StatusOK)
}

func main() {
	cache = NewLRUCache(10)

	r := mux.NewRouter()

	// CORS middleware
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	r.HandleFunc("/set", setHandler).Methods("POST")
	r.HandleFunc("/get", getHandler).Methods("GET")
	r.HandleFunc("/delete", deleteHandler).Methods("DELETE")

	http.Handle("/", handlers.CORS(originsOk, headersOk, methodsOk)(r))

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
