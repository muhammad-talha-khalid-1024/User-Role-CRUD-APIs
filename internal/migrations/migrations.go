package migrations

import (
	"database/sql"
	"fmt"
	"mustaqel/internal/config"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// func Migrations(db *sql.DB, cfg *config.Config) error {
// 	// Find the correct migration path
// 	path, err := findMigrationPath()
// 	if err != nil {
// 		return err
// 	}

// 	// Convert to URL format with proper encoding for Windows
// 	absPath, err := filepath.Abs(path)
// 	if err != nil {
// 		return err
// 	}

// 	// IMPORTANT: For Windows, use triple slash and URL encode spaces
// 	urlPath := "file:///" + filepath.ToSlash(absPath)
// 	urlPath = strings.ReplaceAll(urlPath, " ", "%20")

// 	log.Printf("Migration path: %s", urlPath)

// 	// Create MySQL driver instance
// 	driver, err := mysql.WithInstance(db, &mysql.Config{
// 		DatabaseName: cfg.DBName,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	// Create migrate instance with file source
// 	m, err := migrate.NewWithDatabaseInstance(
// 		urlPath,
// 		"mysql",
// 		driver,
// 	)
// 	if err != nil {
// 		return fmt.Errorf("failed to create migrate instance: %w", err)
// 	}

// 	// Run migrations
// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 		return err
// 	}

// 	if err == migrate.ErrNoChange {
// 		log.Println("No new migrations to apply")
// 	} else {
// 		log.Println("Migrations applied successfully")
// 	}

// 	return nil
// }

// func findMigrationPath() (string, error) {
// 	// Get current working directory
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		return "", err
// 	}

// 	log.Printf("Current working directory: %s", cwd)

// 	// Try different possible paths (using filepath.Join for Windows compatibility)
// 	possiblePaths := []string{
// 		filepath.Join(cwd, "migrations"),
// 		filepath.Join(cwd, "../migrations"),
// 		filepath.Join(cwd, "../../migrations"),
// 		"./migrations",
// 		"../migrations",
// 		"migrations",
// 	}

// 	for _, path := range possiblePaths {
// 		cleanPath := filepath.Clean(path)
// 		if info, err := os.Stat(cleanPath); err == nil && info.IsDir() {
// 			// Check if there are any .sql files
// 			files, err := os.ReadDir(cleanPath)
// 			if err == nil && len(files) > 0 {
// 				log.Printf("✓ Found migrations at: %s (contains %d files)", cleanPath, len(files))
// 				return cleanPath, nil
// 			}
// 		}
// 	}

// 	return "", fmt.Errorf("migrations directory not found in: %v", possiblePaths)
// }

func Migrations(db *sql.DB, cfg *config.Config) error {
	var folder = "./migrations"
	entries, err := os.ReadDir(folder)
	if err != nil {
		return fmt.Errorf("read SQL folder: %w", err)
	}

	var files []string

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if strings.HasSuffix(entry.Name(), ".sql") {
			files = append(files, entry.Name())
		}
	}

	// Ensure predictable execution order
	sort.Strings(files)

	for _, file := range files {
		path := filepath.Join(folder, file)

		fmt.Println("Running:", path)

		sqlBytes, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", file, err)
		}

		if _, err := db.Exec(string(sqlBytes)); err != nil {
			return fmt.Errorf("execute %s: %w", file, err)
		}

		fmt.Println("Completed:", file)
	}

	return nil
}
