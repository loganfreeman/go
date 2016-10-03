import (
  "time"
  "fmt"

  "github.com/diegobernardes/ttlcache"
)

func main () {
  newItemCallback := func(key string, value interface{}) {
        fmt.Printf("New key(%s) added\n", key)
  }
  checkExpirationCallback := func(key string, value interface{}) bool {
        if key == "key1" {
            // if the key equals "key1", the value
            // will not be allowed to expire
            return false
        }
        // all other values are allowed to expire
        return true
    }
  expirationCallback := func(key string, value interface{}) {
        fmt.Printf("This key(%s) has expired\n", key)
    }

  cache := ttlcache.NewCache()
  cache.SetTTL(time.Duration(10 * time.Second))
  cache.SetExpirationCallback(expirationCallback)

  cache.Set("key", "value")
  cache.SetWithTTL("keyWithTTL", "value", 10 * time.Second)

  value, exists := cache.Get("key")
  count := cache.Count()
  result := cache.Remove("key")
}
