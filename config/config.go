package config
import (
  "time"
)
type ServerConfig struct {
    Addr         string
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
}
func ConfigMy()ServerConfig{
  config := ServerConfig{
        Addr:         ":8080",
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }
    return config
}
