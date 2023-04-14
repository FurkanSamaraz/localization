package config

//OptionsSession manages the session generation and caching options
type OptionsSession struct {
	//MemCacheExp dictates how long the in-memory keys are valid in minutes
	MemCacheExp int `json:"memcache-exp"`

	//MemCacheClean dictates the interval in which the cleaning thread is ran in minutes
	MemCacheClean int `json:"memcache-clean"`

	//KeyLength is the byte length of the key generation, it get's base64'd so it is not output length
	KeyLength int `json:"key-length"`

	//RedisExp dictates how long the Redis key life should be in minutes
	RedisExp int `json:"redis-exp"`
}
