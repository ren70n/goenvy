package goenvy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

	err := json.Unmarshal(data, &content)

	if err != nil {
		return nil, err
	}

	jsonmap := content.(map[string]interface{})

	readTree("", jsonmap, m)

	return m, nil
}

func readTree(key string,
	leaf map[string]interface{},
	m map[string]string) {

	for k, v := range leaf {
		switch v.(type) {
		case string:
			mkey := fmt.Sprintf("%s_%s", key, k)
			mkey = strings.ToUpper(mkey)

			m[mkey] = v.(string)

		case interface{}:

			jmm := v.(map[string]interface{})
			readTree(key+k, jmm, m)
		}
	}
}

func PushToOSEnvs(envsmap map[string]string) {
	for k, v := range envsmap {
		err := os.Setenv(k, v)
		if err != nil {
			fmt.Println(err)
			// return
		}
	}
}
