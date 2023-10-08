package datasource

import (
	"fmt"

	"github.com/k0825/go-gin-ent-sample/config"
	"github.com/k0825/go-gin-ent-sample/ent"
	_ "github.com/lib/pq"
)

func NewConnection(conf *config.Config) (*ent.Client, error) {
	client, err := ent.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			conf.RDB.Host,
			conf.RDB.Port,
			conf.RDB.UserName,
			conf.RDB.Database,
			conf.RDB.Password))

	if err != nil {
		return nil, err
	}

	return client, nil
}
