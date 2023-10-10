package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Config struct {
	ApiKey  string `pulumi:"api_key" provider:"secret"`
	Version string `pulumi:"version"`
}

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.ApiKey, "API key for the ImprovMX")
}

func (c *Config) Configure(ctx p.Context) error {
	return nil
}
