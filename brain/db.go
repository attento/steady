package brain

import (
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

type Datastore struct {
	Db *bolt.DB
}

func (ds *Datastore) Open(path string) error {
	db, err := bolt.Open(path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	ds.Db = db
	return nil
}

func (ds *Datastore) Close() error {
	return ds.Db.Close()
}

func (ds *Datastore) GetEnableInstance() error {
	ds.Db.View(func(tx *bolt.Tx) error {
		lb := tx.Bucket([]byte("lb"))
		c := lb.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("walk all: %q %q\n", k, v)
		}
		return nil
	})
	return nil
}

func (ds *Datastore) Init(nodes []string) error {
	err := ds.Db.Update(func(tx *bolt.Tx) error {
		lb, err := tx.CreateBucketIfNotExists([]byte("lb"))
		if err != nil {
			fmt.Printf("%s \n", err)
		}
		for _, ip := range nodes {
			err = lb.Put(nil, []byte(ip))
			if err != nil {
				fmt.Printf("%s \n", err)
			}
		}
		return err
	})
	return err
}
