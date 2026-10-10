package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Config struct {
	DBHost, DBPort, DBUsername, DBPassword, DBName, AppPort string
}

// load config
func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"), // Key + fallback
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUsername: getEnv("DB_USERNAME", "root"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBName:     getEnv("DB_NAME", "sprig_db"),
		AppPort:    getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}

// Connect Driver SQL
func OpenConnection(cfg *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Gagal membuka database : %w", err)
	}

	// log.Printf("DSN DETAIL : %s", dsn)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Gagal terhubung ke database : %w", err)
	}

	// ---- AUTO MIGRATE DATABASE
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Gagal membuat driver database : %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Gagal init migrate : %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Gagal menjalankan migrasi database : %v", err)
	}

	log.Println("migrasi berhasil / sudah up-to-date")
	// ---- END

	db.SetMaxIdleConns(10)                  // kalau lagi bengong, minimal buka koneksinya brp..
	db.SetMaxOpenConns(100)                 // maksimal buka koneksi
	db.SetConnMaxIdleTime(3 * time.Minute)  // kalau request ada yang bengong, minimal brp mnt/dtk untuk di close koneksinya
	db.SetConnMaxLifetime(60 * time.Minute) // maksimal semua request untuk konek ke db

	return db, nil
}
