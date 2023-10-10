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
type Domain struct{}

// Each resource has in input struct, defining what arguments it accepts.
type DomainArgs struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but its generally a
	// good idea.
	Name              string  `pulumi:"name"`
	NotificationEmail *string `pulumi:"notification_email,optional"`
	Whitelabel        *string `pulumi:"whitelabel,optional"`
}

type DomainState struct {
	DomainArgs
	// Here we define a required output called result.
	Result string `pulumi:"result"`
}

func (Domain) Create(ctx p.Context, name string, input DomainArgs, preview bool) (string, DomainState, error) {
	state := DomainState{DomainArgs: input}
	if preview {
		return name, state, nil
	}

	// TODO: create the domain

	return name, state, nil
}
