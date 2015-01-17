package client

const (
	MOUNT_TYPE = "mount"
)

type Mount struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    InstanceId string `json:"instanceId,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    Path string `json:"path,omitempty"`
    
    Permissions string `json:"permissions,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
    VolumeId string `json:"volumeId,omitempty"`
    
}

type MountCollection struct {
	Collection
	Data []Mount `json:"data,omitempty"`
}

type MountClient struct {
	rancherClient *RancherClient
}

type MountOperations interface {
	List(opts *ListOpts) (*MountCollection, error)
	Create(opts *Mount) (*Mount, error)
	Update(existing *Mount, updates interface{}) (*Mount, error)
	ById(id string) (*Mount, error)
	Delete(container *Mount) error
}

func newMountClient(rancherClient *RancherClient) *MountClient {
	return &MountClient{
		rancherClient: rancherClient,
	}
}

func (c *MountClient) Create(container *Mount) (*Mount, error) {
	resp := &Mount{}
	err := c.rancherClient.doCreate(MOUNT_TYPE, container, resp)
	return resp, err
}

func (c *MountClient) Update(existing *Mount, updates interface{}) (*Mount, error) {
	resp := &Mount{}
	err := c.rancherClient.doUpdate(MOUNT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *MountClient) List(opts *ListOpts) (*MountCollection, error) {
	resp := &MountCollection{}
	err := c.rancherClient.doList(MOUNT_TYPE, opts, resp)
	return resp, err
}

func (c *MountClient) ById(id string) (*Mount, error) {
	resp := &Mount{}
	err := c.rancherClient.doById(MOUNT_TYPE, id, resp)
	return resp, err
}

func (c *MountClient) Delete(container *Mount) error {
	return c.rancherClient.doResourceDelete(MOUNT_TYPE, &container.Resource)
}
