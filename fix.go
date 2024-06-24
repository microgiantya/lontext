package lontext

func fixPrefix(prefix string) (fixedPrefix string) {
	fixedPrefix = prefix
	if fixedPrefix == "" {
		fixedPrefix = "unknown"
	}
	return
}

func fixVersion(version string) (fixedVersion string) {
	fixedVersion = version
	if fixedVersion == "" {
		fixedVersion = "v-"
	}
	return
}

func fixView(view view) (fixedView view) {
	fixedView = view
	if fixedView != lontextViewPlain && fixedView != lontextViewJSON {
		fixedView = lontextViewPlain
	}
	return
}
