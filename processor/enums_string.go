// Code generated by "stringer -linecomment -type OutputFmt,NotesFmt,TOCPlacement,TOCType,APNXGeneration,StampPlacement,CoverProcessing -output processor/enums_string.go processor/enums.go"; DO NOT EDIT.

package processor

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OEpub-0]
	_ = x[OKepub-1]
	_ = x[OAzw3-2]
	_ = x[OMobi-3]
	_ = x[UnsupportedOutputFmt-4]
}

const _OutputFmt_name = "epubkepubazw3mobi"

var _OutputFmt_index = [...]uint8{0, 4, 9, 13, 17, 17}

func (i OutputFmt) String() string {
	if i < 0 || i >= OutputFmt(len(_OutputFmt_index)-1) {
		return "OutputFmt(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OutputFmt_name[_OutputFmt_index[i]:_OutputFmt_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NDefault-0]
	_ = x[NInline-1]
	_ = x[NBlock-2]
	_ = x[NFloat-3]
	_ = x[NFloatOld-4]
	_ = x[NFloatNew-5]
	_ = x[UnsupportedNotesFmt-6]
}

const _NotesFmt_name = "defaultinlineblockfloatfloat-oldfloat-new"

var _NotesFmt_index = [...]uint8{0, 7, 13, 18, 23, 32, 41, 41}

func (i NotesFmt) String() string {
	if i < 0 || i >= NotesFmt(len(_NotesFmt_index)-1) {
		return "NotesFmt(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NotesFmt_name[_NotesFmt_index[i]:_NotesFmt_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TOCNone-0]
	_ = x[TOCBefore-1]
	_ = x[TOCAfter-2]
	_ = x[UnsupportedTOCPlacement-3]
}

const _TOCPlacement_name = "nonebeforeafter"

var _TOCPlacement_index = [...]uint8{0, 4, 10, 15, 15}

func (i TOCPlacement) String() string {
	if i < 0 || i >= TOCPlacement(len(_TOCPlacement_index)-1) {
		return "TOCPlacement(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TOCPlacement_name[_TOCPlacement_index[i]:_TOCPlacement_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TOCTypeNormal-0]
	_ = x[TOCTypeKindle-1]
	_ = x[TOCTypeFlat-2]
	_ = x[UnsupportedTOCType-3]
}

const _TOCType_name = "normalkindleflat"

var _TOCType_index = [...]uint8{0, 6, 12, 16, 16}

func (i TOCType) String() string {
	if i < 0 || i >= TOCType(len(_TOCType_index)-1) {
		return "TOCType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TOCType_name[_TOCType_index[i]:_TOCType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[APNXNone-0]
	_ = x[APNXEInk-1]
	_ = x[APNXApp-2]
	_ = x[UnsupportedAPNXGeneration-3]
}

const _APNXGeneration_name = "noneeinkapp"

var _APNXGeneration_index = [...]uint8{0, 4, 8, 11, 11}

func (i APNXGeneration) String() string {
	if i < 0 || i >= APNXGeneration(len(_APNXGeneration_index)-1) {
		return "APNXGeneration(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _APNXGeneration_name[_APNXGeneration_index[i]:_APNXGeneration_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StampNone-0]
	_ = x[StampTop-1]
	_ = x[StampMiddle-2]
	_ = x[StampBottom-3]
	_ = x[UnsupportedStampPlacement-4]
}

const _StampPlacement_name = "nonetopmiddlebottom"

var _StampPlacement_index = [...]uint8{0, 4, 7, 13, 19, 19}

func (i StampPlacement) String() string {
	if i < 0 || i >= StampPlacement(len(_StampPlacement_index)-1) {
		return "StampPlacement(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StampPlacement_name[_StampPlacement_index[i]:_StampPlacement_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CoverNone-0]
	_ = x[CoverKeepAR-1]
	_ = x[CoverStretch-2]
	_ = x[UnsupportedCoverProcessing-3]
}

const _CoverProcessing_name = "nonekeepARstretch"

var _CoverProcessing_index = [...]uint8{0, 4, 10, 17, 17}

func (i CoverProcessing) String() string {
	if i < 0 || i >= CoverProcessing(len(_CoverProcessing_index)-1) {
		return "CoverProcessing(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CoverProcessing_name[_CoverProcessing_index[i]:_CoverProcessing_index[i+1]]
}
