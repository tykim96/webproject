module github.com/tykim96/webproject

go 1.20

replace github.com/tykim96/webproject/src => ./src

require github.com/tykim96/webproject/src v0.0.0-00010101000000-000000000000

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)
