package handlers

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/apparentlymart/go-versions/versions"
)

// ProcessVersions filters and selects the newest version based on the given constraint.
func ProcessVersions(ctx context.Context, versionList string, constraint string) (string, error) {
	versionStrings := strings.Split(versionList, ",")
	var available versions.List
	for _, v := range versionStrings {
		parsedVersion, err := versions.ParseVersion(strings.TrimSpace(v))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid version format: %s\n %+v", v, err)
			return "", err
		}
		available = append(available, parsedVersion)
	}

	// Parse the version constraint
	allowed, err := versions.MeetingConstraintsString(constraint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid version constraint: %s\n", err)
		return "", err
	}

	// Filter versions that meet the constraint
	candidates := available.Filter(allowed)
	if len(candidates) == 0 {
		fmt.Fprintf(os.Stderr, "No versions match the given constraint")
		return "", nil
	}

	// Select the newest matching version
	chosen := candidates.Newest()
	return chosen.String(), nil
}
