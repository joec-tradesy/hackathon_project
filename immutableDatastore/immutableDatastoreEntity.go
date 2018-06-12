package immutableDatastore

import (
	"cloud.google.com/go/datastore"
	"fmt"
	"golang.org/x/net/context"
	"time"
)

// ImmutableEntity is the base object that handles immutability
type ImmutableEntity struct {
	VersionID          int64
	Client             datastore.Client `datastore:"-"`
	KindName           string           `datastore:"-"`
	Key                datastore.Key    `datastore:"-"`
	firstEntityInChain datastore.Key
	consistentID       int64
	isTip              bool
	timestamp          time.Time
}

var ctx = context.Background()

func (e ImmutableEntity) Get(name string) error {
	e.Key = *datastore.NameKey(e.KindName, name, nil)
	err := e.Client.Get(ctx, &e.Key, &e)
	if err != nil {
		return err
	}
	return nil
}

func (e ImmutableEntity) Create() error {

	e.Key = *datastore.NameKey(e.KindName, "testkey1", nil)
	fmt.Println("2", e)
	_, err := e.Client.Put(ctx, &e.Key, &e)
	if err != nil {
		return err
	}

	return nil
}
