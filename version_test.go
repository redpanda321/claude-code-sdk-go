package claudecodesdk

import (
	"regexp"
	"testing"
)

// semverPreRelease matches X.Y.Z-prerelease per https://semver.org section 9.
var semverPreRelease = regexp.MustCompile(`^\d+\.\d+\.\d+-[0-9A-Za-z.-]+$`)

func TestVersion_IsSemverPrerelease(t *testing.T) {
	if !semverPreRelease.MatchString(Version) {
		t.Fatalf("Version = %q, want match of semver pre-release regex", Version)
	}
}

func TestVersion_IsAlpha(t *testing.T) {
	if Version == "" || Version[0] == 'v' {
		t.Fatalf("Version = %q, want non-empty without leading 'v'", Version)
	}
}
