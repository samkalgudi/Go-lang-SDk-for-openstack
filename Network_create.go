package main
import (
"fmt"
"github.com/rackspace/gophercloud"
"github.com/rackspace/gophercloud/openstack"
//"github.com/rackspace/gophercloud/pagination"
"github.com/rackspace/gophercloud/openstack/networking/v2/networks"
"github.com/rackspace/gophercloud/openstack/networking/v2/subnets"
)

func main() {

// Client authentication with the openstack//
opts := gophercloud.AuthOptions{
IdentityEndpoint: "http://10.100.60.231:5000/v2.0",
Username: "osmosis",
Password: "osmosis",
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

// Creating network//
opts2 := networks.CreateOpts{Name: "Test123", AdminStateUp: networks.Up}
network, err := networks.Create(client, opts2).Extract()
fmt.Println("network created")
fmt.Println(network.ID)

if err != nil {
fmt.Println("Network creation error")
fmt.Println(err)
return
}
// Creating Subnets//
opts3 := subnets.CreateOpts{
NetworkID:  network.ID,
CIDR:       "10.0.10.0/24",
IPVersion:  subnets.IPv4,
Name:       network.Name,
}
subnet, err := subnets.Create(client, opts3).Extract()
fmt.Println("Subnet created")
fmt.Println(subnet.ID)

if err != nil {
fmt.Println("subnet creation error")
fmt.Println(err)
return
}

}

