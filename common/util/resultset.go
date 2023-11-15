package util

import (
	"database/sql/driver"
	"strconv"
	"time"

	"github.com/godror/godror"
)

type ResultSet struct {
	colNums map[string]int
	vals    []driver.Value
	rows    driver.Rows
}

func NewResultSet(rows driver.Rows) *ResultSet {
	cols := rows.(driver.RowsColumnTypeScanType).Columns()
	colNums := make(map[string]int, len(cols))
	for i := 0; i < len(cols); i++ {
		colNums[cols[i]] = i
	}
	vals := make([]driver.Value, len(cols))

	return &ResultSet{
		colNums: colNums,
		vals:    vals,
		rows:    rows,
	}
}

func (rs *ResultSet) Close() {
	rs.rows.Close()
}

func (rs *ResultSet) Next() error {
	return rs.rows.Next(rs.vals)
}

func (rs *ResultSet) GetString(fieldName string) string {
	return rs.vals[rs.colNums[fieldName]].(string)
}

func (rs *ResultSet) GetDouble(fieldName string) float64 {
	f, err := strconv.ParseFloat(string(rs.vals[rs.colNums[fieldName]].(godror.Number)), 64)
	if err != nil {
		return 0
	}
	return f
}

func (rs *ResultSet) GetInt(fieldName string) int {
	i, err := strconv.Atoi(string(rs.vals[rs.colNums[fieldName]].(godror.Number)))
	if err != nil {
		return 0
	}
	return i
}

func (rs *ResultSet) GetLong(fieldName string) int64 {
	l, err := strconv.ParseInt(string(rs.vals[rs.colNums[fieldName]].(godror.Number)), 10, 64)
	if err != nil {
		return 0
	}
	return l
}

func (rs *ResultSet) GetTime(fieldName string) *time.Time {
	v := rs.vals[rs.colNums[fieldName]]
	if v != nil {
		t, err := v.(time.Time)
		if !err {
			return nil
		}
		return &t
	} else {
		return nil
	}
}
