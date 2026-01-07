package initializers

import (
    "go-authentication-api/models"
    "log"
    "gorm.io/gorm"
)

func SyncDatabase() {
    // Auto-migrate schema for User model
    if err := DB.AutoMigrate(&models.User{}); err != nil {
        log.Fatalf("AutoMigrate failed: %v", err)
    }

    // Ensure legacy columns are removed when not present in the model
    // Drop the 'name' column if it exists, since the model no longer has it
    migrator := DB.Migrator()
    if migrator.HasColumn(&models.User{}, "name") {
        if err := migrator.DropColumn(&models.User{}, "name"); err != nil && err != gorm.ErrPrimaryKeyRequired {
            log.Printf("Warning: failed to drop 'name' column: %v", err)
        }
    }

    // Drop the legacy 'email' column if present to avoid NOT NULL insertion errors
    if migrator.HasColumn(&models.User{}, "email") {
        if err := migrator.DropColumn(&models.User{}, "email"); err != nil && err != gorm.ErrPrimaryKeyRequired {
            log.Printf("Warning: failed to drop 'email' column: %v", err)
        }
    }
}