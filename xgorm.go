package xc

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type GormConfig struct {
	Dsn         string   `json:"Dsn"`
	DsnReplicas []string `json:"DsnReplicas,optional"`
}

func GormNewClient(c GormConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 判断是否集群
	if len(c.DsnReplicas) > 0 {
		var replicas []gorm.Dialector
		for _, replica := range c.DsnReplicas {
			replicas = append(replicas, mysql.Open(replica))
		}
		err = db.Use(dbresolver.Register(dbresolver.Config{
			Sources:           []gorm.Dialector{mysql.Open(c.Dsn)},
			Replicas:          replicas,
			Policy:            dbresolver.RandomPolicy{},
			TraceResolverMode: true,
		}))
		if err != nil {
			panic(err)
		}
	}

	return db
}

func GormFirstErr(err error) error {
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return nil
}

func GormDbErr(err error) error {
	if err != nil {
		return err
	}
	return nil
}

// 生成模型
func GormGeneratorModel(c GormConfig, outPath string) {
	g := gen.NewGenerator(gen.Config{
		OutPath: outPath,
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	dsn := c.Dsn
	gormDb, _ := gorm.Open(mysql.Open(dsn))
	g.UseDB(gormDb)
	g.WithDataTypeMap(map[string]func(gorm.ColumnType) (dataType string){
		"decimal": func(columnType gorm.ColumnType) (dataType string) {
			return "decimal.Decimal"
		},
	})
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
