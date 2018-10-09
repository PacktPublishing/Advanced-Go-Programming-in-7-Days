package gomongo

import (
	"sync"
	"log"
	"path/filepath"
	"os"
	"fmt"
	"bytes"
	"encoding/json"
	"io/ioutil"
)

const (
	dbExtension = ".json"
	tempExtension = ".tmp"
	permissionDirectory = 0755
	permissionFile = 0644
)

type Driver struct {
	mu      *sync.Mutex
	mutexes map[string]*sync.Mutex
	baseDir string
}

// New creates a new database at the specified directory location
func New(dir string) (*Driver, error) {
	dir = filepath.Clean(dir)

	driver := Driver{
		baseDir: dir,
		mu:      &sync.Mutex{},
		mutexes: make(map[string]*sync.Mutex),
	}

	// if the database already exists, just use it
	if _, err := os.Stat(dir); err == nil {
		log.Printf("Using '%s'", dir)
		return &driver, nil
	}

	// if the database doesn't exist create it
	log.Printf("Creating database at '%s'...\n", dir)
	return &driver, os.MkdirAll(dir, permissionDirectory)
}

// Write locks the database and attempts to write the record to the database under
// the [collection] specified with the [resource] name given
func (d *Driver) Write(document, resource string, v []byte) error {
	if document == "" || resource == "" {
		return fmt.Errorf("Read: missing collection/resource name")
	}

	mu := d.getDocumentMutex(document)
	mu.Lock()
	defer mu.Unlock()

	dir := filepath.Join(d.baseDir, document)
	fnlPath := filepath.Join(dir, resource + dbExtension)
	tmpPath := fnlPath + tempExtension

	// create collection directory
	if err := os.MkdirAll(dir, permissionDirectory); err != nil {
		return err
	}

	buf := bytes.Buffer{}
	raw := json.RawMessage(v)
	err := json.NewEncoder(&buf).Encode(&raw)
	if err != nil {
		return err
	}

	// write marshaled data to the temp file
	if err := ioutil.WriteFile(tmpPath, buf.Bytes(), permissionFile); err != nil {
		return err
	}

	// move final file into place
	return os.Rename(tmpPath, fnlPath)
}

// Read a record from the database
func (d *Driver) Read(document, resource string, v interface{}) error {
	if document == "" || resource == "" {
		return fmt.Errorf("Read: missing collection/resource name")
	}

	record := filepath.Join(d.baseDir, document, resource)

	// check to see if file exists
	if _, err := stat(record); err != nil {
		return err
	}

	// read record from database
	b, err := ioutil.ReadFile(record + dbExtension)
	if err != nil {
		return err
	}

	buf := bytes.NewReader(b)
	err = json.NewDecoder(buf).Decode(&v)
	if err != nil {
		return err
	}

	return err
}

func (d *Driver) Delete(document, resource string) error {
	path := filepath.Join(document, resource)

	mu := d.getDocumentMutex(document)
	mu.Lock()
	defer mu.Unlock()

	dir := filepath.Join(d.baseDir, path)

	info, err := stat(dir)
	if err != nil {
		return fmt.Errorf("Delete: Cannot find file or directory named %v\n", path)
	}

	if info.Mode().IsDir() {
		err = os.RemoveAll(dir)
		return err
	}

	if info.Mode().IsRegular() {
		err = os.RemoveAll(dir + dbExtension)
		return err
	}

	return nil
}

func stat(path string) (os.FileInfo, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		info, err = os.Stat(path + dbExtension)
	}

	return info, err
}

func (d *Driver) getDocumentMutex(document string) *sync.Mutex {
	d.mu.Lock()
	defer d.mu.Unlock()

	m, ok := d.mutexes[document]
	if !ok {
		m = &sync.Mutex{}
		d.mutexes[document] = m
	}

	return m
}

func (d *Driver) Close() error {
	return nil
}