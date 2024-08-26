package catalog

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/khulnasoft/meshkit/models/catalog/v1alpha1"
)

func BuildArtifactHubPkg(name, downloadURL, user, version, createdAt string, catalogData *v1alpha1.CatalogData) *ArtifactHubMetadata {
	artifacthubPkg := &ArtifactHubMetadata{
		Name:        toKebabCase(name),
		DisplayName: name,
		Description: valueOrElse(catalogData.PatternInfo, "A Meshplay Design"),
		Provider: Provider{
			Name: user,
		},
		Links: []Link{
			{
				Name: "download",
				URL:  downloadURL, // this depends on where the design is stored by the user, we can give remote provider URL otherwise
			},
			{
				Name: "Meshplay Catalog",
				URL:  "https://meshplay.io/catalog",
			},
		},
		HomeURL:   "https://docs.meshplay.io/concepts/logical/designs",
		Version:   valueOrElse(version, "0.0.1"),
		CreatedAt: createdAt,
		License:   "Apache-2.0",
		LogoURL:   "https://raw.githubusercontent.com/meshplay/meshplay.io/0b8585231c6e2b3251d38f749259360491c9ee6b/assets/images/brand/meshplay-logo.svg",
		Install:   "meshplayctl design import -f",
		Readme:    fmt.Sprintf("%s \n ##h4 Caveats and Consideration \n", catalogData.PatternCaveats),
	}

	if len(catalogData.SnapshotURL) > 0 {
		artifacthubPkg.Screenshots = append(artifacthubPkg.Screenshots, Screenshot{
			Title: "MeshMap Snapshot",
			URL:   catalogData.SnapshotURL[0],
		})

		if len(catalogData.SnapshotURL) > 1 {
			artifacthubPkg.Screenshots = append(artifacthubPkg.Screenshots, Screenshot{
				Title: "MeshMap Snapshot",
				URL:   catalogData.SnapshotURL[1],
			})
		}
	}

	artifacthubPkg.Screenshots = append(artifacthubPkg.Screenshots, Screenshot{
		Title: "Meshplay Project",
		URL:   "https://raw.githubusercontent.com/meshplay/meshplay.io/master/assets/images/logos/meshplay-gradient.png",
	})

	return artifacthubPkg
}

func valueOrElse(s string, fallback string) string {
	if len(s) == 0 {
		return fallback
	} else {
		return s
	}
}

func toKebabCase(s string) string {
	s = strings.ToLower(s)
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")
	s = strings.ReplaceAll(s, " ", "-")

	return s
}
