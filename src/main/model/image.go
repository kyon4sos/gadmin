package model

type Image struct {
	name string
	url string
}

func CachePrefix() string {
	return cachePrefix
}