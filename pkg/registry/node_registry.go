package registry

import (
	"context"
	"errors"
	"path"

	"gokube/pkg/api"
	"gokube/pkg/storage"
)

const (
	nodePrefix = "/registry/nodes/"
)

var (
	ErrNodeNotFound      = errors.New("node not found")
	ErrNodeAlreadyExists = errors.New("node already exists")
	ErrListNodesFailed   = errors.New("failed to list nodes")
	ErrNodeInvalid       = errors.New("invalid node")
)

// NodeRegistry provides CRUD operations for Node objects
type NodeRegistry struct {
	storage storage.Storage
}

// NewNodeRegistry creates a new NodeRegistry
func NewNodeRegistry(storage storage.Storage) *NodeRegistry {
	return &NodeRegistry{storage: storage}
}

// generateKey generates the storage key for a given node name
func generateKey(prefix, name string) string {
	return path.Join(prefix, name)
}

// CreateNode stores a new Node
func (r *NodeRegistry) CreateNode(ctx context.Context, node *api.Node) error {
	//assignment 1. Implement create node.
	return nil
}

// GetNode retrieves a Node by name
func (r *NodeRegistry) GetNode(ctx context.Context, name string) (*api.Node, error) {
	//assignment 1. Implement get node.
	return nil, ErrNodeNotFound
}

// UpdateNode updates an existing Node
func (r *NodeRegistry) UpdateNode(ctx context.Context, node *api.Node) error {
	//assignment 3. Implement updateNode.
	return nil
}

// DeleteNode removes a Node by name
func (r *NodeRegistry) DeleteNode(ctx context.Context, name string) error {
	//assignment 4. Implement updateNode.

	return nil
}

// ListNodes retrieves all Nodes
func (r *NodeRegistry) ListNodes(ctx context.Context) ([]*api.Node, error) {
	//assignment 5. Implement ListNodes.

	return nil, nil
}
