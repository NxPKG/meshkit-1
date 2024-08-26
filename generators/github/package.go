package github

import (
	"bytes"
	"os"

	"github.com/khulnasoft/meshkit/utils"
	"github.com/khulnasoft/meshkit/utils/component"
	"github.com/khulnasoft/meshkit/utils/manifests"
	"github.com/meshplay/schemas/models/v1beta1/category"
	_component "github.com/meshplay/schemas/models/v1beta1/component"
	"github.com/meshplay/schemas/models/v1beta1/model"
)

type GitHubPackage struct {
	Name       string `yaml:"name" json:"name"`
	filePath   string
	branch     string
	repository string
	version    string
	SourceURL  string `yaml:"source_url" json:"source_url"`
}

func (gp GitHubPackage) GetVersion() string {
	return gp.version
}

func (gp GitHubPackage) GenerateComponents() ([]_component.ComponentDefinition, error) {
	components := make([]_component.ComponentDefinition, 0)

	data, err := os.ReadFile(gp.filePath)
	if err != nil {
		return nil, ErrGenerateGitHubPackage(err, gp.Name)
	}

	manifestBytes := bytes.Split(data, []byte("\n---\n"))
	crds, errs := component.FilterCRDs(manifestBytes)

	for _, crd := range crds {
		comp, err := component.Generate(crd)
		if err != nil {
			continue
		}
		if comp.Model.Metadata == nil {
			comp.Model.Metadata = &model.ModelDefinition_Metadata{}
		}
		if comp.Model.Metadata.AdditionalProperties == nil {
			comp.Model.Metadata.AdditionalProperties = make(map[string]interface{})
		}

		comp.Model.Metadata.AdditionalProperties["source_uri"] = gp.SourceURL
		comp.Model.Version = gp.version
		comp.Model.Name = gp.Name
		comp.Model.Category = category.CategoryDefinition{
			Name: "",
		}
		comp.Model.DisplayName = manifests.FormatToReadableString(comp.Model.Name)
		components = append(components, comp)
	}

	return components, utils.CombineErrors(errs, "\n")
}
