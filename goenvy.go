package goenvy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// GetEnvsFromJSON ...
func GetEnvsFromJSON(file string) (map[string]string, error) {
	data, err := readJSON(file)

	if err != nil {
		return nil, err
	}
	pdata, err := parseJSON(data)
	if err != nil {
		return nil, err
	}

	return pdata, nil
}

func readJSON(file string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func parseJSON(data []byte) (map[string]string, error) {
	m := make(map[string]string)

	var content interface{}

	if err := json.Unmarshal(data, &content); err != nil {
		return nil, err
	}
	toPaths(content, &m, "_", nil)
	return m, nil
}

// PushToOSEnvs ...
func PushToOSEnvs(envsmap map[string]string) {
	for k, v := range envsmap {
		err := os.Setenv(k, v)
		if err != nil {
			fmt.Println(err)
			// return
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
