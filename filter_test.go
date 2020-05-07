	"github.com/reviewdog/reviewdog/difffilter"
const diffContentAddedStrip = `diff --git a/test_added.go b/test_added.go
new file mode 100644
index 0000000..264c67e
--- /dev/null
+++ b/test_added.go
@@ -0,0 +1,3 @@
+package reviewdog
+
+var TestAdded = 14
`

			ShouldReport: false,
			InDiffFile:   true,
			OldPath:      "sample.old.txt",
			OldLine:      1,
			ShouldReport: true,
			InDiffFile:   true,
			LnumDiff:     3,
			OldPath:      "sample.old.txt",
			OldLine:      0,
			ShouldReport: false,
			InDiffFile:   true,
			OldPath:      "nonewline.old.txt",
			OldLine:      1,
			ShouldReport: true,
			InDiffFile:   true,
			LnumDiff:     5,
			OldPath:      "nonewline.old.txt",
			OldLine:      0,
	got := FilterCheck(results, filediffs, 0, "", difffilter.ModeAdded)
			ShouldReport: true,
			InDiffFile:   true,
			LnumDiff:     1,
			OldPath:      "sample.old.txt",
			OldLine:      1,
			ShouldReport: true,
			InDiffFile:   true,
			LnumDiff:     3,
			OldPath:      "sample.old.txt",
			OldLine:      0,
			ShouldReport: true,
			InDiffFile:   true,
			LnumDiff:     4,
			OldPath:      "sample.old.txt",
			OldLine:      0,
	got := FilterCheck(results, filediffs, 0, "", difffilter.ModeDiffContext)
func findFileDiff(filediffs []*diff.FileDiff, path string, strip int) *diff.FileDiff {
	for _, file := range filediffs {
		if difffilter.NormalizeDiffPath(file.PathNew, strip) == path {
			return file
		}
	}
	return nil
}

func TestGetOldPosition(t *testing.T) {
	const strip = 0
	tests := []struct {
		newPath     string
		newLine     int
		wantOldPath string
		wantOldLine int
	}{
		{
			newPath:     "sample.new.txt",
			newLine:     1,
			wantOldPath: "sample.old.txt",
			wantOldLine: 1,
		},
		{
			newPath:     "sample.new.txt",
			newLine:     2,
			wantOldPath: "sample.old.txt",
			wantOldLine: 0,
		},
		{
			newPath:     "sample.new.txt",
			newLine:     3,
			wantOldPath: "sample.old.txt",
			wantOldLine: 0,
		},
		{
			newPath:     "sample.new.txt",
			newLine:     14,
			wantOldPath: "sample.old.txt",
			wantOldLine: 13,
		},
		{
			newPath:     "not_found",
			newLine:     14,
			wantOldPath: "",
			wantOldLine: 0,
		},
	for _, tt := range tests {
		fdiff := findFileDiff(filediffs, tt.newPath, strip)
		gotPath, gotLine := getOldPosition(fdiff, strip, tt.newPath, tt.newLine)
		if !(gotPath == tt.wantOldPath && gotLine == tt.wantOldLine) {
			t.Errorf("getOldPosition(..., %s, %d) = (%s, %d), want (%s, %d)",
				tt.newPath, tt.newLine, gotPath, gotLine, tt.wantOldPath, tt.wantOldLine)
}

func TestGetOldPosition_added(t *testing.T) {
	const strip = 1
	filediffs, _ := diff.ParseMultiFile(strings.NewReader(diffContentAddedStrip))
	path := "test_added.go"
	fdiff := findFileDiff(filediffs, path, strip)
	gotPath, _ := getOldPosition(fdiff, strip, path, 1)
	if gotPath != "" {
		t.Errorf("got %q as old path for addedd diff file, want empty", gotPath)