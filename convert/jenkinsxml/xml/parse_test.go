// Copyright 2024 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xml

import (
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseFile(t *testing.T) {
	got, err := ParseFile("testdata/hello.xml")
	if err != nil {
		t.Error(err)
	}

	want := &Project{
		Disabled: false,
		Builders: &Builders{
			Tasks: []Task{
				{
					XMLName: xml.Name{
						Local: "hudson.tasks.Shell",
					},
					Content: `
      <command>echo hello</command>
      <configuredLocalRules/>
    `,
				},
			},
		},
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected parsing result")
		t.Log(diff)
	}

	got, err = ParseFile("testdata/shell-and-ant.xml")
	if err != nil {
		t.Error(err)
	}

	want = &Project{
		Disabled: false,
		Builders: &Builders{
			Tasks: []Task{
				{
					XMLName: xml.Name{
						Local: "hudson.tasks.Shell",
					},
					Content: `
      <command>echo hello</command>
      <configuredLocalRules/>
    `,
				},
				{
					XMLName: xml.Name{
						Local: "hudson.tasks.Ant",
					},
					Content: `
      <targets>one/two/three</targets>
    `,
				},
				{
					XMLName: xml.Name{
						Local: "hudson.tasks.Shell",
					},
					Content: `
      <command>echo hello again</command>
      <configuredLocalRules/>
    `,
				},
			},
		},
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected parsing result")
		t.Log(diff)
	}
}

func TestParseFile_Error(t *testing.T) {
	_, err := ParseFile("testdata/file-does-not-exist.yaml")
	if err == nil {
		t.Errorf("Expect parsing error")
	}
}

func TestParseString(t *testing.T) {
	out, _ := ioutil.ReadFile("testdata/hello.xml")
	got, err := ParseString(string(out))
	if err != nil {
		t.Error(err)
	}

	want := &Project{
		Disabled: false,
		Builders: &Builders{
			Tasks: []Task{
				{
					XMLName: xml.Name{
						Local: "hudson.tasks.Shell",
					},
					Content: `
      <command>echo hello</command>
      <configuredLocalRules/>
    `,
				},
			},
		},
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected parsing result")
		t.Log(diff)
	}
}
