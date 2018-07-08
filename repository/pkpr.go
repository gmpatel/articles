package repository

import (
	"context"
	"database/sql"
	"fmt"
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
func (repo *SQLRepository) ensureConnected() error {
	if repo.mutex != nil {
		repo.mutex.Lock()
		defer repo.mutex.Unlock()
	}

	log.Info("Attempting to connect to the db...")
	count := 0

	for {
		count++

		err := repo.db.Ping()
		if err != nil {
			if count >= 5 {
				return fmt.Errorf("Database connection attempt failed. Cause: %s", err.Error())
			}
			log.Warnf("Database connection attempt failed. Cause: %s", err.Error())
			time.Sleep(time.Second)
			continue
		}

		break
	}

	log.Infof("Successfully connected to db after %d attempts", count)
	return nil
}

func (repo *SQLRepository) queryContext(ctx context.Context, cmd string) (*sql.Rows, []string, error) {
	// Querying the database for given query and args
	rows, err := repo.db.QueryContext(ctx, cmd)
	if err != nil {
		log.Errorf("Failed to query '%s': %v", cmd, err)
		return nil, nil, err
	}

	// Check if our query timed out
	if ctx.Err() == context.DeadlineExceeded {
		log.Errorf("Query timeout: %v", ctx.Err())
		return nil, nil, ctx.Err()
	}

	// Checking any errors while fetching rows
	if rows.Err() != nil {
		log.Errorf("Failed to fetch rows: %v", rows.Err())
		return nil, nil, rows.Err()
	}

	// Checking any errors while columns for rows
	cols, colsErr := rows.Columns()
	if colsErr != nil {
		log.Errorf("failed to get columns from the rows: %v", colsErr)
		return nil, nil, colsErr
	}

	// Return query assets
	return rows, cols, nil
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

func getScallerValue(rows *sql.Rows) interface{} {
	// Checking any errors while fetching rows
	if rows.Err() != nil {
		log.Errorf("failed to fetch rows: %v", rows.Err())
		return rows.Err()
	}

	// Checking any errors while columns for rows
	cols, colsErr := rows.Columns()
	if colsErr != nil {
		log.Errorf("failed to get columns from the rows: %v", colsErr)
		return colsErr
	}

	// Moving rows to the first result row
	rows.Next()
	vals := getCellsArray(len(cols))
	scanErr := rows.Scan(vals...)
	if scanErr != nil {
		log.Errorf("failed to scan rows data: %v", scanErr)
		return rows.Err()
	}

	// Return the first cells value
	return getCellValue(vals[0].(*interface{}))
}
