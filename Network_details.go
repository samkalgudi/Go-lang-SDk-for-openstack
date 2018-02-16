package main
import (
"fmt"
"github.com/rackspace/gophercloud"
"github.com/rackspace/gophercloud/openstack"
"github.com/rackspace/gophercloud/pagination"
"github.com/rackspace/gophercloud/openstack/networking/v2/networks"
)

func main() {

// Client authentication with the openstack//
opts := gophercloud.AuthOptions{
IdentityEndpoint: "http://openstack_url/v2.0",
Username: "Admin",
Password: "*******",
TenantID: "b84bd7c16a6e47479b4aa99d81095b71",
}
provider, err := openstack.AuthenticatedClient(opts)

if err != nil {
fmt.Println("AuthenticatedClient")
fmt.Println(err)
return
}
//intilizing network client
client, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
Region: "RegionOne",
})

if err != nil {
fmt.Println("NewNetworkV2 Error")
fmt.Println(err)
return
}

opts2 := networks.ListOpts{}
pager := networks.List(client, opts2)

pager.EachPage(func(page pagination.Page) (bool, error) {
networkList, err := networks.ExtractNetworks(page)

        if err != nil {
        return false, err
        }
	for _, n := range networkList{fmt.Println(n.ID, n.Name, n.Status) 
		// "n" will be a networks.Network
	}
        return true, nil
})

}


