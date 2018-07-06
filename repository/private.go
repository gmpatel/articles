package repository

import (
	"time"

	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func init() {
	SetLogger(logrus.StandardLogger())
}

// ensureConnected will test the connection and if not connected, will continually retry on a configured interval
func (repo *SQLRepository) ensureConnected() {
	if repo.mutex != nil {
		repo.mutex.Lock()
		defer repo.mutex.Unlock()
	}

	log.Info("Attempting to connect to db")
	count := 0

	for {
		count++
		err := repo.db.Ping()
		if err != nil {
			log.Warnf("Database connection attempt failed. Cause: %s", err.Error())
			time.Sleep(time.Second)
			continue
		}
		break
	}

	log.Infof("Successfully connected to db after %d attempts", count)
}

func getCellsArray(len int) []interface{} {
	vals := make([]interface{}, len)
	for i := 0; i < len; i++ {
		vals[i] = new(interface{})
	}
	return vals
}

func getCellValue(p *interface{}) interface{} {
	var ret interface{}
	switch v := (*p).(type) {
	default:
		ret = v
	}
	return ret
}
