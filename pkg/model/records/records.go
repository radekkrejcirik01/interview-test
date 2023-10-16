package records

import (
	"encoding/gob"
	"os"
	"time"

	"github.com/samber/lo"
)

const FILE_NAME = "records.bin"

type Record struct {
	ID        int64
	IntValue  int64
	StrValue  string
	BoolValue bool
	TimeValue time.Time
}

// CreateRecord create new record in binary file
func CreateRecord(t *Record) error {
	// Read existing records from the file.
	existingData, err := readBinaryFile()
	if err != nil {
		return err
	}

	if t != nil {
		t.ID = getID(existingData)

		// Append new record to the existing records.
		existingData = append(existingData, *t)
	}

	// Write the combined records back to the file.
	err = writeBinaryFile(existingData)
	return err
}

// GetRecord get record by ID from binary file
func GetRecord(id int64) (*Record, error) {
	// Read existing records from the file.
	data, err := readBinaryFile()
	if err != nil {
		return nil, err
	}

	// Get record
	record := getRecordByID(data, id)

	return record, nil
}

// UpdateRecord update record by ID in binary file
func UpdateRecord(id int64, t *Record) error {
	// Read existing records from the file.
	existingData, err := readBinaryFile()
	if err != nil {
		return err
	}

	// Check if provided record is not empty and update value by ID
	if t != nil {
		for i, v := range existingData {
			if v.ID == id {
				t.ID = id
				existingData[i] = *t
				break
			}
		}
	}

	// Write the adjusted records back to the file.
	err = writeBinaryFile(existingData)
	return err
}

// DeleteRecord delete record by ID in binary file
func DeleteRecord(id int64) error {
	// Read existing records from the file.
	existingData, err := readBinaryFile()
	if err != nil {
		return err
	}

	// Remove value from records array by ID
	records := lo.Filter(existingData, func(x Record, index int) bool {
		return x.ID != id
	})

	// Write the adjusted records back to the file.
	err = writeBinaryFile(records)
	return err
}

// writeBinaryFile write binary file
func writeBinaryFile(data []Record) error {
	// Open the file for writing. Create it if it doesn't exist.
	file, err := os.Create(FILE_NAME)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new encoder for the file.
	encoder := gob.NewEncoder(file)

	// Encode and write the data to the file.
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

// readBinaryFile read binary file and retrieve data
func readBinaryFile() ([]Record, error) {
	var data []Record

	// Open the file for reading.
	file, err := os.Open(FILE_NAME)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, return an empty slice.
			return data, nil
		}
		return data, err
	}
	defer file.Close()

	// Create a new decoder for the file.
	decoder := gob.NewDecoder(file)

	// Decode and read the data from the file.
	if err := decoder.Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}

// getID get new ID from stored data
func getID(data []Record) int64 {
	if len(data) > 0 {
		return data[0].ID + 1
	}
	return 1
}

// getRecordByID get record from list by ID
func getRecordByID(data []Record, id int64) *Record {
	for _, item := range data {
		if item.ID == id {
			return &item
		}
	}
	return nil
}
