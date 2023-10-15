package records

import (
	"encoding/gob"
	"os"
	"time"
)

const FILE_NAME = "records.bin"

type Record struct {
	ID        int64
	IntValue  int64
	StrValue  string
	BoolValue bool
	TimeValue time.Time
}

// CreateRecord writes new record to binary file
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

// GetRecord gets recordy by ID from binary file
func GetRecord(id int64) (*Record, error) {
	// Read existing records from the file.
	data, err := readBinaryFile()
	if err != nil {
		return nil, err
	}

	// Get records
	record := getRecordByID(data, id)

	return record, nil
}

// UpdateRecord updates record by ID in binary file
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
			}
		}
	}

	// Write the adjusted records back to the file.
	err = writeBinaryFile(existingData)
	return err
}

// DeleteRecord deletes record by ID in binary file
func DeleteRecord(id int64) error {
	var data []Record

	// Read existing records from the file.
	existingData, err := readBinaryFile()
	if err != nil {
		return err
	}

	// Remove value from records array by ID
	for _, v := range existingData {
		if v.ID != id {
			data = append(data, v)
		}
	}

	// Write the adjusted records back to the file.
	err = writeBinaryFile(data)
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

// getID helper function for retrieving id by already stored data
func getID(data []Record) int64 {
	if len(data) > 0 {
		return data[0].ID + 1
	}
	return 1
}

// getRecordByID helper function to get recordy by ID
func getRecordByID(data []Record, id int64) *Record {
	for _, item := range data {
		if item.ID == id {
			return &item
		}
	}
	return nil
}
