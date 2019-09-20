package goenvy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetEnvsFromJSON parses the JSON and returns map of KEYs/values
// default delimiter for KEY naming is "_"
func GetEnvsFromJSON(file string) (map[string]string, error) {
	return GetEnvsFromJSONDelim(file, "_")
}

// GetEnvsFromJSON parses the JSON and returns map of KEYs/values
// it is possible to define delimiter separating parts of the KEY name
func GetEnvsFromJSONDelim(file string, delim string) (map[string]string, error) {
	data, err := readJSON(file)

	if err != nil {
		return nil, err
	}
	pdata, err := parseJSON(data, delim)
	if err != nil {
		return nil, err
	}

	return pdata, nil
}

// readJSON is a simplest possible file reader
// it takes whatever is passed and returns byte array (or error)
func readJSON(file string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

// parseJSON takes a data array (read from JSON file) and returns
// map of environmental variables. Map is an input for PushToOSEnvs
func parseJSON(data []byte, delim string) (map[string]string, error) {
	m := make(map[string]string)

	var content interface{}

	if err := json.Unmarshal(data, &content); err != nil {
		return nil, err
	}
	toPaths(content, &m, delim, nil)
	return m, nil
}

// PushToOSEnvs sets environmental variables passed in map in OS
func PushToOSEnvs(envsmap map[string]string) {
	for k, v := range envsmap {
		err := os.Setenv(k, v)
		if err != nil {
			log.Println(err)
		}
	}
}

// toPaths implements DFS over arbitrary json structure and flattens results
// to PATH-value pairs
// Path parts are separated by delimiter
func toPaths(in interface{}, res *map[string]string, delim string, path []string) {
	if path == nil {
		path = []string{}
	}
	switch t := in.(type) {
	case []interface{}:
		for i, v := range t {
			toPaths(v, res, delim, copyAppend(path, strconv.Itoa(i)))
		}
	case map[string]interface{}:
		for k, v := range t {
			toPaths(v, res, delim, copyAppend(path, k))
		}
	default:
		(*res)[strings.ToUpper(strings.Join(path, delim))] = fmt.Sprint(t)
	}
}

// copyAppend is a utility func for toPaths
// create a slice one longer than src
// copy src to dst
// set elem at tail
// return
func copyAppend(src []string, elem string) (dst []string) {
	l := len(src)
	dst = make([]string, l+1)
	copy(dst, src)
	dst[l] = elem
	return
}
