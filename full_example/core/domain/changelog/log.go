package changelog

import "fmt"

type Log struct {
	field    Field
	oldValue *string
	newValue *string
}

func NewLog(field string, oldValue, newValue *string) (Log, error) {
	f, err := NewField(field)
	if err != nil {
		return Log{}, err
	}

	return Log{
		field:    f,
		oldValue: oldValue,
		newValue: newValue,
	}, nil
}

func (l Log) Field() Field {
	return l.field
}

func (l Log) OldValue() *string {
	return l.oldValue
}

func (l Log) NewValue() *string {
	return l.newValue
}

func (l Log) String() string {
	var strOldValue, strNewValue string

	if l.oldValue != nil {
		strOldValue = *l.oldValue
	}

	if l.newValue != nil {
		strNewValue = *l.newValue
	}

	return fmt.Sprintf("Log{field: %s, oldValue: %s, newValue: %s}", l.field.String(), strOldValue, strNewValue)
}
