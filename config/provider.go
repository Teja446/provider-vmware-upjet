/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/accenture/provider-vmware/config/virtualdisk"
	"github.com/accenture/provider-vmware/config/virtualmachine"
	"github.com/accenture/provider-vmware/config/datastorecluster"
)

const (
	resourcePrefix = "vmware"
	modulePath     = "github.com/accenture/provider-vmware"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		virtualmachine.Configure,
		virtualdisk.Configure,
		datastorecluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
