// Code generated by "stringer -type=tokenKind -trimprefix=tok -linecomment=true"; DO NOT EDIT.

package syntax

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[tokNone-0]
	_ = x[tokChar-1]
	_ = x[tokGroupFlags-2]
	_ = x[tokPosixClass-3]
	_ = x[tokConcat-4]
	_ = x[tokRepeat-5]
	_ = x[tokEscapeChar-6]
	_ = x[tokEscapeMeta-7]
	_ = x[tokEscapeOctal-8]
	_ = x[tokEscapeUni-9]
	_ = x[tokEscapeUniFull-10]
	_ = x[tokEscapeHex-11]
	_ = x[tokEscapeHexFull-12]
	_ = x[tokComment-13]
	_ = x[tokQ-14]
	_ = x[tokMinus-15]
	_ = x[tokLbracket-16]
	_ = x[tokLbracketCaret-17]
	_ = x[tokRbracket-18]
	_ = x[tokDollar-19]
	_ = x[tokCaret-20]
	_ = x[tokQuestion-21]
	_ = x[tokDot-22]
	_ = x[tokPlus-23]
	_ = x[tokStar-24]
	_ = x[tokPipe-25]
	_ = x[tokLparen-26]
	_ = x[tokLparenName-27]
	_ = x[tokLparenNameAngle-28]
	_ = x[tokLparenNameQuote-29]
	_ = x[tokLparenFlags-30]
	_ = x[tokLparenAtomic-31]
	_ = x[tokLparenPositiveLookahead-32]
	_ = x[tokLparenPositiveLookbehind-33]
	_ = x[tokLparenNegativeLookahead-34]
	_ = x[tokLparenNegativeLookbehind-35]
	_ = x[tokRparen-36]
}

const _tokenKind_name = "NoneCharGroupFlagsPosixClassConcatRepeatEscapeCharEscapeMetaEscapeOctalEscapeUniEscapeUniFullEscapeHexEscapeHexFullComment\\Q-[[^]$^?.+*|((?P<name>(?<name>(?'name'(?flags(?>(?=(?<=(?!(?<!)"

var _tokenKind_index = [...]uint8{0, 4, 8, 18, 28, 34, 40, 50, 60, 71, 80, 93, 102, 115, 122, 124, 125, 126, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 146, 154, 162, 169, 172, 175, 179, 182, 186, 187}

func (i tokenKind) String() string {
	if i >= tokenKind(len(_tokenKind_index)-1) {
		return "tokenKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _tokenKind_name[_tokenKind_index[i]:_tokenKind_index[i+1]]
}