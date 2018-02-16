package main
import (
"fmt"
"github.com/rackspace/gophercloud"
"github.com/rackspace/gophercloud/openstack"
"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
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
//Intilizing Compute client
client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
	Region: "RegionOne",
})

//Creating Server......
server, err := servers.Create(client, servers.CreateOpts{
  Name:       "Test123",
  FlavorName: "m1.tiny",
  ImageName: "cirros",
}).Extract()

if err != nil {
  fmt.Println("Unable to create server: ", err)
}
fmt.Println("Server ID: ", server.ID)

}
