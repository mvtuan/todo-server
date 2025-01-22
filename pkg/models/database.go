package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"server/pkg/common"
)

type Database interface {
	Create(data any) *common.APIResponse
	Query(filter any, offset, limit int) *common.APIResponse
	QueryOne(filter any) *common.APIResponse
	Update(filter, updater any) *common.APIResponse
	Migration() error
}

type DB struct {
	db *gorm.DB
}

func (db *DB) Create(data any) *common.APIResponse {
	if data == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.Invalid,
			Message: "Invalid input",
		}
	}
	result := db.db.Create(data)
	if result.Error != nil {
		return &common.APIResponse{
			Status:    common.APIStatus.InternalServerError,
			Message:   "Something went wrong",
			RootCause: result.Error,
		}
	}

	return &common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Data inserted successfully",
		Data:    data,
	}
}

func (db *DB) Query(filter any, offset, limit int) *common.APIResponse {
	query := db.db.Model(filter).Limit(limit).Offset(offset)

	// Apply additional filters (if any)
	if filter != nil {
		query = query.Where(filter)
	}

	// Execute the query
	result := query.Find(filter)

	// Check if no records are found
	if result.RowsAffected == 0 {
		return &common.APIResponse{
			Status:  common.APIStatus.NotFound,
			Message: "Data not found",
		}
	}

	// Check for any other errors
	if result.Error != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.InternalServerError,
			Message: result.Error.Error(),
		}
	}

	return &common.APIResponse{
		Status: common.APIStatus.Ok,
		Data:   filter,
	}
}

func (db *DB) QueryOne(filter any) *common.APIResponse {
	query := db.db.Model(filter).Limit(1).Offset(0)

	// Apply additional filters (if any)
	if filter != nil {
		query = query.Where(filter)
	}

	// Execute the query
	result := query.Find(filter)

	// Check if no records are found
	if result.RowsAffected == 0 {
		return &common.APIResponse{
			Status:  common.APIStatus.NotFound,
			Message: "Data not found",
		}
	}

	// Check for any other errors
	if result.Error != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.InternalServerError,
			Message: result.Error.Error(),
		}
	}

	return &common.APIResponse{
		Status: common.APIStatus.Ok,
		Data:   filter,
	}
}

func (db *DB) Update(filter, updater any) *common.APIResponse {
	query := db.db.Model(filter).Where(filter)
	query.Updates(updater)

	return &common.APIResponse{
		Status: common.APIStatus.Ok,
		Data:   filter,
	}
}

func (db *DB) Migration() error {
	if exist := db.db.Migrator().HasTable(&Task{}); exist == false {
		err := db.db.Migrator().CreateTable(&Task{})
		if err != nil {
			return err
		}
	}
	if exist := db.db.Migrator().HasTable(&User{}); exist == false {
		err := db.db.Migrator().CreateTable(&User{})
		if err != nil {
			return err
		}
	}

	return nil
}

func NewDB() (Database, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to database")
	return &DB{db: db}, nil
}
