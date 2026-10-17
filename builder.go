package main
 
import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	table   string
	columns []string
	where   []string
}

func NewQueryBuilder() *QueryBuilder { return &QueryBuilder{} }

func (b *QueryBuilder) Select(cols ...string) *QueryBuilder {
	b.columns = append(b.columns, cols...)
	return b
}

func (b *QueryBuilder) From(table string) *QueryBuilder {
	b.table = table
	return b
}

func (b *QueryBuilder) Where(condition string) *QueryBuilder {
	b.where = append(b.where, condition)
	return b
}

func (b *QueryBuilder) Build() string {
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(b.columns, ", "), b.table)
	if len(b.where) > 0 {
		query += " WHERE " + strings.Join(b.where, " AND ")
	}
	return query
}

func main() {
	q := NewQueryBuilder().
		Select("id", "name").
		From("users").
		Where("age > 18").
		Where("status = 'active'").
		Build()
	
	fmt.Println(q) // SELECT id, name FROM users WHERE age > 18 AND status = 'active'
}