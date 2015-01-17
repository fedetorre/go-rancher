package client

const (
	HOST_TYPE = "host"
)

type Host struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    AgentId string `json:"agentId,omitempty"`
    
    ComputeFree int `json:"computeFree,omitempty"`
    
    ComputeTotal int `json:"computeTotal,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    IsPublic bool `json:"isPublic,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    PhysicalHostId string `json:"physicalHostId,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uri string `json:"uri,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
    ZoneId string `json:"zoneId,omitempty"`
    
}

type HostCollection struct {
	Collection
	Data []Host `json:"data,omitempty"`
}

type HostClient struct {
	rancherClient *RancherClient
}

type HostOperations interface {
	List(opts *ListOpts) (*HostCollection, error)
	Create(opts *Host) (*Host, error)
	Update(existing *Host, updates interface{}) (*Host, error)
	ById(id string) (*Host, error)
	Delete(container *Host) error
}

func newHostClient(rancherClient *RancherClient) *HostClient {
	return &HostClient{
		rancherClient: rancherClient,
	}
}

func (c *HostClient) Create(container *Host) (*Host, error) {
	resp := &Host{}
	err := c.rancherClient.doCreate(HOST_TYPE, container, resp)
	return resp, err
}

func (c *HostClient) Update(existing *Host, updates interface{}) (*Host, error) {
	resp := &Host{}
	err := c.rancherClient.doUpdate(HOST_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HostClient) List(opts *ListOpts) (*HostCollection, error) {
	resp := &HostCollection{}
	err := c.rancherClient.doList(HOST_TYPE, opts, resp)
	return resp, err
}

func (c *HostClient) ById(id string) (*Host, error) {
	resp := &Host{}
	err := c.rancherClient.doById(HOST_TYPE, id, resp)
	return resp, err
}

func (c *HostClient) Delete(container *Host) error {
	return c.rancherClient.doResourceDelete(HOST_TYPE, &container.Resource)
}
