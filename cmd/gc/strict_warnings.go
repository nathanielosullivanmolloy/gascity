package main

import "strings"

// splitStrictConfigWarnings separates warnings that should remain fatal in
// strict mode from compatibility/migration guidance that should stay warnings.
func splitStrictConfigWarnings(warnings []string) (fatal []string, nonFatal []string) {
	for _, warning := range warnings {
		if strictWarningIsNonFatal(warning) {
			nonFatal = append(nonFatal, warning)
			continue
		}
		fatal = append(fatal, warning)
	}
	return fatal, nonFatal
}

func strictWarningIsNonFatal(warning string) bool {
	// Only migration guidance that preserves a usable config stays non-fatal
	// in strict mode. Missing rig bindings still remain fatal.
	return strings.Contains(warning, "still declares path in city.toml; move it to .gc/site.toml") ||
		strings.HasPrefix(warning, ".gc/site.toml declares a binding for unknown rig ")
}
