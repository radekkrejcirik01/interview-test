package records

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRecord(t *testing.T) {
	record := Record{
		IntValue:  42,
		StrValue:  "foo",
		BoolValue: true,
	}

	tests := []struct {
		name     string
		provided Record
		expected error
	}{
		{
			name:     "no provided record",
			provided: Record{},
			expected: nil,
		},
		{
			name:     "create new record",
			provided: record,
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, CreateRecord(&test.provided))
		})
	}
}

func TestGetRecord(t *testing.T) {
	record := Record{
		ID:        1,
		IntValue:  42,
		StrValue:  "foo",
		BoolValue: true,
	}

	tests := []struct {
		name           string
		provided       int64
		expectedRecord *Record
		expectedErr    error
	}{
		{
			name:           "no record for provided id",
			provided:       12,
			expectedRecord: nil,
			expectedErr:    nil,
		},
		{
			name:           "retrieve record",
			provided:       1,
			expectedRecord: &record,
			expectedErr:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			record, err := GetRecord(test.provided)

			assert.Equal(t, test.expectedRecord, record)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestUpdateRecord(t *testing.T) {
	record := Record{
		ID:        1,
		IntValue:  42,
		StrValue:  "foo",
		BoolValue: true,
	}

	tests := []struct {
		name          string
		providedId    int64
		proviedRecord *Record
		expected      error
	}{
		{
			name:          "no data provided",
			providedId:    1,
			proviedRecord: nil,
			expected:      nil,
		},
		{
			name:          "no record for provided id",
			providedId:    12,
			proviedRecord: &record,
			expected:      nil,
		},
		{
			name:          "update record",
			providedId:    1,
			proviedRecord: &record,
			expected:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, UpdateRecord(test.providedId, test.proviedRecord))
		})
	}
}

func TestDeleteRecord(t *testing.T) {
	tests := []struct {
		name     string
		provided int64
		expected error
	}{
		{
			name:     "no record for provided id",
			provided: 12,
			expected: nil,
		},
		{
			name:     "delete record",
			provided: 1,
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, DeleteRecord(test.provided))
		})
	}
}

func TestWriteBinaryFile(t *testing.T) {
	data := []Record{{
		ID:        1,
		IntValue:  42,
		StrValue:  "foo",
		BoolValue: true,
	}}

	tests := []struct {
		name     string
		provided []Record
		expected error
	}{
		{
			name:     "no provided record",
			provided: []Record{},
			expected: nil,
		},
		{
			name:     "write binary file",
			provided: data,
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, writeBinaryFile(test.provided))
		})
	}
}

func TestReadBinaryFile(t *testing.T) {
	records := []Record{{
		ID:        1,
		IntValue:  42,
		StrValue:  "foo",
		BoolValue: true,
	}}

	tests := []struct {
		name            string
		expectedRecords []Record
		expectedErr     error
	}{
		{
			name:            "read binary file",
			expectedRecords: records,
			expectedErr:     nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			record, err := readBinaryFile()

			assert.Equal(t, test.expectedRecords, record)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestGetId(t *testing.T) {
	records := []Record{{
		ID:        1,
		IntValue:  42,
		StrValue:  "foo",
		BoolValue: true,
	}}

	tests := []struct {
		name     string
		provided []Record
		expected int64
	}{
		{
			name:     "no previous ids",
			provided: []Record{},
			expected: 1,
		},
		{
			name:     "previous id 1",
			provided: records,
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, getID(test.provided))
		})
	}
}

func TestGetRecordByID(t *testing.T) {
	record := Record{
		ID:        1,
		IntValue:  42,
		StrValue:  "foo",
		BoolValue: true,
	}

	data := []Record{record}

	tests := []struct {
		name         string
		providedData []Record
		providedID   int64
		expected     *Record
	}{
		{
			name:         "no records",
			providedData: []Record{},
			providedID:   1,
			expected:     nil,
		},
		{
			name:         "no record for provided id",
			providedData: data,
			providedID:   2,
			expected:     nil,
		},
		{
			name:         "get record",
			providedData: data,
			providedID:   1,
			expected:     &record,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, getRecordByID(test.providedData, test.providedID))
		})
	}
}
