package client

const (
	PUBLISH_TYPE = "publish"
)

type Publish struct {
	Resource
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    Publisher string `json:"publisher,omitempty"`
    
    ResourceId string `json:"resourceId,omitempty"`
    
    ResourceType string `json:"resourceType,omitempty"`
    
    Time int `json:"time,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningInternalMessage string `json:"transitioningInternalMessage,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
}

type PublishCollection struct {
	Collection
	Data []Publish `json:"data,omitempty"`
}

type PublishClient struct {
	rancherClient *RancherClient
}

type PublishOperations interface {
	List(opts *ListOpts) (*PublishCollection, error)
	Create(opts *Publish) (*Publish, error)
	Update(existing *Publish, updates interface{}) (*Publish, error)
	ById(id string) (*Publish, error)
	Delete(container *Publish) error
}

func newPublishClient(rancherClient *RancherClient) *PublishClient {
	return &PublishClient{
		rancherClient: rancherClient,
	}
}

func (c *PublishClient) Create(container *Publish) (*Publish, error) {
	resp := &Publish{}
	err := c.rancherClient.doCreate(PUBLISH_TYPE, container, resp)
	return resp, err
}

func (c *PublishClient) Update(existing *Publish, updates interface{}) (*Publish, error) {
	resp := &Publish{}
	err := c.rancherClient.doUpdate(PUBLISH_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PublishClient) List(opts *ListOpts) (*PublishCollection, error) {
	resp := &PublishCollection{}
	err := c.rancherClient.doList(PUBLISH_TYPE, opts, resp)
	return resp, err
}

func (c *PublishClient) ById(id string) (*Publish, error) {
	resp := &Publish{}
	err := c.rancherClient.doById(PUBLISH_TYPE, id, resp)
	return resp, err
}

func (c *PublishClient) Delete(container *Publish) error {
	return c.rancherClient.doResourceDelete(PUBLISH_TYPE, &container.Resource)
}
