package gorm

type Orm interface {
	TableName() string
}

func ParseTableName(v interface{}) string {
	table, ok := v.(Orm)
	if !ok {
		return "default"
	}
	return table.TableName()
}
