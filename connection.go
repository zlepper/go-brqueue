package brqueue

import (
	"net"
	"sync"
)

type connection struct {
	conn net.Conn
	readLock sync.Mutex
}

func (c *connection) close() error {
	return c.conn.Close()
}