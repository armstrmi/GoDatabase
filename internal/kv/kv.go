package kv

type KV interface {
	Get(key []byte) (val []byte, ok bool)
}
