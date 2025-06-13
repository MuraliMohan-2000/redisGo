package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tidwall/resp"
)

func TestOfficialRedisCLient(t *testing.T) {

	listenAddr := ":8080"

	server := NewServer(Config{
		ListenAddress: listenAddr,
	})

	go func() {
		log.Fatal(server.Start())
	}()

	time.Sleep(time.Millisecond * 400)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:8080",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	testCases := map[string]string{
		"foo":  "bar",
		"a":    "bbb",
		"key1": "val1",
		"key2": "val2",
	}

	for key, val := range testCases {

		if err := rdb.Set(context.Background(), key, val, 0).Err(); err != nil {
			t.Fatal(err)
		}

		newval, err := rdb.Get(context.Background(), key).Result()
		if err != nil {
			t.Fatal(err)
		}
		if newval != val {
			t.Fatalf("expected %s but got %s", val, newval)
		}
	}
}
func TestFooBar(t *testing.T) {
	buf := &bytes.Buffer{}
	rw := resp.NewWriter(buf)
	rw.WriteString("Ok")
	fmt.Println(buf.String())

	in := map[string]string{
		"server":  "redis",
		"version": "7.0.12",
	}
	out := respWriteMap(in)
	fmt.Println(string(out))
}
func TestServerWithMultiClients(t *testing.T) {

	server := NewServer(Config{})

	go func() {
		log.Fatal(server.Start())
	}()

	time.Sleep(time.Second)

	nClients := 10
	wg := sync.WaitGroup{}
	wg.Add(nClients)
	for i := 0; i < nClients; i++ {
		go func(it int) {

			rdb := redis.NewClient(&redis.Options{
				Addr:     "localhost:8080",
				Password: "", // no password set
				DB:       0,  // use default DB
			})

			defer rdb.Close()

			key := fmt.Sprintf("client_foo_%d", i)
			val := fmt.Sprintf("client_bar_%d", i)
			if err := rdb.Set(context.Background(), key, val, 0).Err(); err != nil {
				t.Fatal(err)
			}

			newval, err := rdb.Get(context.Background(), key).Result()
			if err != nil {
				t.Fatal(err)
			}
			if newval != val {
				t.Fatalf("expected %s but got %s", val, newval)
			}
			wg.Done()

		}(i)
	}
	wg.Wait()

	time.Sleep(time.Second)

	if len(server.peers) != 0 {
		t.Fatalf("expected 0 peers but got %d", len(server.peers))
	}

}
