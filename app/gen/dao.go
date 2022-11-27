package gen

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenerateDao(dsn string) (err error) {
	var db *gorm.DB
	if db, err = gorm.Open(mysql.Open(dsn)); err != nil {
		return
	}

	var tableList []string
	if tableList, err = db.Migrator().GetTables(); err != nil {
		return
	}

	generator := gen.NewGenerator(gen.Config{
		ModelPkgPath: "model/entity",
	})

	generator.UseDB(db)

	for _, tableName := range tableList {
		generator.GenerateModel(tableName)
	}

	generator.Execute()
	return
}
