package datastorecluster
import (

    ujconfig "github.com/upbound/upjet/pkg/config"

)

func Configure(p *ujconfig.Provider) {

    p.AddResourceConfigurator("vsphere_datastore_cluster", func(r *ujconfig.Resource) {

    })
}
