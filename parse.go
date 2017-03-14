/*
 * Copyright 2017 agwlvssainokuni
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/go-yaml/yaml"
)

var (
	JsonParam bool
	YamlParam bool
)

func Parse(paramFile string) (interface{}, error) {
	if JsonParam {
		return ParseJson(paramFile)
	} else if YamlParam {
		return ParseYaml(paramFile)
	} else {
		return ParseProps(paramFile)
	}
}

func ParseProps(paramFile string) (interface{}, error) {

	data, err := ioutil.ReadFile(paramFile)
	if err != nil {
		return nil, err
	}

	param := make(map[string]interface{})
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {

		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue
		}

		index := strings.Index(line, "=")
		if index < 0 {
			continue
		}

		param[line[:index]] = line[index+1:]
	}

	return param, nil
}

func ParseJson(paramFile string) (interface{}, error) {

	data, err := ioutil.ReadFile(paramFile)
	if err != nil {
		return nil, err
	}

	var param map[string]interface{}
	err = json.Unmarshal(data, &param)
	if err != nil {
		return nil, err
	}

	return param, nil
}

func ParseYaml(paramFile string) (interface{}, error) {

	data, err := ioutil.ReadFile(paramFile)
	if err != nil {
		return nil, err
	}

	var param map[string]interface{}
	err = yaml.Unmarshal(data, &param)
	if err != nil {
		return nil, err
	}

	return param, nil
}
