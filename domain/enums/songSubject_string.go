// Code generated by "stringer -type=SongSubject -trimprefix=Subject -output=songSubject_string.go"; DO NOT EDIT.

package enums

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SongSubjectRomantic-1]
	_ = x[SongSubjectSocial-2]
	_ = x[SongSubjectEpic-3]
	_ = x[SongSubjectReligious-4]
}

const _SongSubject_name = "SongSubjectRomanticSongSubjectSocialSongSubjectEpicSongSubjectReligious"

var _SongSubject_index = [...]uint8{0, 19, 36, 51, 71}

func (i SongSubject) String() string {
	i -= 1
	if i >= SongSubject(len(_SongSubject_index)-1) {
		return "SongSubject(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _SongSubject_name[_SongSubject_index[i]:_SongSubject_index[i+1]]
}