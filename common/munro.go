// Copyright 2016 Josh Deprez
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

package common

import "github.com/DrJosh9000/awakengine"

const (
	munroHeight  = 11
	munroYOffset = -2
)

var (
	munroMap = map[byte]awakengine.CharInfo{
		' ':  {Width: 0, X: 1, Y: 9, XOffset: 0, Height: 0, YOffset: 11, XAdvance: 3},
		'!':  {Width: 1, X: 2, Y: 2, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 2},
		'"':  {Width: 3, X: 4, Y: 1, XOffset: 0, Height: 2, YOffset: 3, XAdvance: 4},
		'#':  {Width: 5, X: 8, Y: 3, XOffset: 0, Height: 6, YOffset: 5, XAdvance: 6},
		'$':  {Width: 3, X: 14, Y: 1, XOffset: 0, Height: 9, YOffset: 3, XAdvance: 4},
		'%':  {Width: 7, X: 18, Y: 2, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 8},
		'&':  {Width: 6, X: 26, Y: 2, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 7},
		'\'': {Width: 1, X: 33, Y: 1, XOffset: 0, Height: 2, YOffset: 3, XAdvance: 2},
		'(':  {Width: 2, X: 35, Y: 1, XOffset: 0, Height: 9, YOffset: 3, XAdvance: 3},
		')':  {Width: 2, X: 38, Y: 1, XOffset: 0, Height: 9, YOffset: 3, XAdvance: 3},
		'*':  {Width: 3, X: 41, Y: 1, XOffset: 0, Height: 3, YOffset: 3, XAdvance: 4},
		'+':  {Width: 5, X: 45, Y: 3, XOffset: 0, Height: 5, YOffset: 5, XAdvance: 6},
		',':  {Width: 1, X: 51, Y: 8, XOffset: 0, Height: 2, YOffset: 10, XAdvance: 2},
		'-':  {Width: 3, X: 53, Y: 6, XOffset: 0, Height: 1, YOffset: 8, XAdvance: 4},
		'.':  {Width: 1, X: 57, Y: 8, XOffset: 0, Height: 1, YOffset: 10, XAdvance: 2},
		'/':  {Width: 3, X: 59, Y: 1, XOffset: 0, Height: 9, YOffset: 3, XAdvance: 4},
		'0':  {Width: 4, X: 63, Y: 2, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'1':  {Width: 2, X: 1, Y: 11, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 3},
		'2':  {Width: 3, X: 4, Y: 11, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 4},
		'3':  {Width: 3, X: 8, Y: 11, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 4},
		'4':  {Width: 4, X: 12, Y: 11, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'5':  {Width: 3, X: 17, Y: 11, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 4},
		'6':  {Width: 4, X: 21, Y: 11, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'7':  {Width: 3, X: 26, Y: 11, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 4},
		'8':  {Width: 4, X: 30, Y: 11, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'9':  {Width: 4, X: 35, Y: 11, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		':':  {Width: 1, X: 40, Y: 14, XOffset: 0, Height: 4, YOffset: 7, XAdvance: 2},
		';':  {Width: 1, X: 42, Y: 14, XOffset: 0, Height: 5, YOffset: 7, XAdvance: 2},
		'<':  {Width: 5, X: 44, Y: 12, XOffset: 0, Height: 5, YOffset: 5, XAdvance: 5},
		'=':  {Width: 5, X: 50, Y: 13, XOffset: 0, Height: 3, YOffset: 6, XAdvance: 5},
		'>':  {Width: 5, X: 56, Y: 12, XOffset: 0, Height: 5, YOffset: 5, XAdvance: 5},
		'?':  {Width: 3, X: 62, Y: 11, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 4},
		'@':  {Width: 7, X: 1, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 8},
		'A':  {Width: 4, X: 9, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'B':  {Width: 4, X: 14, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'C':  {Width: 3, X: 19, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 4},
		'D':  {Width: 4, X: 23, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'E':  {Width: 3, X: 28, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 4},
		'F':  {Width: 3, X: 32, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 4},
		'G':  {Width: 4, X: 36, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'H':  {Width: 4, X: 41, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'I':  {Width: 1, X: 46, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 2},
		'J':  {Width: 2, X: 48, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 3},
		'K':  {Width: 4, X: 51, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'L':  {Width: 3, X: 56, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 4},
		'M':  {Width: 5, X: 60, Y: 20, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 6},
		'N':  {Width: 4, X: 1, Y: 28, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'O':  {Width: 4, X: 6, Y: 28, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'P':  {Width: 4, X: 11, Y: 28, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'Q':  {Width: 4, X: 16, Y: 28, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'R':  {Width: 4, X: 21, Y: 28, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'S':  {Width: 3, X: 26, Y: 28, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 4},
		'T':  {Width: 3, X: 30, Y: 28, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 4},
		'U':  {Width: 4, X: 34, Y: 28, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'V':  {Width: 5, X: 39, Y: 28, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 6},
		'W':  {Width: 5, X: 45, Y: 28, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 6},
		'X':  {Width: 5, X: 51, Y: 28, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 6},
		'Y':  {Width: 5, X: 57, Y: 28, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 6},
		'Z':  {Width: 3, X: 63, Y: 28, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 4},
		'[':  {Width: 2, X: 1, Y: 36, XOffset: 0, Height: 11, YOffset: 2, XAdvance: 3},
		'\\': {Width: 3, X: 4, Y: 37, XOffset: 0, Height: 9, YOffset: 3, XAdvance: 4},
		']':  {Width: 2, X: 8, Y: 36, XOffset: 0, Height: 11, YOffset: 2, XAdvance: 3},
		'^':  {Width: 3, X: 11, Y: 37, XOffset: 0, Height: 2, YOffset: 3, XAdvance: 4},
		'_':  {Width: 5, X: 15, Y: 45, XOffset: 0, Height: 1, YOffset: 11, XAdvance: 6},
		'a':  {Width: 4, X: 21, Y: 40, XOffset: 0, Height: 5, YOffset: 6, XAdvance: 5},
		'b':  {Width: 4, X: 26, Y: 38, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'c':  {Width: 3, X: 31, Y: 40, XOffset: 0, Height: 5, YOffset: 6, XAdvance: 4},
		'd':  {Width: 4, X: 35, Y: 38, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'e':  {Width: 4, X: 40, Y: 40, XOffset: 0, Height: 5, YOffset: 6, XAdvance: 5},
		'f':  {Width: 3, X: 45, Y: 38, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 3},
		'g':  {Width: 4, X: 49, Y: 40, XOffset: 0, Height: 7, YOffset: 6, XAdvance: 5},
		'h':  {Width: 4, X: 54, Y: 38, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'i':  {Width: 1, X: 59, Y: 38, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 2},
		'j':  {Width: 2, X: 61, Y: 38, XOffset: -1, Height: 9, YOffset: 4, XAdvance: 2},
		'k':  {Width: 4, X: 64, Y: 38, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 5},
		'l':  {Width: 1, X: 1, Y: 48, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 2},
		'm':  {Width: 7, X: 3, Y: 50, XOffset: 0, Height: 5, YOffset: 6, XAdvance: 8},
		'n':  {Width: 4, X: 11, Y: 50, XOffset: 0, Height: 5, YOffset: 6, XAdvance: 5},
		'o':  {Width: 4, X: 16, Y: 50, XOffset: 0, Height: 5, YOffset: 6, XAdvance: 5},
		'p':  {Width: 4, X: 21, Y: 50, XOffset: 0, Height: 7, YOffset: 6, XAdvance: 5},
		'q':  {Width: 5, X: 26, Y: 50, XOffset: 0, Height: 7, YOffset: 6, XAdvance: 5},
		'r':  {Width: 3, X: 32, Y: 50, XOffset: 0, Height: 5, YOffset: 6, XAdvance: 4},
		's':  {Width: 3, X: 36, Y: 50, XOffset: 0, Height: 5, YOffset: 6, XAdvance: 4},
		't':  {Width: 3, X: 40, Y: 48, XOffset: 0, Height: 7, YOffset: 4, XAdvance: 4},
		'u':  {Width: 4, X: 44, Y: 50, XOffset: 0, Height: 5, YOffset: 6, XAdvance: 5},
		'v':  {Width: 5, X: 49, Y: 50, XOffset: 0, Height: 5, YOffset: 6, XAdvance: 6},
		'w':  {Width: 7, X: 55, Y: 50, XOffset: 0, Height: 5, YOffset: 6, XAdvance: 8},
		'x':  {Width: 5, X: 63, Y: 50, XOffset: 0, Height: 5, YOffset: 6, XAdvance: 6},
		'y':  {Width: 4, X: 1, Y: 62, XOffset: 0, Height: 7, YOffset: 6, XAdvance: 5},
		'z':  {Width: 3, X: 6, Y: 62, XOffset: 0, Height: 5, YOffset: 6, XAdvance: 4},
		'{':  {Width: 3, X: 10, Y: 59, XOffset: 0, Height: 9, YOffset: 3, XAdvance: 4},
		'|':  {Width: 1, X: 14, Y: 58, XOffset: 1, Height: 11, YOffset: 2, XAdvance: 4},
		'}':  {Width: 3, X: 16, Y: 59, XOffset: 0, Height: 9, YOffset: 3, XAdvance: 4},
		'~':  {Width: 5, X: 20, Y: 63, XOffset: 0, Height: 3, YOffset: 7, XAdvance: 6},
	}
)

// MunroFont implements awakengine.Font with the Munro typeface.
type MunroFont struct {
	Invert bool
}

// Source implements awkaengine.Font.
func (m MunroFont) ImageKey(invert bool) string {
	if m.Invert != invert {
		return "inv_munro"
	}
	return "munro"
}

// Metrics implements awakengine.Font.
func (MunroFont) Metrics() awakengine.CharMetrics { return munroMap }

// LineHeight implements awakengine.Font.
func (MunroFont) LineHeight() int { return munroHeight }

// YOffset implements awakengine.Font.
func (MunroFont) YOffset() int { return munroYOffset }
