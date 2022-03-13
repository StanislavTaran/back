migrateup:
migrate -path migrations -database "mysql://root:test@tcp(localhost:3306)/testDb?charset=utf8mb4&parseTime=True&loc=Local" up
migratedown:
migrate -path migrations -database "mysql://root:test@tcp(localhost:3306)/testDb?charset=utf8mb4&parseTime=True&loc=Local" down