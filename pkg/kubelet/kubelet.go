package kubelet

import (
	"encoding/json"
	"fmt"
	"gokube/pkg/api"
)

type Kubelet struct {
	nodeName     string
	apiServerURL string
}

func NewKubelet(nodeName, apiServerURL string) (*Kubelet, error) {

	return &Kubelet{
		nodeName:     nodeName,
		apiServerURL: apiServerURL,
	}, nil
}

func (k *Kubelet) Start() error {
	// Register the node with the API server
	if err := k.registerNode(); err != nil {
		return fmt.Errorf("failed to register node: %w", err)
	}
	return nil
}

func (k *Kubelet) registerNode() error {
	node := &api.Node{
		ObjectMeta: api.ObjectMeta{
			Name: k.nodeName,
		},
		Status: api.NodeReady,
	}

	err2 := k.registerWithAPIServer(node)
	if err2 != nil {
		return err2
	}

	return nil
}

func (k *Kubelet) registerWithAPIServer(node *api.Node) error {
	jsonData, err := json.Marshal(node)
	if err != nil {
		return fmt.Errorf("failed to marshal node data: %w", err)
	}
	//assignment 9: Complete kubelete registration with the apiserver.
	_ = jsonData

	return nil
}
