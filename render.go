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
	"io"
	"io/ioutil"
	"os"

	"github.com/cbroglie/mustache"
)

func RenderStream2Stream(in io.Reader, out io.Writer, context ...interface{}) error {
	data, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	template, err := mustache.ParseString(string(data))
	if err != nil {
		return err
	}
	return template.FRender(out, context...)
}

func RenderStream2File(in io.Reader, target string, context ...interface{}) error {
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	err = RenderStream2Stream(in, out, context...)
	if err2 := out.Close(); err2 != nil && err == nil {
		return err2
	}
	return err
}

func RenderFiles2Stream(template []string, out io.Writer, context ...interface{}) error {
	for _, file := range template {
		in, err := os.Open(file)
		if err != nil {
			return err
		}
		err = RenderStream2Stream(in, out, context...)
		if err2 := in.Close(); err2 != nil && err == nil {
			return err2
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func RenderFiles2File(template []string, target string, context ...interface{}) error {
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	err = RenderFiles2Stream(template, out, context...)
	if err2 := out.Close(); err2 != nil && err == nil {
		return err2
	}
	return err
}
