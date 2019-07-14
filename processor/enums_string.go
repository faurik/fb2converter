// Code generated by "stringer -linecomment -type OutputFmt,NotesFmt,TOCPlacement,TOCType,APNXGeneration,StampPlacement -output /mnt/e/projects/books/fb2converter/processor/enums_string.go /mnt/e/projects/books/fb2converter/processor/enums.go"; DO NOT EDIT.

package processor

import "strconv"

const _OutputFmt_name = "epubkepubazw3mobi"

var _OutputFmt_index = [...]uint8{0, 4, 9, 13, 17, 17}

func (i OutputFmt) String() string {
	if i < 0 || i >= OutputFmt(len(_OutputFmt_index)-1) {
		return "OutputFmt(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OutputFmt_name[_OutputFmt_index[i]:_OutputFmt_index[i+1]]
}

const _NotesFmt_name = "defaultinlineblockfloatfloat-old"

var _NotesFmt_index = [...]uint8{0, 7, 13, 18, 23, 32, 32}

func (i NotesFmt) String() string {
	if i < 0 || i >= NotesFmt(len(_NotesFmt_index)-1) {
		return "NotesFmt(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NotesFmt_name[_NotesFmt_index[i]:_NotesFmt_index[i+1]]
}

const _TOCPlacement_name = "nonebeforeafter"

var _TOCPlacement_index = [...]uint8{0, 4, 10, 15, 15}

func (i TOCPlacement) String() string {
	if i < 0 || i >= TOCPlacement(len(_TOCPlacement_index)-1) {
		return "TOCPlacement(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TOCPlacement_name[_TOCPlacement_index[i]:_TOCPlacement_index[i+1]]
}

const _TOCType_name = "normalkindleflat"

var _TOCType_index = [...]uint8{0, 6, 12, 16, 16}

func (i TOCType) String() string {
	if i < 0 || i >= TOCType(len(_TOCType_index)-1) {
		return "TOCType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TOCType_name[_TOCType_index[i]:_TOCType_index[i+1]]
}

const _APNXGeneration_name = "noneeinkapp"

var _APNXGeneration_index = [...]uint8{0, 4, 8, 11, 11}

func (i APNXGeneration) String() string {
	if i < 0 || i >= APNXGeneration(len(_APNXGeneration_index)-1) {
		return "APNXGeneration(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _APNXGeneration_name[_APNXGeneration_index[i]:_APNXGeneration_index[i+1]]
}

const _StampPlacement_name = "nonetopmiddlebottom"

var _StampPlacement_index = [...]uint8{0, 4, 7, 13, 19, 19}

func (i StampPlacement) String() string {
	if i < 0 || i >= StampPlacement(len(_StampPlacement_index)-1) {
		return "StampPlacement(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StampPlacement_name[_StampPlacement_index[i]:_StampPlacement_index[i+1]]
}
