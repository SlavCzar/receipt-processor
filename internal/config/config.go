package config

type Config struct {
    Port string
}

func Load() *Config {
    return &Config{
        Port: "8080", // Can load from an env file in the future
    }
}
