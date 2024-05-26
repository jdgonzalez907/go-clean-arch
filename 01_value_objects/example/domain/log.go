package domain

import "fmt"

type Log struct {
	field   Field
	oldData *string
	newData *string
}

func NewLog(field string, oldData, newData *string) (Log, error) {
	newField, err := NewField(field)
	if err != nil {
		return Log{}, err
	}
	return Log{field: newField, oldData: oldData, newData: newData}, nil
}

func (l Log) Field() Field {
	return l.field
}

func (l Log) OldData() *string {
	return l.oldData
}

func (l Log) NewData() *string {
	return l.newData
}

func (l Log) Equals(other Log) bool {
	return l.field == other.Field() &&
		l.oldData == other.OldData() &&
		l.newData == other.NewData()
}

func (l Log) String() string {
	var oldData, newData string

	if l.oldData != nil {
		oldData = *l.oldData
	}
	if l.newData != nil {
		newData = *l.newData
	}

	return fmt.Sprintf("Log{field: %s, oldData: %s, newData: %s}", l.field.Value(), oldData, newData)
}
