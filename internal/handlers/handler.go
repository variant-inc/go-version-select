package handlers

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/Masterminds/semver/v3"
)

// ProcessVersions filters and selects the newest version based on the given constraint.
func ProcessVersions(_ context.Context, versionList, constraint string) (string, error) {
	// Split the comma-separated version list
	versionStrings := strings.Split(versionList, ",")

	// Parse all versions into semver.Version objects
	var available []*semver.Version
	for _, v := range versionStrings {
		parsedVersion, err := semver.NewVersion(strings.TrimSpace(v))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid version format: %s\n %+v", v, err)
			return "", err
		}
		available = append(available, parsedVersion)
	}

	// Parse the version constraint
	allowed, err := semver.NewConstraint(constraint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid version constraint: %s\n", err)
		return "", err
	}

	// Filter versions that meet the constraint
	var candidates []*semver.Version
	for _, v := range available {
		if allowed.Check(v) {
			candidates = append(candidates, v)
		}
	}

	// Handle case where no versions match the constraint
	if len(candidates) == 0 {
		fmt.Fprintf(os.Stderr, "No versions match the given constraint")
		return "", nil
	}

	// Sort matching versions and select the newest one
	sort.Sort(semver.Collection(candidates))
	chosen := candidates[len(candidates)-1]

	return chosen.String(), nil
}
