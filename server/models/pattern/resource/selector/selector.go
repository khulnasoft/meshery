package selector

import (
	"fmt"

	meshmodel "github.com/khulnasoft/meshkit/models/meshmodel/registry"
)

const (
	CoreResource = "pattern.meshery.khulnasoft.com/core"
	MeshResource = "pattern.meshery.khulnasoft.com/mesh/workload"
	K8sResource  = "pattern.meshery.khulnasoft.com/k8s"
)

type Helpers interface {
	GetServiceMesh() (name string, version string)
	GetAPIVersionForKind(kind string) string
}

type Selector struct {
	registry *meshmodel.RegistryManager
	helpers  Helpers
}

func New(reg *meshmodel.RegistryManager, helpers Helpers) *Selector {
	return &Selector{
		registry: reg,
		helpers:  helpers,
	}
}

func generateTraitKey(name string) string {
	return fmt.Sprintf(
		"/meshery/registry/definition/%s/%s/%s",
		"core.oam.dev/v1alpha1",
		"TraitDefinition",
		name,
	)
}
