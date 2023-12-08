package sqlu

import (
	"context"
	"database/sql"
	"strings"
)

type Execer interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
}

type Fields map[string]any

// Update builds an UPDATE statement.
func Update(tableName string) *stmt {
	return &stmt{
		tableName: quote(tableName),
		fields:    make(map[string]any),
		where:     make(map[string][]any),
	}
}

type stmt struct {
	tableName string
	fields    map[string]any
	where     map[string][]any
}

func (q *stmt) Set(field string, value any) *stmt {
	q.fields[quote(field)] = value
	return q
}

func (q *stmt) Setm(fields Fields) *stmt {
	for field, value := range fields {
		q.Set(field, value)
	}
	return q
}

func (q *stmt) Where(expr string, args ...any) *stmt {
	q.where[expr] = args
	return q
}

func (q *stmt) Exec(ctx context.Context, db Execer) (sql.Result, error) {
	sql, args := q.Build()
	return db.ExecContext(ctx, sql, args...)
}

func (q *stmt) Build() (string, []any) {
	var (
		fields []string
		where  []string
		args   []any
	)

	for field, value := range q.fields {
		fields = append(fields, field+"=?")
		args = append(args, value)
	}

	for expr, values := range q.where {
		where = append(where, expr)
		args = append(args, values...)
	}

	sql := "UPDATE " + q.tableName +
		" SET " + strings.Join(fields, ", ") +
		" WHERE " + strings.Join(where, " AND ")

	return sql, args
}

func quote(s string) string {
	s = strings.Trim(s, "`")
	return "`" + s + "`"
}
