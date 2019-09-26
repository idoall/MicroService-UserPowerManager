package common

import (
	"bytes"
	"fmt"
)

// const vars related to the app version
const (
	MajorVersion = "0"
	MinorVersion = "1"

	PrereleaseBlurb = "This version is pre-release, use at your own risk."
	IsRelease       = false
	Copyright       = "Copyright (c) 2019 The mshk.top Developers."
	GitHub          = "GitHub: https://github.com/idoall/MicroService-UserPowerManager"
	Issues          = "Issues: https://github.com/idoall/MicroService-UserPowerManager/issues"
)

// The following fields are populated at build time using -ldflags -X.
// Note that DATE is omitted for reproducible builds
var (
	buildGitCommit     = "unknown"
	buildGitRevision   = "unknown"
	buildUser          = "unknown"
	buildHost          = "unknown"
	buildStatus        = "unknown"
	buildTime          = "unknown"
	buildGolangVersion = "unknown"
)

// buildVersion returns the version string
func buildVersion(short bool) string {
	var b bytes.Buffer
	b.WriteString("MicroService-UserPowerManager v" + MajorVersion + "." + MinorVersion)
	// versionStr := fmt.Sprintf("MicroService-UserPowerManager v%s.%s",
	// 	MajorVersion, MinorVersion)
	if !IsRelease {
		b.WriteString(" pre-release.\n")
		if !short {
			b.WriteString(PrereleaseBlurb + "\n")
		}
	} else {
		b.WriteString(" release.\n")
	}
	if short {
		return b.String()
	}
	b.WriteString(Copyright + "\n")
	b.WriteString(fmt.Sprintf(`Git Commit: %v
Git Revision: %v
Golang Build Version: %v
Build User: %v@%v
Build Status: %v
Build Time: %v
`,
		buildGitCommit,
		buildGitRevision,
		buildGolangVersion,
		buildUser,
		buildHost,
		buildStatus,
		buildTime))
	b.WriteString(GitHub + "\n")
	b.WriteString(Issues + "\n")
	return b.String()
}

// OutVersion 输出版本
func OutVersion() {
	fmt.Println(buildVersion(false))
}
