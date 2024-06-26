package models

import (
	"github.com/khulnasoft/meshsync/pkg/model"
)

type MeshSyncResourcesAPIResponse struct {
	Page       int                        `json:"page"`
	PageSize   int                        `json:"page_size"`
	TotalCount int64                      `json:"total_count"`
	Resources  []model.KubernetesResource `json:"resources"`
}

type MeshSyncResourcesKindsAPIResponse struct {
	Kinds      []string `json:"kinds"`
	Page       int      `json:"page"`
	PageSize   int      `json:"page_size"`
	TotalCount int64    `json:"total_count"`
}
