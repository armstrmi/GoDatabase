package kv

type iterator interface {
	HasNext() bool
	Next() (key []byte, val []byte)
}
