package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
)

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type Alias struct{}

type AliasArgs struct {
	Domain  string `pulumi:"domain"`
	Name    string `pulumi:"name"`
	Forward string `pulumi:"forward"`
}

type AliasState struct {
	AliasArgs
	Id int `pulumi:"alias_id"`
}

func (Alias) Create(ctx p.Context, name string, input AliasArgs, preview bool) (string, AliasState, error) {
	state := AliasState{AliasArgs: input}
	if preview {
		return name, state, nil
	}

	// TODO: create the alias

	return name, state, nil
}
