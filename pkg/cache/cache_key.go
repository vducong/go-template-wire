package cache

import "fmt"

func GetCacheKey(template CacheKeyTemplate, params ...string) string {
	return fmt.Sprintf(string(template), params)
}

type CacheKeyTemplate string

const (
	KeyReusableCodeAllActive CacheKeyTemplate = "REUSABLE_CODE_ALL_ACTIVE"
)
