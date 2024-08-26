package registration

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/khulnasoft/meshkit/models/meshmodel/entity"
	"github.com/khulnasoft/meshkit/utils"
	"github.com/meshplay/schemas/models/v1alpha3/relationship"
	"github.com/meshplay/schemas/models/v1beta1/component"
	"github.com/meshplay/schemas/models/v1beta1/model"
)

type Dir struct {
	dirpath string
}

/*
The directory should contain one and only one `model`.
A directory containing multiple `model` will be invalid.
*/
func NewDir(path string) Dir {
	return Dir{dirpath: path}
}

/*
PkgUnit parses all the files inside the directory and finds out if they are any valid meshplay definitions. Valid meshplay definitions are added to the packagingUnit struct.
Invalid definitions are stored in the regErrStore with error data.
*/
func (d Dir) PkgUnit(regErrStore RegistrationErrorStore) (_ packagingUnit, err error) {
	pkg := packagingUnit{}
	// check if the given is a directory
	_, err = os.ReadDir(d.dirpath)
	if err != nil {
		return pkg, ErrDirPkgUnitParseFail(d.dirpath, fmt.Errorf("Could not read the directory: %e", err))
	}
	err = filepath.Walk(d.dirpath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		byt, _ := os.ReadFile(path)
		if byt == nil {
			return nil
		}

		var e entity.Entity
		e, err = getEntity(byt)
		if err != nil {
			regErrStore.AddInvalidDefinition(path, err)
			return nil
		}

		// set it to pkgunit
		switch e.Type() {
		case entity.Model:
			if !reflect.ValueOf(pkg.model).IsZero() {
				// currently models inside models are not handled
				return nil
			}
			model, err := utils.Cast[*model.ModelDefinition](e)
			if err != nil {
				regErrStore.AddInvalidDefinition(path, ErrGetEntity(err))
			}
			pkg.model = *model
		case entity.ComponentDefinition:
			comp, err := utils.Cast[*component.ComponentDefinition](e)
			if err != nil {
				regErrStore.AddInvalidDefinition(path, ErrGetEntity(err))
			}
			pkg.components = append(pkg.components, *comp)
		case entity.RelationshipDefinition:
			rel, err := utils.Cast[*relationship.RelationshipDefinition](e)
			if err != nil {
				regErrStore.AddInvalidDefinition(path, ErrGetEntity(err))
			}
			pkg.relationships = append(pkg.relationships, *rel)
		}
		return nil
	})
	if err != nil {
		return pkg, ErrDirPkgUnitParseFail(d.dirpath, fmt.Errorf("Could not completely walk the file tree: %e", err))
	}
	if reflect.ValueOf(pkg.model).IsZero() {
		err := fmt.Errorf("Model definition not found in imported package. Model definitions often use the filename `model.json`, but are not required to have this filename. One and exactly one entity containing schema: model.core....... ...... must be present, otherwise the model package is considered malformed..")
		regErrStore.AddInvalidDefinition(d.dirpath, err)
		return pkg, err
	}
	return pkg, err
}
