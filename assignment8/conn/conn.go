package conn

import (
	"bytes"
	"errors"
	"net"
	"strings"
	"time"
)

type connImpl struct {
	connectionClosed bool
	buffer           *bytes.Buffer
	pos              int
}

func NewConnection() net.Conn {
	return &connImpl{
		connectionClosed: false,
		buffer:           bytes.NewBufferString(strings.Repeat("Money\nVijay\nParitosh\nShivam\n", 1)),
		pos:              0,
	}
}

func (c *connImpl) Read(b []byte) (int, error) {
	if c.connectionClosed {
		return 0, errors.New("Connection is Closed")
	}
	return c.buffer.Read(b)
}

func (c *connImpl) Write(b []byte) (n int, err error) {
	return 0, errors.New("Operation write not permitted on this implementation") // client never writes to server
}

func (c *connImpl) Close() error {
	if c.connectionClosed {
		return errors.New("Connection already closed")
	}
	c.connectionClosed = true
	return nil
}

func (c *connImpl) LocalAddr() net.Addr {
	return nil
}

func (c *connImpl) RemoteAddr() net.Addr {
	return nil
}

func (c *connImpl) SetDeadline(t time.Time) error {
	return nil
}

func (c *connImpl) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *connImpl) SetWriteDeadline(t time.Time) error {
	return nil
}
