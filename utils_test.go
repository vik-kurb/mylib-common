package common

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToNullTime_EmptyString(t *testing.T) {
	res := ToNullTime("")
	assert.Equal(t, res, sql.NullTime{})
}

func TestToNullTime_InvalidDate(t *testing.T) {
	res := ToNullTime("1 January")
	assert.Equal(t, res, sql.NullTime{})
}

func TestToNullTime_Success(t *testing.T) {
	res := ToNullTime("10.04.2003")
	assert.Equal(t, res, sql.NullTime{Time: time.Date(2003, time.April, 10, 0, 0, 0, 0, time.UTC), Valid: true})
}
