package api

import (
	"code.google.com/p/go.net/context"
	"encoding/json"
	"github.com/brain/instance"
	"net/http"
)

func MetricGetList(w http.ResponseWriter, r *http.Request, ctx context.Context) {
	w.Header().Add("content-type", "application/json")

	ds := ctx.Value("ds").(*core.Datastore)
	db := ds.Db
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("instances")).Cursor()

		var data []brain.Instance
		for k, val := c.First(); k != nil; k, val = c.Next() {
			var m brain.Instance
			if err := json.Unmarshal(val, &m); err != nil {
				return err
			}
			data = append(data, m)
		}

		jsonData, err := json.Marshal(data)
		w.Write(jsonData)
		return err
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("get metric list: %v", err)
	}
}
