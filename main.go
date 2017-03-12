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
	"flag"
	"fmt"
	"os"
)

func main() {

	var (
		paramFile  string
		targetFile string
		jsonParam  bool
		yamlParam  bool
	)

	flag.StringVar(&paramFile, "p", "", "置換文字列定義ファイルのパス")
	flag.StringVar(&targetFile, "o", "", "生成するファイルのパス")
	flag.BoolVar(&jsonParam, "j", false, "JSON形式の置換文字列定義ファイルを使用")
	flag.BoolVar(&yamlParam, "y", false, "YAML形式の置換文字列定義ファイルを使用")
	flag.Parse()

	var parse func(string) (interface{}, error)
	if jsonParam {
		parse = ParseJson
	} else if yamlParam {
		parse = ParseYaml
	} else {
		parse = ParseProps
	}

	param, err := parse(paramFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "置換文字列定義ファイル読み込みエラー: %s\n", err.Error())
		os.Exit(-1)
	}

	err = render(flag.Args(), targetFile, param)
	if err != nil {
		fmt.Fprintf(os.Stderr, "生成処理エラー: %s\n", err.Error())
		os.Exit(-1)
	}
}

func render(template []string, target string, context ...interface{}) error {
	if len(template) <= 0 {
		if len(target) <= 0 {
			return RenderStream2Stream(os.Stdin, os.Stdout, context...)
		} else {
			return RenderStream2File(os.Stdin, target, context...)
		}
	} else {
		if len(target) <= 0 {
			return RenderFiles2Stream(template, os.Stdout, context...)
		} else {
			return RenderFiles2File(template, target, context...)
		}
	}
}
