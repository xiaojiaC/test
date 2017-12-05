package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const unitSize = 1400 // read/write chunk size. ~MTU size.

type Rate struct {
	KBps    int // or 0, to not rate-limit bandwidth
	Latency time.Duration
}

// byteTime returns the time required for n bytes.
func (r Rate) byteTime(n int) time.Duration {
	if r.KBps == 0 {
		return 0
	}
	return time.Duration(float64(n)/1024/float64(r.KBps)) * time.Second
}

type Listener struct {
	net.Listener
	Down Rate // server Writes to Client
	Up   Rate // server Reads from client
}

func (ln *Listener) Accept() (net.Conn, error) {
	c, err := ln.Listener.Accept()
	time.Sleep(ln.Up.Latency)
	if err != nil {
		return nil, err
	}
	tc := &conn{Conn: c, Down: ln.Down, Up: ln.Up}
	tc.start()
	return tc, nil
}

type nErr struct {
	n   int
	err error
}

type writeReq struct {
	writeAt time.Time
	p       []byte
	resc    chan nErr
}

type conn struct {
	net.Conn
	Down Rate // for reads
	Up   Rate // for writes

	wchan     chan writeReq
	closeOnce sync.Once
	closeErr  error
}

func (c *conn) start() {
	c.wchan = make(chan writeReq, 1024)
	go c.writeLoop()
}

func (c *conn) writeLoop() {
	for req := range c.wchan {
		time.Sleep(req.writeAt.Sub(time.Now()))
		var res nErr
		for len(req.p) > 0 && res.err == nil {
			writep := req.p
			if len(writep) > unitSize {
				writep = writep[:unitSize]
			}
			n, err := c.Conn.Write(writep)
			time.Sleep(c.Up.byteTime(len(writep)))
			res.n += n
			res.err = err
			req.p = req.p[n:]
		}
		req.resc <- res
	}
}

func (c *conn) Close() error {
	c.closeOnce.Do(func() {
		err := c.Conn.Close()
		close(c.wchan)
		c.closeErr = err
	})
	return c.closeErr
}

var (
	mu      sync.Mutex
	start   = time.Now()
	last    string
	lastDir string
)

func transfer(dir string, n int) {
	if !*verbose {
		return
	}
	mu.Lock()
	t := fmt.Sprintf("%.3fs", time.Since(start).Seconds())
	sep := "+"
	if t != last {
		fmt.Printf("\n%s\n\t%s", t, dir)
		last = t
		lastDir = dir
		sep = ""
	} else if dir != lastDir {
		fmt.Printf("\n\t%s", dir)
		lastDir = dir
		sep = ""
	}
	fmt.Printf("%s%d", sep, n)
	mu.Unlock()
}

func (c *conn) Write(p []byte) (n int, err error) {
	defer func() {
		if e := recover(); e != nil {
			n = 0
			err = fmt.Errorf("%v", err)
			return
		}
	}()
	resc := make(chan nErr, 1)
	c.wchan <- writeReq{time.Now().Add(c.Up.Latency), p, resc}
	res := <-resc

	transfer("->", res.n)
	return res.n, res.err
}

func (c *conn) Read(p []byte) (n int, err error) {
	const max = 1024
	if len(p) > max {
		p = p[:max]
	}
	n, err = c.Conn.Read(p)
	time.Sleep(c.Down.byteTime(n))

	transfer("<-", n)

	return
}
