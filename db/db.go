package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Record defines the fields of the note.
type Record struct {
	FileID    string    `json:"-"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Records is an array of Record structs.
type Records []Record

// Len returns the number of records in the database.
func (slice Records) Len() int {
	return len(slice)
}

// Less compares two database records and returns whether the first is less than the second one.
func (slice Records) Less(i, j int) bool {
	return strings.ToLower(slice[i].Title) < strings.ToLower(slice[j].Title)
}

// Swap swaps the two database records.
func (slice Records) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// FormattedCreatedAt formats the CreatedAt field into a human readable format.
func (rec *Record) FormattedCreatedAt() string {
	return rec.CreatedAt.Format(time.RFC822)
}

// FormattedUpdatedAt formats the UpdatedAt field into a human readable format.
func (rec *Record) FormattedUpdatedAt() string {
	return rec.UpdatedAt.Format(time.RFC822)
}

// Create writes a new record to the database.
func Create(dataDir string, rec *Record) (string, error) {
	var err error
	var fileID string

	if fileID, err = nextAvailableFileID(dataDir); err != nil {
		return "", err
	}

	if err = writeRec(dataDir, rec, fileID); err != nil {
		return "", err
	}

	return fileID, nil
}

//Count returns the number of records in the database.
func Count(dataDir string) int {
	return len(fileIDsInDataDir(dataDir))
}

// Delete removes a record from the database.
func Delete(dataDir string, fileID string) error {
	var err error
	var filename string

	filename = filePath(dataDir, fileID)

	if err = os.Remove(filename); err != nil {
		return errors.New("(Delete) error doing os.Remove " + err.Error())
	}

	return nil
}

// Find returns the record from the database whose fileID matches the input.
func Find(dataDir string, fileID string) (*Record, error) {
	var data []byte
	var err error
	var filename string
	var rec *Record

	filename = filePath(dataDir, fileID)

	if data, err = ioutil.ReadFile(filename); err != nil {
		return nil, errors.New("(Find) error doing ioutil.ReadFile " + err.Error())
	}

	if err = json.Unmarshal(data, &rec); err != nil {
		return nil, errors.New("(Find) error doing json.Unmarshal " + err.Error())
	}

	rec.FileID = fileID

	return rec, nil
}

// Search returns the record from the database whose Title or Text fields contain the searchString.
func Search(dataDir string, searchString string) ([]Record, error) {
	var data []byte
	var err error
	var fileID string
	var filename string
	var rec Record
	var results Records
	var searchValue string
	var searchValues []string
	var searchValuesCount int
	var valuesFound int

	searchValues = strings.Split(strings.ToLower(searchString), " ")
	searchValuesCount = len(searchValues)

	for _, fileID = range fileIDsInDataDir(dataDir) {
		filename = filePath(dataDir, fileID)

		if data, err = ioutil.ReadFile(filename); err != nil {
			return nil, errors.New("(Search) error doing ioutil.ReadFile " + err.Error())
		}

		if err = json.Unmarshal(data, &rec); err != nil {
			return nil, errors.New("(Search) error doing json.Unmarshal " + err.Error())
		}

		rec.FileID = fileID

		valuesFound = 0

		for _, searchValue = range searchValues {
			if searchValue == "" || strings.Contains(strings.ToLower(rec.Title), searchValue) || strings.Contains(
				strings.ToLower(rec.Text),
				searchValue) {

				valuesFound++
			} else {
				break
			}
		}

		if valuesFound == searchValuesCount {
			results = append(results, rec)
		}
	}

	sort.Sort(results)
	return results, nil
}

// Update updates the record in the database whose fileID matches the input fileID.
func Update(dataDir string, rec *Record, fileID string) error {
	var err error

	if stringInSlice(fileID, fileIDsInDataDir(dataDir)) {
		if err = writeRec(dataDir, rec, fileID); err != nil {
			return err
		}
	} else {
		return errors.New("(Update) error File ID not found")
	}
	return nil
}

//*****************************************************************************
// Private Functions
//*****************************************************************************

// fileIDsInDataDir returns all file ids in the data directory.
func fileIDsInDataDir(dataDir string) []string {
	var file os.FileInfo
	var files []os.FileInfo
	var ids []string

	files, _ = ioutil.ReadDir(dataDir)
	for _, file = range files {
		if !file.IsDir() {
			if path.Ext(file.Name()) == ".json" {
				ids = append(ids, file.Name()[:len(file.Name())-5])
			}
		}
	}

	return ids
}

// filePath returns a file name for a file id.
func filePath(dataDir string, fileID string) string {
	return fmt.Sprintf("%v/%v.json", dataDir, fileID)
}

// loadRec reads a json file into the supplied Note struct.
func loadRec(dataDir string, rec Record, fileID string) error {
	var data []byte
	var err error
	var filename string

	filename = filePath(dataDir, fileID)

	if data, err = ioutil.ReadFile(filename); err != nil {
		return errors.New("(loadRec) error doing ioutil.ReadFile " + err.Error())
	}

	err = json.Unmarshal(data, rec)

	return err
}

// nextAvailableFileID returns the next ascending available file id in a
// directory.
func nextAvailableFileID(dataDir string) (string, error) {
	var err error
	var fileID int
	var fileIDs []int
	var lastFileID int
	var nextFileID string
	var strFileID string

	for _, strFileID = range fileIDsInDataDir(dataDir) {
		if fileID, err = strconv.Atoi(strFileID); err != nil {
			return "", errors.New("(nextAvailableFileID) error doing strconv.Atoi " + err.Error())
		}

		fileIDs = append(fileIDs, fileID)
	}

	if len(fileIDs) == 0 {
		nextFileID = "1"
	} else {
		sort.Ints(fileIDs)
		lastFileID = fileIDs[len(fileIDs)-1]

		nextFileID = strconv.Itoa(lastFileID + 1)
	}

	return nextFileID, nil
}

func stringInSlice(strToFind string, strSlice []string) bool {
	var s string

	for _, s = range strSlice {
		if s == strToFind {
			return true
		}
	}
	return false
}

func writeRec(dataDir string, rec *Record, fileID string) error {
	var err error
	var filename string
	var marshalledRec []byte

	if marshalledRec, err = json.Marshal(rec); err != nil {
		return errors.New("(writeRec) error doing json.Marshal " + err.Error())
	}

	filename = filePath(dataDir, fileID)

	err = ioutil.WriteFile(filename, marshalledRec, 0600)
	if err != nil {
		return errors.New("(writeRec) error doing ioutil.WriteFile " + err.Error())
	}

	return nil
}
