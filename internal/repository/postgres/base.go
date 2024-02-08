package postgres

import (
	"log"
	"strconv"
	"strings"

	"github.com/engagerocketco/templates-api-svc/internal/repository/repomodel"

	_ "github.com/lib/pq"

	// dot import so that jet go code would resemble as much as native SQL
	// dot import is not mandatory
	. "github.com/go-jet/jet/v2/postgres"
)

// Assuming you have a function for an IntColumn
func handleIntColumn(column ColumnInteger, value int32) BoolExpression {
	return column.EQ(Int32(value))
}

// And similarly for a StringColumn
func handleStringColumn(column ColumnString, value string) BoolExpression {
	return column.LIKE(String("%" + value + "%"))
}

func convertToIntPointer(value int32) *int {
	convertedValue := int(value) // Convert int32 to int
	return &convertedValue       // Return the address of the converted value
}

func GenerateDynamicOrderByClause(table_columns []Column, req repomodel.PaginationRequest) OrderByClause {

	var column Column
	var id Column
	for _, table_column := range table_columns {
		if table_column.Name() == "id" {
			id = table_column
		}
		if req.SortBy == table_column.Name() {
			column = table_column
		}
	}

	if column == nil {
		return id // Column not found, skip
	}

	// Determine sort direction
	direction := strings.ToUpper(strings.TrimSpace(req.SortType)) // Assuming this applies universally, adjust logic as needed for per-column direction

	switch direction {
	case "DESC":
		return column.DESC()
	default:
		return column.ASC()
	}
}

func GenerateDynamicWhereClause(table_columns []Column, req repomodel.PaginationRequest) BoolExpression {

	search_columns := req.SearchBy
	values := req.SearchValue
	operatorType := req.SearchLogicOpeator

	if len(values) != len(search_columns) {
		// Handle error: column and value slices must have the same length
		return nil
	}

	var expressions []BoolExpression

	for i, search_column := range search_columns {

		next_column := true
		var column Column
		for _, table_column := range table_columns {
			if search_column == table_column.Name() {
				next_column = false
				column = table_column
				break
			}
		}

		if next_column {
			continue
		}

		value := values[i]

		log.Println(search_column)
		log.Println(value)

		switch typedColumn := column.(type) {
		case ColumnInteger:
			i, err := strconv.Atoi(value)
			if err != nil {
				return nil //TODO
			}
			expressions = append(expressions, handleIntColumn(typedColumn, int32(i)))
		case ColumnString:
			expressions = append(expressions, handleStringColumn(typedColumn, value))
			// handle other types
		}
	}

	if len(expressions) == 0 {
		return nil
	}

	var combinedExpression BoolExpression = expressions[0]
	for _, expr := range expressions[1:] {
		if operatorType == "OR" {
			combinedExpression = combinedExpression.OR(expr)
		} else { // Defaults to "AND"
			combinedExpression = combinedExpression.AND(expr)
		}
	}
	return combinedExpression
}
