package db

import (
	"fmt"
	"strings"

	"github.com/fatih/structs"
)

type ModelController interface {
	Save(tableName string, inter interface{}) (string, []interface{}, error)
	Update(tableName string, inter interface{}) (string, []interface{}, error)
	Delete(tableName string, inter interface{}) (string, []interface{}, error)
	Get(tableName string, inter interface{}) (string, []interface{}, error)
}

type HelperDBModel struct {
	TableName string
}

func (m *HelperDBModel) InterToFields(inter interface{}, withDollar bool) (fields []string, values []interface{}) {
	fields = make([]string, 0)
	values = make([]interface{}, 0)
	i := 0
	// Get db tag name
	for _, field := range structs.Fields(inter) {
		if field.IsExported() {
			if withDollar {
				fields = append(fields, fmt.Sprintf("%s = $%d", field.Tag("db"), i+1))
			} else {
				fields = append(fields, field.Tag("db"))
			}
			values = append(values, field.Value())
			i++
		}
	}

	return fields, values
}
func (m *HelperDBModel) RepeatDollarid(n int, sep string) string {
	// $1, $2...
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		arr[i] = "$" + fmt.Sprintf("%d", i+1)
	}
	return strings.Join(arr, sep)
}
func (m *HelperDBModel) Save(tableName string, inter interface{}) (string, []interface{}, error) {
	fields, values := m.InterToFields(inter, false)

	return fmt.Sprintf("insert into %s (%s) values (%s)", tableName, strings.Join(fields, ", "), m.RepeatDollarid(len(fields), ", ")), values, nil
}

func (m *HelperDBModel) Update(tableName string, inter interface{}) (string, []interface{}, error) {
	fields, values := m.InterToFields(inter, true)

	return fmt.Sprintf("update %s set %s", tableName, strings.Join(fields, ", ")), values, nil
}

func (m *HelperDBModel) Delete(tableName string, inter interface{}) (string, []interface{}, error) {
	fields, values := m.InterToFields(inter, true)

	return fmt.Sprintf("delete from %s where %s", tableName, strings.Join(fields, " and ")), values, nil
}

func (m *HelperDBModel) Get(tableName string, pickList []string, inter interface{}) (string, []interface{}, error) {
	fields, values := m.InterToFields(inter, true)
	pickListStr := "*"
	if len(pickList) > 0 {
		pickListStr = strings.Join(pickList, ", ")
	}
	return fmt.Sprintf("select %s from %s where %s", pickListStr, tableName, strings.Join(fields, " and ")), values, nil
}
