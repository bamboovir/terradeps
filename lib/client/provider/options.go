package provider

// Opt is a configuration option to initialize a client
type Opt func(*Client) error

// Default defination
func Default(c *Client) error {
	err := WithHashicorpRegistry(c)
	if err != nil {
		return err
	}

	return nil
}

// WithHashicorpRegistry defination
func WithHashicorpRegistry(c *Client) error {
	c.HostURL = HashicorpProviderRegistryHost
	return nil
}

// WithHTTPBasicAuth defination
func WithHTTPBasicAuth(username string, password string) Opt {
	return func (c *Client) error {
		
		return 
	}
}