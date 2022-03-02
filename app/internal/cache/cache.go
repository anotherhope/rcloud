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
	originalStat, _ := os.Stat(sourcePath)
	if originalStat.IsDir() {
		return false
	}

	original, _ := os.OpenFile(sourcePath, os.O_RDONLY, 0700)
	defer original.Close()

	cache, _ := os.OpenFile(cachePath, os.O_CREATE|os.O_RDWR, 0700)
	defer cache.Close()

	sourceChecksum := calculateChecksum(original)
	cacheChecksum := getCacheCheksum(cache)

	return sourceChecksum != cacheChecksum
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

func (c *Cache) Add(sourcePath string) {
	c.Sign(sourcePath, c.MakeCachePath(sourcePath))
}

func (c *Cache) Update(sourcePath string) {
	c.Add(sourcePath)
}

func (c *Cache) Sign(sourcePath string, cachePath string) error {
	sourceInfo, err := os.Stat(sourcePath)
	if err != nil {
		return err
	}

	if sourceInfo.IsDir() {
		return os.MkdirAll(cachePath, 0700)
	}

	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	cache, err := os.OpenFile(cachePath, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer cache.Close()

	cache.Truncate(0)
	cache.Seek(0, 0)
	cache.WriteString(
		calculateChecksum(source),
	)

	return nil
}

func (c *Cache) MakeCachePath(sourcePath string) string {
	if len(sourcePath) <= len(c.Base) {
		return CachePath
	}

	relative := sourcePath[len(c.Base):]
	cachePath := CachePath + "/" + c.Id + relative

	return cachePath
}

func NewCache(Id string, Base string) *Cache {
	return &Cache{
		Id:   Id,
		Base: Base,
	}
}
