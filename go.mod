module github.com/tykim96/webproject

go 1.20

replace github.com/tykim96/webproject/src => ./src

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.5.1
)
