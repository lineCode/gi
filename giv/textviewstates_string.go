// Code generated by "stringer -type=TextViewStates"; DO NOT EDIT.

package giv

import (
	"fmt"
	"strconv"
)

const _TextViewStates_name = "TextViewActiveTextViewFocusTextViewInactiveTextViewSelTextViewHighlightTextViewStatesN"

var _TextViewStates_index = [...]uint8{0, 14, 27, 43, 54, 71, 86}

func (i TextViewStates) String() string {
	if i < 0 || i >= TextViewStates(len(_TextViewStates_index)-1) {
		return "TextViewStates(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TextViewStates_name[_TextViewStates_index[i]:_TextViewStates_index[i+1]]
}

func (i *TextViewStates) FromString(s string) error {
	for j := 0; j < len(_TextViewStates_index)-1; j++ {
		if s == _TextViewStates_name[_TextViewStates_index[j]:_TextViewStates_index[j+1]] {
			*i = TextViewStates(j)
			return nil
		}
	}
	return fmt.Errorf("String %v is not a valid option for type TextViewStates", s)
}
