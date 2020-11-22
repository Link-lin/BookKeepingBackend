package tool

import (
	"github.com/Link-lin/BookKeepingBackend/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Orm struct {
	*xorm.Engine
}

var DbEngine *Orm

func OrmEngine(cfg *Config) (*Orm, error) {

	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine(database.Driver, conn)

	if err != nil {
		return nil, err
	}

	engine.ShowSQL(database.ShowSQL)

	err = engine.Sync2(new(model.User))

	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm

	return orm, nil
}
