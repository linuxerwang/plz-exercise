package main

import (
	"time"

	etcd "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := etcd.New(etcd.Config{
		Endpoints:   []string{"localhost:2300"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	_ = cli
}
