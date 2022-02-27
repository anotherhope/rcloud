package cache

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/anotherhope/rcloud/app/internal/system"
)

type Cache struct {
	Base string
	Id   string
}

// CachePath is the path of cache folder
var CachePath string

func init() {
	CachePath = system.User.HomeDir + "/.config/rcloud/cache"
}

func calculateChecksum(original io.Reader) string {
	hash := sha256.New()
	io.Copy(hash, original)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func getCacheCheksum(cache io.Reader) string {
	buffer := new(strings.Builder)
	io.Copy(buffer, cache)
	return strings.TrimSpace(buffer.String())
}

func (c *Cache) DetectChange(sourcePath string) bool {
	cachePath := c.MakeCachePath(sourcePath)
	original, _ := os.OpenFile(sourcePath, os.O_RDONLY, 0700)
	originalStat, _ := original.Stat()
	defer original.Close()

	if originalStat.IsDir() {
		if _, err := os.Stat(cachePath); os.IsNotExist(err) {
			os.MkdirAll(cachePath, 0700)
		}

		os.Chtimes(
			cachePath,
			originalStat.ModTime().Local(),
			originalStat.ModTime().Local(),
		)
		return true
	}

	cache, _ := os.OpenFile(cachePath, os.O_CREATE|os.O_RDWR, 0700)
	defer cache.Close()

	sourceChecksum := calculateChecksum(original)
	cacheChecksum := getCacheCheksum(cache)

	if sourceChecksum != cacheChecksum {

		cache.Truncate(0)
		cache.Seek(0, 0)
		cache.WriteString(sourceChecksum)

		os.Chtimes(
			cache.Name(),
			originalStat.ModTime().Local(),
			originalStat.ModTime().Local(),
		)

		return true
	}

	return false
}

func (c *Cache) Remove(sourcePaths ...string) {
	if len(sourcePaths) > 0 {
		for _, sourcePath := range sourcePaths {
			cachePath := c.MakeCachePath(sourcePath)
			os.RemoveAll(cachePath)
		}
	} else {
		cachePath := c.MakeCachePath(c.Base)
		os.RemoveAll(cachePath)
	}
}

func (c *Cache) MakeCachePath(sourcePath string) string {
	relative := sourcePath[len(c.Base):]
	cachePath := CachePath + "/" + c.Id + relative

	return cachePath
}

func Make(Id string, Base string) *Cache {
	return &Cache{
		Id:   Id,
		Base: Base,
	}
}
