package client

import (
	"bytes"
	"context"
	"net"

	"github.com/tidwall/resp"
)

type client struct {
	addr string
	conn net.Conn
}

func New(addr string) (*client, error) {

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &client{
		addr: addr,
		conn: conn,
	}, nil
}

func (c *client) Set(ctx context.Context, key string, val string) error {

	buf := &bytes.Buffer{}
	wr := resp.NewWriter(buf)
	wr.WriteArray([]resp.Value{
		resp.StringValue("SET"),
		resp.StringValue(key),
		resp.StringValue(val),
	})

	_, err := c.conn.Write(buf.Bytes())
	return err
}

func (c *client) Get(ctx context.Context, key string) (string, error) {

	buf := &bytes.Buffer{}
	wr := resp.NewWriter(buf)
	wr.WriteArray([]resp.Value{
		resp.StringValue("GET"),
		resp.StringValue(key),
	})

	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		return "", err
	}
	b := make([]byte, 1024)
	n, err := c.conn.Read(b)
	return string(b[:n]), err
}
