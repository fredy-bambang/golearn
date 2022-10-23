package database

import (
	"path/filepath"
	"runtime"
)

// MigrationsDir returns database migration directory.
func MigrationsDir() string {
	return migrationsDir()
}

func migrationsDir() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return filepath.Join(filepath.Dir(filename), "./migrations")
}
