package repository

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/denisenkom/go-mssqldb" // Blank import is important to register the sql drivers "mssql" / "sqlserver"
	"github.com/sirupsen/logrus"
)

// SetLogger sets the logger for this package
func SetLogger(logger *logrus.Logger) {
	log = logger
}

// SQLRepository object to provide repository methods talking to the db
type SQLRepository struct {
	db       *sql.DB
	doneChan chan bool
	mutex    *sync.Mutex
	timeout  time.Duration
}

// NewRepository returns the repository object constructed with the parameter given
func NewRepository(connString string, conns int, timeoutMillis int) (*SQLRepository, error) {
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(conns)
	db.SetMaxIdleConns(conns)

	return &SQLRepository{
		db: db, doneChan: make(chan bool),
		mutex:   new(sync.Mutex),
		timeout: time.Millisecond * time.Duration(timeoutMillis),
	}, nil
}

// Start will start the repository
func (repo *SQLRepository) Start() {
	repo.ensureConnected()
}

// Stop closes any resources used by the repository
func (repo *SQLRepository) Stop() {
	repo.doneChan <- true
	err := repo.db.Close()
	if err != nil {
		log.Debugf("Failed to close connection: %v", err)
	}
}

/* FetchChanges checks a temporal table for any new row changes
func (repo *SQLRepository) FetchChanges(config *TrackingConfig) ([]interface{}, error) {
	resultSet := make([]interface{}, 0)

	ctx, cancel := context.WithTimeout(context.Background(), repo.queryTimeout)
	defer cancel()

	query := getQueryFetchChanges(config)

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		log.Errorf("failed to query from table '%s': %v", config.tableName, err)
		metricsClient.Increment("error.fetch_changes_failed")
		return nil, err
	}

	// check if our query timed out
	if ctx.Err() == context.DeadlineExceeded {
		metricsClient.Increment("error.fetch_changes_timeout")
		return nil, fmt.Errorf("query timeout")
	}

	// Checking any errors while fetching rows
	if rows.Err() != nil {
		log.Errorf("failed to fetch rows from ProcessTracking table: %v", rows.Err())
		return nil, rows.Err()
	}

	// Iterating over the result rows
	for rows.Next() {
		scanData, scanErr := scanFetchChanges(config.tableName, rows)

		if scanErr != nil {
			log.Errorf("failed to scan row cells into the object interface: %v", scanErr)
			return nil, scanErr
		}

		resultSet = append(resultSet, scanData)
	}

	return resultSet, err
}
*/

/* GetLastProcessed returns the last id processed for the given tableName
func (repo *SQLRepository) GetLastProcessed(config *TrackingConfig) error {
	// Setting up context
	ctx, cancel := context.WithTimeout(context.Background(), repo.queryTimeout)
	defer cancel()

	// Building and executing query
	query := "select top 1 Position from ProcessTracking where TableName = @tableName"
	rows, err := repo.db.QueryContext(ctx, query, sql.Named("tableName", config.tableName))

	if err != nil {
		log.Errorf("failed to query from ProcessTracking table for the column '%s': %v", config.tableName, err)
		return err
	}

	// Checking any errors while fetching rows
	if rows.Err() != nil {
		log.Errorf("failed to fetch rows from ProcessTracking table: %v", rows.Err())
		return rows.Err()
	}

	// Checking any errors while columns for rows
	cols, colsErr := rows.Columns()
	if colsErr != nil {
		log.Errorf("failed to get columns from ProcessTracking table rows: %v", rows.Err())
		return colsErr
	}

	// Moving rows to the first result row
	rows.Next()
	vals := getCellsArray(len(cols))
	scanErr := rows.Scan(vals...)
	if scanErr != nil {
		log.Errorf("failed to get rows from ProcessTracking table: %v", rows.Err())
		return rows.Err()
	}

	// Iterating through the cells of result row
	for i := 0; i < len(vals); i++ {
		if strings.EqualFold(cols[i], "Position") {
			config.trackingPosition = getCellValue(vals[i].(*interface{})).(string)
			break
		}
	}

	// Return on success
	return err
}
*/

/* SetLastProcessed updates the given id processed for the given tableName
func (repo *SQLRepository) SetLastProcessed(config *TrackingConfig, newPosition string) error {
	ctx, cancel := context.WithTimeout(context.Background(), repo.queryTimeout)
	defer cancel()

	cmd := "update ProcessTracking set Position = @position where TableName = @tableName"
	_, err := repo.db.ExecContext(ctx, cmd, sql.Named("position", newPosition), sql.Named("tableName", config.tableName))

	config.trackingPosition = newPosition

	return err
}
*/
