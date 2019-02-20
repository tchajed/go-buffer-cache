package cache

import "github.com/tchajed/go-buffer-cache/disk"

type Cache struct {
	m map[int]disk.Block
}

func New() Cache {
	return Cache{m: make(map[int]disk.Block)}
}

func Recover() Cache {
	return New()
}

func (c Cache) Read(addr int) disk.Block {
	b, ok := c.m[addr]
	if ok {
		return b
	}
	// miss, go to disk
	b = disk.Read(addr)
	c.m[addr] = b
	return b
}

func (c Cache) Write(addr int, b disk.Block) {
	c.m[addr] = b
	disk.Write(addr, b)
}
