// CV_project/api/internal/database/db.go

package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
)

// * connection to the database
func ConnectToDatabases() (*sql.DB, error) {
	getEnv := func(key, fallback string) string {
		if value, exists := os.LookupEnv(key); exists {
			return value
		}
		return fallback
	}

	cfg := mysql.Config{
		User:                 getEnv("MYSQL_USER", getEnv("MYSQL_ROOT_USER", "")),
		Passwd:               getEnv("MYSQL_PASSWORD", getEnv("MYSQL_ROOT_PASSWORD", "")),
		Net:                  "tcp",
		Addr:                 getEnv("MYSQL_HOST", "localhost") + ":" + getEnv("MYSQL_PORT", "3306"),
		DBName:               getEnv("MYSQL_DATABASE", ""),
		AllowNativePasswords: true,
	}

	fmt.Println("\n\033[1;34;1m * * * Establishing connection to the database...")
	fmt.Printf("\n\033[1;37;1m * Environment variables print from \033[1;36;1mmain.go:\n\n\033[1;36;1m")
	fmt.Printf("	User:		   ➮ %s \n", cfg.User)
	fmt.Printf("	Password:	   ➮ %s*pass*%s \n", string(cfg.Passwd[0]), string(cfg.Passwd[len(cfg.Passwd)-1]))
	fmt.Printf("	Address:	   ➮ %s \n", cfg.Addr)
	fmt.Printf("	Database Name:     ➮ %s \n\n", cfg.DBName)

	dsn := cfg.FormatDSN()
	maskedPasswd := fmt.Sprintf("%s*pass*%s", string(cfg.Passwd[0]), string(cfg.Passwd[len(cfg.Passwd)-1]))
	maskedDSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", cfg.User, maskedPasswd, cfg.Addr, cfg.DBName, dsn[strings.Index(dsn, "?")+1:])
	fmt.Printf("	DSN:		  \033[1;36;5m ➮ %s\033[0m\n", maskedDSN)

	fmt.Println("\n * Opening database connection...")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return nil, err
	}

	fmt.Println(" * Pinging DB...")
	if err = db.Ping(); err != nil {
		fmt.Printf("\033[31m	Error pinging database: %v\033[0m\n", err)
		db.Close()
		return nil, err
	}

	fmt.Println("\033[1;37;1m * Connecting to database to the address: ➮\033[1;94;1m", cfg.Addr, "\033[0m")
	return db, nil
}
