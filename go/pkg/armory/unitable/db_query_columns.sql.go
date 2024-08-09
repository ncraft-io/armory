// Code generated by mojo. DO NOT EDIT.
// Rerunning mojo will overwrite this file.

package unitable

import (
	"database/sql/driver"
	"github.com/mojo-lang/db/go/pkg/mojo/db"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func (x DbQueryColumns) Value() (driver.Value, error) {
	return db.JSONValuer{}.Value(x)
}

func (x *DbQueryColumns) Scan(value interface{}) error {
	return db.JSONScanner{}.Scan(x, value)
}

func (x DbQueryColumns) GormDBDataType(gdb *gorm.DB, field *schema.Field) string {
	return db.JSONDbDataType{}.GormDBDataType(gdb, field)
}

func (x DbQueryColumns) GormDataType() string {
	return "string"
}