// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package influxdb

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The user resource allows a user to be created on an InfluxDB server.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-influxdb/blob/master/website/docs/r/user.html.markdown.
type User struct {
	s *pulumi.ResourceState
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOpt) (*User, error) {
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["admin"] = nil
		inputs["grants"] = nil
		inputs["name"] = nil
		inputs["password"] = nil
	} else {
		inputs["admin"] = args.Admin
		inputs["grants"] = args.Grants
		inputs["name"] = args.Name
		inputs["password"] = args.Password
	}
	s, err := ctx.RegisterResource("influxdb:index/user:User", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &User{s: s}, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UserState, opts ...pulumi.ResourceOpt) (*User, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["admin"] = state.Admin
		inputs["grants"] = state.Grants
		inputs["name"] = state.Name
		inputs["password"] = state.Password
	}
	s, err := ctx.ReadResource("influxdb:index/user:User", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &User{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *User) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *User) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Mark the user as admin.
func (r *User) Admin() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["admin"])
}

// A list of grants for non-admin users
func (r *User) Grants() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["grants"])
}

// The name for the user. 
func (r *User) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The password for the user. 
func (r *User) Password() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["password"])
}

// Input properties used for looking up and filtering User resources.
type UserState struct {
	// Mark the user as admin.
	Admin interface{}
	// A list of grants for non-admin users
	Grants interface{}
	// The name for the user. 
	Name interface{}
	// The password for the user. 
	Password interface{}
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Mark the user as admin.
	Admin interface{}
	// A list of grants for non-admin users
	Grants interface{}
	// The name for the user. 
	Name interface{}
	// The password for the user. 
	Password interface{}
}
