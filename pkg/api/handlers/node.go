package handlers

import (
	"net/http"

	"gokube/pkg/api"
	"gokube/pkg/registry"

	"github.com/emicklei/go-restful/v3"
)

// NodeHandler handles Node-related HTTP requests
type NodeHandler struct {
	nodeRegistry *registry.NodeRegistry
}

// NewNodeHandler creates a new NodeHandler
func NewNodeHandler(nodeRegistry *registry.NodeRegistry) *NodeHandler {
	return &NodeHandler{nodeRegistry: nodeRegistry}
}

// CreateNode handles POST requests to create a new Node
func (h *NodeHandler) CreateNode(request *restful.Request, response *restful.Response) {
	node := new(api.Node)
	if err := request.ReadEntity(node); err != nil {
		api.WriteError(response, http.StatusBadRequest, err)
		return
	}
	//assignment 6: Hook in nodeRegistry with the handler.

	api.WriteResponse(response, http.StatusNotImplemented, node)
}

// GetNode handles GET requests to retrieve a Node
func (h *NodeHandler) GetNode(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	//assignment 7: Hook nodeRegistry.
	api.WriteResponse(response, http.StatusNotImplemented, name)

}

// UpdateNode handles PUT requests to update a Node
func (h *NodeHandler) UpdateNode(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	node := new(api.Node)
	if err := request.ReadEntity(node); err != nil {
		api.WriteError(response, http.StatusBadRequest, err)
		return
	}
	_ = name //Assignment: Hook nodeRegistry.

	api.WriteResponse(response, http.StatusNotImplemented, node)
}

// DeleteNode handles DELETE requests to remove a Node
func (h *NodeHandler) DeleteNode(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	//
	_ = name //assignment 8: Hook nodeRegistry.

	api.WriteResponse(response, http.StatusNotImplemented, name)

}

// ListNodes handles GET requests to list all Nodes
func (h *NodeHandler) ListNodes(request *restful.Request, response *restful.Response) {
	api.WriteResponse(response, http.StatusNotImplemented, "nodes")

}

// RegisterNodeRoutes registers Node routes with the WebService
func RegisterNodeRoutes(ws *restful.WebService, handler *NodeHandler) {
	ws.Route(ws.POST("/nodes").To(handler.CreateNode))
	ws.Route(ws.GET("/nodes").To(handler.ListNodes))
	ws.Route(ws.GET("/nodes/{name}").To(handler.GetNode))
	ws.Route(ws.PUT("/nodes/{name}").To(handler.UpdateNode))
	ws.Route(ws.DELETE("/nodes/{name}").To(handler.DeleteNode))
}
