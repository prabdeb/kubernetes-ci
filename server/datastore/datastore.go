package datastore

import (
	"time"

	"github.com/prabdeb/kubernetes-ci/k8s"

	log "github.com/Sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
)

// DataStore type
type DataStore struct {
	logger  *bolt.DB
	configs *k8s.Kubernetes
}

// New Client
func New() *DataStore {
	var dataSource DataStore
	err := dataSource.Init()
	if err != nil {
		log.Fatal("error initializing datastore - ", err.Error())
	}
	dataSource.configs = k8s.New()
	return &dataSource
}

// Init func
func (d *DataStore) Init() error {
	db, err := bolt.Open("kubernetes.ci.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	defer db.Close()
	d.logger = db
	return nil
}
