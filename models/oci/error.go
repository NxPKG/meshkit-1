package oci

import (
	"fmt"

	"github.com/khulnasoft/meshkit/errors"
)

var (
	ErrAppendingLayerCode       = "meshkit-11147"
	ErrReadingFileCode          = "meshkit-11148"
	ErrUnSupportedLayerTypeCode = "meshkit-11149"
	ErrGettingLayerCode         = "meshkit-11150"
	ErrCompressingLayerCode     = "meshkit-11151"
	ErrUnTaringLayerCode        = "meshkit-11152"
	ErrGettingImageCode         = "meshkit-11153"
	ErrValidatingImageCode      = "meshkit-11154"
)

func ErrAppendingLayer(err error) error {
	return errors.New(ErrAppendingLayerCode, errors.Alert, []string{"appending content to artifact failed"}, []string{err.Error()}, []string{"layer is not compatible with the base image"}, []string{"Try using a different base image", "use a different media type for the layer"})
}

func ErrReadingFile(err error) error {
	return errors.New(ErrReadingFileCode, errors.Alert, []string{"reading file failed"}, []string{err.Error()}, []string{"failed to read the file", "Insufficient permissions"}, []string{"Try using a different file", "check if appropriate read permissions are given to the file"})
}

func ErrUnSupportedLayerType(err error) error {
	return errors.New(ErrUnSupportedLayerTypeCode, errors.Alert, []string{"unsupported layer type"}, []string{err.Error()}, []string{"layer type is not supported"}, []string{"Try using a different layer type", fmt.Sprintf("supported layer types are: %s, %s", LayerTypeTarball, LayerTypeStatic)})
}

func ErrGettingLayer(err error) error {
	return errors.New(ErrGettingLayerCode, errors.Alert, []string{"getting layer failed"}, []string{err.Error()}, []string{"failed to get the layer"}, []string{"Try using a different layer", "check if OCI image is not malformed"})
}

func ErrCompressingLayer(err error) error {
	return errors.New(ErrCompressingLayerCode, errors.Alert, []string{"compressing layer failed"}, []string{err.Error()}, []string{"failed to compress the layer"}, []string{"Try using a different layer", "check if layers are compatible with the base image"})
}

func ErrUnTaringLayer(err error) error {
	return errors.New(ErrUnTaringLayerCode, errors.Alert, []string{"untaring layer failed"}, []string{err.Error()}, []string{"failed to untar the layer"}, []string{"Try using a different layer", "check if image is not malformed"})
}

func ErrGettingImage(err error) error {
	return errors.New(ErrGettingImageCode, errors.Alert, []string{"getting image failed"}, []string{err.Error()}, []string{"failed to get the image"}, []string{"Try using a different image", "check if image is not malformed"})
}

func ErrValidatingImage(err error) error {
	return errors.New(ErrValidatingImageCode, errors.Alert, []string{"validating image failed"}, []string{err.Error()}, []string{"failed to validate the image"}, []string{"Try using a different image", "check if image is not malformed"})
}
