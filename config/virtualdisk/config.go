package virtualdisk
import (

    ujconfig "github.com/upbound/upjet/pkg/config"

)

func Configure(p *ujconfig.Provider) {

    p.AddResourceConfigurator("vsphere_virtual_disk", func(r *ujconfig.Resource) {

    })
}
