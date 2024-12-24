package main

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/jamesstocktonj1/component-cdc/keyvalue/gen/wasi/keyvalue/store"
)

func init() {
	store.Exports.Open = Open
	store.Exports.Bucket.Destructor = BucketDestructor
	store.Exports.Bucket.Delete = BucketDelete
	store.Exports.Bucket.Exists = BucketExists
	store.Exports.Bucket.Get = BucketGet
	store.Exports.Bucket.ListKeys = BucketListKeys
	store.Exports.Bucket.Set = BucketSet
}

func Open(identifier string) cm.Result[store.ErrorShape_, store.ExportBucket, store.Error] {
	res := store.Open(identifier)
	if res.IsErr() {
		return cm.Err[cm.Result[store.ErrorShape_, store.ExportBucket, store.Error]](*res.Err())
	} else {
		return cm.OK[cm.Result[store.ErrorShape_, store.ExportBucket, store.Error]](store.ExportBucket(*res.OK()))
	}
}

func BucketDestructor(self cm.Rep) {
	store.Bucket(self).ResourceDrop()
}

func BucketDelete(self cm.Rep, key string) cm.Result[store.Error, struct{}, store.Error] {
	return store.Bucket(self).Delete(key)
}

func BucketExists(self cm.Rep, key string) (result cm.Result[store.ErrorShape_, bool, store.Error]) {
	res := store.Bucket(self).Exists(key)
	if res.IsErr() {
		return cm.Err[cm.Result[store.ErrorShape_, bool, store.Error]](*res.Err())
	} else {
		return cm.OK[cm.Result[store.ErrorShape_, bool, store.Error]](*res.OK())
	}
}

func BucketGet(self cm.Rep, key string) cm.Result[store.OptionListU8Shape_, cm.Option[cm.List[uint8]], store.Error] {
	res := store.Bucket(self).Get(key)
	if res.IsErr() {
		return cm.Err[cm.Result[store.OptionListU8Shape_, cm.Option[cm.List[uint8]], store.Error]](*res.Err())
	} else {
		return cm.OK[cm.Result[store.OptionListU8Shape_, cm.Option[cm.List[uint8]], store.Error]](*res.OK())
	}
}

func BucketListKeys(self cm.Rep, cursor cm.Option[uint64]) (result cm.Result[store.KeyResponseShape_, store.KeyResponse, store.Error]) {
	res := store.Bucket(self).ListKeys(cursor)
	if res.IsErr() {
		return cm.Err[cm.Result[store.KeyResponseShape_, store.KeyResponse, store.Error]](*res.Err())
	} else {
		return cm.OK[cm.Result[store.KeyResponseShape_, store.KeyResponse, store.Error]](*res.OK())
	}
}

func BucketSet(self cm.Rep, key string, value cm.List[uint8]) (result cm.Result[store.Error, struct{}, store.Error]) {
	return store.Bucket(self).Set(key, value)
}

//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world capture --out gen ./wit
func main() {}
