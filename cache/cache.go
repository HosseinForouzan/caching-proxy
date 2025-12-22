package cache

type Cache interface {
	Get(url string) (*CachedPage, error)
	Set(url string, cachedPage *CachedPage) (error)
}