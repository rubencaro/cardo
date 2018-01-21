package db

import (
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/rubencaro/cardo/lib/cnf"
)

// GetCollection returns a Collection struct for given name
func GetCollection(name string, c *cnf.Cnf) (driver.Collection, error) {
	database, err := getDatabase("devel", c)
	if err != nil {
		return nil, err
	}

	exists, err := database.CollectionExists(nil, name)
	if err != nil {
		return nil, err
	}

	var coll driver.Collection
	if !exists {
		coll, err = database.CreateCollection(nil, name, nil)
	} else {
		coll, err = database.Collection(nil, name)
	}
	if err != nil {
		return nil, err
	}

	return coll, nil
}

func connect(c *cnf.Cnf) (driver.Client, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		return nil, err
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(c.GetString("arango.user"), c.GetString("arango.pass")),
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getDatabase(dbname string, c *cnf.Cnf) (driver.Database, error) {
	client, err := connect(c)
	if err != nil {
		return nil, err
	}

	exists, err := client.DatabaseExists(nil, dbname)
	if err != nil {
		return nil, err
	}

	var database driver.Database
	if !exists {
		database, err = client.CreateDatabase(nil, dbname, nil)
	} else {
		database, err = client.Database(nil, dbname)
	}
	if err != nil {
		return nil, err
	}

	return database, nil
}
