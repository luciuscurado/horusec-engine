// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package javascript_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ZupIT/horusec-engine/internal/ast"
	javascript "github.com/ZupIT/horusec-engine/internal/horusec-javascript"
)

func TestJavaScriptParseFileFillAllPositions(t *testing.T) {
	src := []byte(`
class Foo { f(a, b) { return a + b }}

function f1(s) { console.log(s) } 

const f2 (a, b) => { return a / b }
	`)

	f, err := javascript.ParseFile("", src)
	require.NoError(t, err, "Expected no error to parse source file: %v", err)

	notExpectedPos := ast.Pos{
		Byte:   0,
		Row:    0,
		Column: 0,
	}

	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		start := f.Start()
		end := f.End()

		assert.NotEqual(t, notExpectedPos, start, "Expected not empty start position from node %T: %s", n, start)
		assert.NotEqual(t, notExpectedPos, end, "Expected not empty end position from node %T: %s", n, end)

		return true
	})
}
