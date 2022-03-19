### MIGRATIONS

Create new migration example
```bash
migrate create -ext sql -dir migrations -seq migation_name
```
<br />

Run migrations
### up
```bash
migrate -path migrations -database "mysql://root:test@tcp(localhost:3306)/legendary_waffle?charset=utf8mb4&parseTime=True&loc=Local" up
```

### down
```bash
migrate -path migrations -database "mysql://root:test@tcp(localhost:3306)/legendary_waffle?charset=utf8mb4&parseTime=True&loc=Local" up
```
<br />

### DUMP MYSQL
```
docker exec mysql sh -c 'exec mysqldump --all-databases -uroot -p"$MYSQL_ROOT_PASSWORD"' > ./backups/backup.sql
```