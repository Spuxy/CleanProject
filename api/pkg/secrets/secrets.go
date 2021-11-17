package secrets

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type DotSecret struct {
	Dir    string
	Values map[string]string
}

func NewDotSecret(path string) (*DotSecret, error) {
	ds := &DotSecret{Dir: path, Values: map[string]string{}}
	ds.readAll()
	return ds, nil
}

func (d *DotSecret) Get(key string) (string, error) {
	if _, ok := d.Values[key]; !ok {
		return "", fmt.Errorf("the key:%s was not found", key)
	}

	return d.Values[key], nil
}

func (d *DotSecret) readAll() error {
	err := d.isDir()
	if err != nil {
		fmt.Println(err)
		return err
	}

	files, err := ioutil.ReadDir(d.Dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := d.read(file.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DotSecret) read(file string) error {
	value, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", d.Dir, file))
	if err != nil {
		return err
	}

	d.Values[file] = strings.TrimSpace(string(value))
	return nil
}

func (d DotSecret) isDir() error {
	file, err := os.Stat(d.Dir)
	if err != nil {
		return err
	}
	if !file.IsDir() {
		return fmt.Errorf("%s is not directory", d.Dir)
	}
	return nil
}
