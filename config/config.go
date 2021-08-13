package config

type Config struct {
	Filename string
	Limit    int
	Shuffle  bool
}

func NewConfig(filename string, limit int, shuffle bool) *Config {
	return &Config{Filename: filename, Limit: limit, Shuffle: shuffle}
}
