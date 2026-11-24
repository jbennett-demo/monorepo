package main

// Libraries this service uses beyond the HTTP skeleton in main.go (database, cache, auth,
// metrics, CLI, RPC). Blank imports keep them in go.mod so `go mod tidy` doesn't prune them.
import (
	_ "github.com/golang-jwt/jwt/v5"
	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/redis/go-redis/v9"
	_ "github.com/spf13/cobra"
	_ "google.golang.org/grpc"
	_ "google.golang.org/protobuf/proto"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
)
