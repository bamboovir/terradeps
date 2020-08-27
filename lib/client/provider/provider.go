package provider

import (
	"fmt"

	"net/http"

	"github.com/bamboovir/terradeps/lib/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

const (
	// HashicorpProviderRegistryHost defination
	HashicorpProviderRegistryHost = "https://releases.hashicorp.com/"
)

// Client defination
type Client struct {
	HostURL     string
	client      *http.Client
	middlewares MiddleWareSet
}

// NewClient initializes a new API client with default values. It takes functors
// to modify values when creating it
func New(ops ...Opt) (*Client, error) {
	c := &Client{
		HostURL: HashicorpProviderRegistryHost,
		client:  &http.Client{},
	}

	for _, op := range ops {
		err := op()
	}

	return c, nil
}

/*

coreZipURL := c.coreURL(config.Terraform.Version, osName, archName)
err = getter.Get(workDir, coreZipURL)

coreZipURL = https://releases....terraform_0.13.zip

reqs {
	"addr.Providers" : [ "v1", "v2" ]
}

providercache.InstallPackage

meta {
	Provider: {
		Type: "azurerm"
		NameSpace: "hashicorp"
		HostName: "registry.terraform.io"
	}
	Version: {
		Major: 2
		Minor: 24
		Patch: 0
	}
	ProtocolVersions: [
		{
			Major: 5
			Minor: 0
			Patch: 0
		}
	]
	TargetPlatform: {
		OS: "darwin"
		Arch: "amd64"
	}
	Filename: "terraform-provider-azurerm_2.24.0_darwin_amd64.zip"
	Location: "https://releases.hashicorp.com/terraform-provider-azurerm/2.24.0...+51 more"

}
*/

// DownloadToDir defination
func (c *Client) DownloadToDir(fs afero.Fs, path string, deps *types.ProviderDeps) error {
	errmsg := fmt.Sprintf("download providers to dir [%s] failed", path)
	isDir, err := afero.IsDir(fs, path)

	if err != nil {
		log.Errorf("%v, %s", err, errmsg)
		return errors.Wrap(err, errmsg)
	}

	if !isDir {
		log.Errorf("path [%s] is not a dir", path)
		return errors.New(fmt.Sprintf("path [%s] is not a dir", path))
	}

	return nil
}
