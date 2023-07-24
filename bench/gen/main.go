package main

import (
	"gorm.io/gen"

	"github.com/efectn/go-orm-benchmarks/bench/gen/models"
)

func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})
	//
	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN:                  helper.OrmSource,
	// 	PreferSimpleProtocol: true, // disables implicit prepared statement usage
	// }), &gorm.Config{
	// 	SkipDefaultTransaction: true,
	// 	PrepareStmt:            false,
	// 	Logger:                 logger.Default.LogMode(logger.Silent),
	// })
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // Use the above `*gorm.DB` instance to initialize the generator,
	// // which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	// g.UseDB(db)

	// Generate default DAO interface for those specified structs
	g.ApplyBasic(models.Model{})

	// Execute the generator
	g.Execute()
}
