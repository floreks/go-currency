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

package common

import "testing"

const EPSILON = 0.00000001

func TestRound(t *testing.T) {
	equals := func(a, b float64) bool {
		if (a-b) < EPSILON && (b-a) < EPSILON {
			return true
		}
		return false
	}

	cases := []struct {
		value    float64
		places   int
		expected float64
	}{
		{1.23456, 0, 1},
		{1.23456, 1, 1.2},
		{1.23456, 2, 1.23},
		{1.23456, 3, 1.235},
	}

	for _, c := range cases {
		actual := Round(c.value, c.places)

		if !equals(actual, c.expected) {
			t.Errorf("Round(%f, %d) == \ngot: %f, \nexpected %f", c.value, c.places,
				actual, c.expected)
		}
	}
}
