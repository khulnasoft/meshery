package stages

import (
	"github.com/gofrs/uuid"
	"github.com/khulnasoft/meshery/server/models/pattern/core"
	meshmodel "github.com/khulnasoft/meshkit/models/meshmodel/registry"
	"github.com/khulnasoft/meshkit/models/oam/core/v1alpha1"
)

type ServiceInfoProvider interface {
	GetMesheryPatternResource(
		name string,
		namespace string,
		typ string,
		oamType string,
	) (ID *uuid.UUID, err error)
	GetServiceMesh() (name string, version string)
	GetAPIVersionForKind(kind string) string
	IsDelete() bool
}

type ServiceActionProvider interface {
	Terminate(error)
	Log(msg string)
	Provision(CompConfigPair) (string, error)
	GetRegistry() *meshmodel.RegistryManager
	Persist(string, core.Service, bool) error
	DryRun([]v1alpha1.Component) (map[string]map[string]core.DryRunResponseWrapper, error)
	Mutate(*core.Pattern) //Uses pre-defined policies/configuration to mutate the pattern
}
