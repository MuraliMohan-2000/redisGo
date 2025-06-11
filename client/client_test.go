package client

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestNewClient1(t *testing.T) {
	client, err := New("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Set(context.TODO(), "foo", "1"); err != nil {
		log.Fatal(err)
	}

	val, err := client.Get(context.TODO(), "foo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)

}

func TestNewClient(t *testing.T) {
	client, err := New("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {

		fmt.Println("SET =>", fmt.Sprintf("foo_%d", i))

		if err := client.Set(context.TODO(), fmt.Sprintf("foo_%d", i), fmt.Sprintf("bar_%d", i)); err != nil {
			log.Fatal(err)
		}

		val, err := client.Get(context.TODO(), fmt.Sprintf("foo_%d", i))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("GET =>", val)

	}
}
