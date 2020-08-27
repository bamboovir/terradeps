package types

type TerraformDeps struct {
	OS string 	   `json:"os"`,
	Arch string    `json:"arch"`
	Version string `json:"version"`
	Sha256 string  `json:"sha256"`
}

type ProviderDeps struct {
	OS string 	   `json:"os"`,
	Arch string    `json:"arch"`
	Version string `json:"version"`
	Sha256 string  `json:"sha256"`
	Source string  `json:"source"`
}

// Deps defination
type Deps struct {
	Terraform []TerraformDeps `json:"terraform"`
	Providers []ProviderDeps  `json:"providers"`
}