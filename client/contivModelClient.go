// contivModelClient.go
// This file is auto generated by modelgen tool
// Do not edit this file manually

package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

// Link is a one way relattion between two objects
type Link struct {
	ObjType string `json:"type,omitempty"`
	ObjKey  string `json:"key,omitempty"`
}

func httpGet(url string, jdata interface{}) error {

	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	switch {
	case r.StatusCode == int(404):
		return errors.New("Page not found!")
	case r.StatusCode == int(403):
		return errors.New("Access denied!")
	case r.StatusCode == int(500):
		response, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		return errors.New(string(response))

	case r.StatusCode != int(200):
		log.Debugf("GET Status '%s' status code %d \n", r.Status, r.StatusCode)
		return errors.New(r.Status)
	}

	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(response, jdata); err != nil {
		return err
	}

	return nil
}

func httpDelete(url string) error {

	req, err := http.NewRequest("DELETE", url, nil)

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// body, _ := ioutil.ReadAll(r.Body)

	switch {
	case r.StatusCode == int(404):
		// return errors.New("Page not found!")
		return nil
	case r.StatusCode == int(403):
		return errors.New("Access denied!")
	case r.StatusCode == int(500):
		response, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		return errors.New(string(response))

	case r.StatusCode != int(200):
		log.Debugf("DELETE Status '%s' status code %d \n", r.Status, r.StatusCode)
		return errors.New(r.Status)
	}

	return nil
}

func httpPost(url string, jdata interface{}) error {
	buf, err := json.Marshal(jdata)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(buf)
	r, err := http.Post(url, "application/json", body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	switch {
	case r.StatusCode == int(404):
		return errors.New("Page not found!")
	case r.StatusCode == int(403):
		return errors.New("Access denied!")
	case r.StatusCode == int(500):
		response, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		return errors.New(string(response))

	case r.StatusCode != int(200):
		log.Debugf("POST Status '%s' status code %d \n", r.Status, r.StatusCode)
		return errors.New(r.Status)
	}

	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	log.Debugf(string(response))

	return nil
}

// ContivClient has the contiv model client instance
type ContivClient struct {
	baseURL string
}

// NewContivClient returns a new client instance
func NewContivClient(baseURL string) (*ContivClient, error) {
	client := ContivClient{
		baseURL: baseURL,
	}

	return &client, nil
}

type AppProfile struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	AppProfileName string   `json:"appProfileName,omitempty"` // Application Profile Name
	EndpointGroups []string `json:"endpointGroups,omitempty"`
	TenantName     string   `json:"tenantName,omitempty"` // Tenant Name

	// add link-sets and links
	LinkSets AppProfileLinkSets `json:"link-sets,omitempty"`
	Links    AppProfileLinks    `json:"links,omitempty"`
}

type AppProfileLinkSets struct {
	EndpointGroups map[string]Link `json:"EndpointGroups,omitempty"`
}

type AppProfileLinks struct {
	Tenant Link `json:"Tenant,omitempty"`
}

type AppProfileInspect struct {
	Config AppProfile
}

type Bgp struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	As         string `json:"as,omitempty"`          // AS id
	Hostname   string `json:"hostname,omitempty"`    // host name
	Neighbor   string `json:"neighbor,omitempty"`    // Bgp  neighbor
	NeighborAs string `json:"neighbor-as,omitempty"` // AS id
	Routerip   string `json:"routerip,omitempty"`    // Bgp router intf ip

}

type BgpInspect struct {
	Config Bgp
}

type EndpointGroup struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	ExtContractsGrps []string `json:"extContractsGrps,omitempty"`
	GroupName        string   `json:"groupName,omitempty"`   // Group name
	NetworkName      string   `json:"networkName,omitempty"` // Network
	Policies         []string `json:"policies,omitempty"`
	TenantName       string   `json:"tenantName,omitempty"` // Tenant

	// add link-sets and links
	LinkSets EndpointGroupLinkSets `json:"link-sets,omitempty"`
	Links    EndpointGroupLinks    `json:"links,omitempty"`
}

type EndpointGroupLinkSets struct {
	ExtContractsGrps map[string]Link `json:"ExtContractsGrps,omitempty"`
	Policies         map[string]Link `json:"Policies,omitempty"`
	Services         map[string]Link `json:"Services,omitempty"`
}

type EndpointGroupLinks struct {
	AppProfile Link `json:"AppProfile,omitempty"`
	Network    Link `json:"Network,omitempty"`
	Tenant     Link `json:"Tenant,omitempty"`
}

type EndpointGroupInspect struct {
	Config EndpointGroup
}

type ExtContractsGroup struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	Contracts          []string `json:"contracts,omitempty"`
	ContractsGroupName string   `json:"contractsGroupName,omitempty"` // Contracts group name
	ContractsType      string   `json:"contractsType,omitempty"`      // Contracts type
	TenantName         string   `json:"tenantName,omitempty"`         // Tenant name

	// add link-sets and links
	LinkSets ExtContractsGroupLinkSets `json:"link-sets,omitempty"`
}

type ExtContractsGroupLinkSets struct {
	EndpointGroups map[string]Link `json:"EndpointGroups,omitempty"`
}

type Global struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	Name             string `json:"name,omitempty"`             // name of this block(must be 'global')
	NetworkInfraType string `json:"networkInfraType,omitempty"` // Network infrastructure type
	Vlans            string `json:"vlans,omitempty"`            // Allowed vlan range
	Vxlans           string `json:"vxlans,omitempty"`           // Allwed vxlan range

}

type GlobalInspect struct {
	Config Global
}

type Network struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	Encap       string `json:"encap,omitempty"`       // Encapsulation
	Gateway     string `json:"gateway,omitempty"`     // Gateway
	Ipv6Gateway string `json:"ipv6Gateway,omitempty"` // IPv6Gateway
	Ipv6Subnet  string `json:"ipv6Subnet,omitempty"`  // IPv6Subnet
	NetworkName string `json:"networkName,omitempty"` // Network name
	NwType      string `json:"nwType,omitempty"`      // Network Type
	PktTag      int    `json:"pktTag,omitempty"`      // Vlan/Vxlan Tag
	Subnet      string `json:"subnet,omitempty"`      // Subnet
	TenantName  string `json:"tenantName,omitempty"`  // Tenant Name

	// add link-sets and links
	LinkSets NetworkLinkSets `json:"link-sets,omitempty"`
	Links    NetworkLinks    `json:"links,omitempty"`
}

type NetworkLinkSets struct {
	EndpointGroups map[string]Link `json:"EndpointGroups,omitempty"`
	Servicelbs     map[string]Link `json:"Servicelbs,omitempty"`
	Services       map[string]Link `json:"Services,omitempty"`
}

type NetworkLinks struct {
	Tenant Link `json:"Tenant,omitempty"`
}

type NetworkOper struct {
	AllocatedAddressesCount int    `json:"allocatedAddressesCount,omitempty"` // Vlan/Vxlan Tag
	AllocatedIPAddresses    string `json:"allocatedIPAddresses,omitempty"`    // allocated IP addresses
	DnsServerIP             string `json:"dnsServerIP,omitempty"`             // dns IP for the network
	ExternalPktTag          int    `json:"externalPktTag,omitempty"`          // external packet tag
	NumEndpoints            int    `json:"numEndpoints,omitempty"`            // external packet tag
	PktTag                  int    `json:"pktTag,omitempty"`                  // internal packet tag

}

type NetworkInspect struct {
	Config Network

	Oper NetworkOper
}

type Policy struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	PolicyName string `json:"policyName,omitempty"` // Policy Name
	TenantName string `json:"tenantName,omitempty"` // Tenant Name

	// add link-sets and links
	LinkSets PolicyLinkSets `json:"link-sets,omitempty"`
	Links    PolicyLinks    `json:"links,omitempty"`
}

type PolicyLinkSets struct {
	EndpointGroups map[string]Link `json:"EndpointGroups,omitempty"`
	Rules          map[string]Link `json:"Rules,omitempty"`
}

type PolicyLinks struct {
	Tenant Link `json:"Tenant,omitempty"`
}

type PolicyInspect struct {
	Config Policy
}

type Rule struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	Action            string `json:"action,omitempty"`            // Action
	Direction         string `json:"direction,omitempty"`         // Direction
	FromEndpointGroup string `json:"fromEndpointGroup,omitempty"` // From Endpoint Group
	FromIpAddress     string `json:"fromIpAddress,omitempty"`     // IP Address
	FromNetwork       string `json:"fromNetwork,omitempty"`       // From Network
	PolicyName        string `json:"policyName,omitempty"`        // Policy Name
	Port              int    `json:"port,omitempty"`              // Port No
	Priority          int    `json:"priority,omitempty"`          // Priority
	Protocol          string `json:"protocol,omitempty"`          // Protocol
	RuleID            string `json:"ruleId,omitempty"`            // Rule Id
	TenantName        string `json:"tenantName,omitempty"`        // Tenant Name
	ToEndpointGroup   string `json:"toEndpointGroup,omitempty"`   // To Endpoint Group
	ToIpAddress       string `json:"toIpAddress,omitempty"`       // IP Address
	ToNetwork         string `json:"toNetwork,omitempty"`         // To Network

	// add link-sets and links
	LinkSets RuleLinkSets `json:"link-sets,omitempty"`
}

type RuleLinkSets struct {
	Policies map[string]Link `json:"Policies,omitempty"`
}

type RuleInspect struct {
	Config Rule
}

type ServiceLB struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	IpAddress   string   `json:"ipAddress,omitempty"`   // Service ip
	NetworkName string   `json:"networkName,omitempty"` // Service network name
	Ports       []string `json:"ports,omitempty"`
	Selectors   []string `json:"selectors,omitempty"`
	ServiceName string   `json:"serviceName,omitempty"` // service name
	TenantName  string   `json:"tenantName,omitempty"`  // Tenant Name

	Links ServiceLBLinks `json:"links,omitempty"`
}

type ServiceLBLinks struct {
	Network Link `json:"Network,omitempty"`
	Tenant  Link `json:"Tenant,omitempty"`
}

type ServiceLBInspect struct {
	Config ServiceLB
}

type Tenant struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	DefaultNetwork string `json:"defaultNetwork,omitempty"` // Network name
	TenantName     string `json:"tenantName,omitempty"`     // Tenant Name

	// add link-sets and links
	LinkSets TenantLinkSets `json:"link-sets,omitempty"`
}

type TenantLinkSets struct {
	AppProfiles    map[string]Link `json:"AppProfiles,omitempty"`
	EndpointGroups map[string]Link `json:"EndpointGroups,omitempty"`
	Networks       map[string]Link `json:"Networks,omitempty"`
	Policies       map[string]Link `json:"Policies,omitempty"`
	Servicelbs     map[string]Link `json:"Servicelbs,omitempty"`
	VolumeProfiles map[string]Link `json:"VolumeProfiles,omitempty"`
	Volumes        map[string]Link `json:"Volumes,omitempty"`
}

type TenantInspect struct {
	Config Tenant
}

type Volume struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	DatastoreType string `json:"datastoreType,omitempty"` //
	MountPoint    string `json:"mountPoint,omitempty"`    //
	PoolName      string `json:"poolName,omitempty"`      //
	Size          string `json:"size,omitempty"`          //
	TenantName    string `json:"tenantName,omitempty"`    // Tenant Name
	VolumeName    string `json:"volumeName,omitempty"`    // Volume Name

	// add link-sets and links
	LinkSets VolumeLinkSets `json:"link-sets,omitempty"`
	Links    VolumeLinks    `json:"links,omitempty"`
}

type VolumeLinkSets struct {
	ServiceInstances map[string]Link `json:"ServiceInstances,omitempty"`
}

type VolumeLinks struct {
	Tenant Link `json:"Tenant,omitempty"`
}

type VolumeInspect struct {
	Config Volume
}

type VolumeProfile struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	DatastoreType     string `json:"datastoreType,omitempty"`     //
	MountPoint        string `json:"mountPoint,omitempty"`        //
	PoolName          string `json:"poolName,omitempty"`          //
	Size              string `json:"size,omitempty"`              //
	TenantName        string `json:"tenantName,omitempty"`        // Tenant Name
	VolumeProfileName string `json:"volumeProfileName,omitempty"` // Volume profile Name

	// add link-sets and links
	LinkSets VolumeProfileLinkSets `json:"link-sets,omitempty"`
	Links    VolumeProfileLinks    `json:"links,omitempty"`
}

type VolumeProfileLinkSets struct {
	Services map[string]Link `json:"Services,omitempty"`
}

type VolumeProfileLinks struct {
	Tenant Link `json:"Tenant,omitempty"`
}

type VolumeProfileInspect struct {
	Config VolumeProfile
}

// AppProfilePost posts the appProfile object
func (c *ContivClient) AppProfilePost(obj *AppProfile) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.AppProfileName
	url := c.baseURL + "/api/v1/appProfiles/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating appProfile %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// AppProfileList lists all appProfile objects
func (c *ContivClient) AppProfileList() (*[]*AppProfile, error) {
	// build key and URL
	url := c.baseURL + "/api/v1/appProfiles/"

	// http get the object
	var objList []*AppProfile
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting appProfiles. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// AppProfileGet gets the appProfile object
func (c *ContivClient) AppProfileGet(tenantName string, appProfileName string) (*AppProfile, error) {
	// build key and URL
	keyStr := tenantName + ":" + appProfileName
	url := c.baseURL + "/api/v1/appProfiles/" + keyStr + "/"

	// http get the object
	var obj AppProfile
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting appProfile %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// AppProfileInspect gets the appProfileInspect object
func (c *ContivClient) AppProfileInspect(tenantName string, appProfileName string) (*AppProfileInspect, error) {
	// build key and URL
	keyStr := tenantName + ":" + appProfileName
	url := c.baseURL + "/api/v1/inspect/appProfiles/" + keyStr + "/"

	// http get the object
	var obj AppProfileInspect
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting appProfile %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// AppProfileDelete deletes the appProfile object
func (c *ContivClient) AppProfileDelete(tenantName string, appProfileName string) error {
	// build key and URL
	keyStr := tenantName + ":" + appProfileName
	url := c.baseURL + "/api/v1/appProfiles/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting appProfile %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// BgpPost posts the Bgp object
func (c *ContivClient) BgpPost(obj *Bgp) error {
	// build key and URL
	keyStr := obj.Hostname
	url := c.baseURL + "/api/v1/Bgps/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating Bgp %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// BgpList lists all Bgp objects
func (c *ContivClient) BgpList() (*[]*Bgp, error) {
	// build key and URL
	url := c.baseURL + "/api/v1/Bgps/"

	// http get the object
	var objList []*Bgp
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting Bgps. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// BgpGet gets the Bgp object
func (c *ContivClient) BgpGet(hostname string) (*Bgp, error) {
	// build key and URL
	keyStr := hostname
	url := c.baseURL + "/api/v1/Bgps/" + keyStr + "/"

	// http get the object
	var obj Bgp
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting Bgp %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// BgpInspect gets the BgpInspect object
func (c *ContivClient) BgpInspect(hostname string) (*BgpInspect, error) {
	// build key and URL
	keyStr := hostname
	url := c.baseURL + "/api/v1/inspect/Bgps/" + keyStr + "/"

	// http get the object
	var obj BgpInspect
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting Bgp %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// BgpDelete deletes the Bgp object
func (c *ContivClient) BgpDelete(hostname string) error {
	// build key and URL
	keyStr := hostname
	url := c.baseURL + "/api/v1/Bgps/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting Bgp %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// EndpointGroupPost posts the endpointGroup object
func (c *ContivClient) EndpointGroupPost(obj *EndpointGroup) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.GroupName
	url := c.baseURL + "/api/v1/endpointGroups/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating endpointGroup %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// EndpointGroupList lists all endpointGroup objects
func (c *ContivClient) EndpointGroupList() (*[]*EndpointGroup, error) {
	// build key and URL
	url := c.baseURL + "/api/v1/endpointGroups/"

	// http get the object
	var objList []*EndpointGroup
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting endpointGroups. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// EndpointGroupGet gets the endpointGroup object
func (c *ContivClient) EndpointGroupGet(tenantName string, groupName string) (*EndpointGroup, error) {
	// build key and URL
	keyStr := tenantName + ":" + groupName
	url := c.baseURL + "/api/v1/endpointGroups/" + keyStr + "/"

	// http get the object
	var obj EndpointGroup
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting endpointGroup %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// EndpointGroupInspect gets the endpointGroupInspect object
func (c *ContivClient) EndpointGroupInspect(tenantName string, groupName string) (*EndpointGroupInspect, error) {
	// build key and URL
	keyStr := tenantName + ":" + groupName
	url := c.baseURL + "/api/v1/inspect/endpointGroups/" + keyStr + "/"

	// http get the object
	var obj EndpointGroupInspect
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting endpointGroup %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// EndpointGroupDelete deletes the endpointGroup object
func (c *ContivClient) EndpointGroupDelete(tenantName string, groupName string) error {
	// build key and URL
	keyStr := tenantName + ":" + groupName
	url := c.baseURL + "/api/v1/endpointGroups/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting endpointGroup %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// ExtContractsGroupPost posts the extContractsGroup object
func (c *ContivClient) ExtContractsGroupPost(obj *ExtContractsGroup) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.ContractsGroupName
	url := c.baseURL + "/api/extContractsGroups/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating extContractsGroup %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// ExtContractsGroupList lists all extContractsGroup objects
func (c *ContivClient) ExtContractsGroupList() (*[]*ExtContractsGroup, error) {
	// build key and URL
	url := c.baseURL + "/api/extContractsGroups/"

	// http get the object
	var objList []*ExtContractsGroup
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting extContractsGroups. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// ExtContractsGroupGet gets the extContractsGroup object
func (c *ContivClient) ExtContractsGroupGet(tenantName string, contractsGroupName string) (*ExtContractsGroup, error) {
	// build key and URL
	keyStr := tenantName + ":" + contractsGroupName
	url := c.baseURL + "/api/extContractsGroups/" + keyStr + "/"

	// http get the object
	var obj ExtContractsGroup
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting extContractsGroup %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// ExtContractsGroupDelete deletes the extContractsGroup object
func (c *ContivClient) ExtContractsGroupDelete(tenantName string, contractsGroupName string) error {
	// build key and URL
	keyStr := tenantName + ":" + contractsGroupName
	url := c.baseURL + "/api/extContractsGroups/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting extContractsGroup %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// GlobalPost posts the global object
func (c *ContivClient) GlobalPost(obj *Global) error {
	// build key and URL
	keyStr := obj.Name
	url := c.baseURL + "/api/v1/globals/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating global %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// GlobalList lists all global objects
func (c *ContivClient) GlobalList() (*[]*Global, error) {
	// build key and URL
	url := c.baseURL + "/api/v1/globals/"

	// http get the object
	var objList []*Global
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting globals. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// GlobalGet gets the global object
func (c *ContivClient) GlobalGet(name string) (*Global, error) {
	// build key and URL
	keyStr := name
	url := c.baseURL + "/api/v1/globals/" + keyStr + "/"

	// http get the object
	var obj Global
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting global %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// GlobalInspect gets the globalInspect object
func (c *ContivClient) GlobalInspect(name string) (*GlobalInspect, error) {
	// build key and URL
	keyStr := name
	url := c.baseURL + "/api/v1/inspect/globals/" + keyStr + "/"

	// http get the object
	var obj GlobalInspect
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting global %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// GlobalDelete deletes the global object
func (c *ContivClient) GlobalDelete(name string) error {
	// build key and URL
	keyStr := name
	url := c.baseURL + "/api/v1/globals/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting global %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// NetworkPost posts the network object
func (c *ContivClient) NetworkPost(obj *Network) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.NetworkName
	url := c.baseURL + "/api/v1/networks/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating network %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// NetworkList lists all network objects
func (c *ContivClient) NetworkList() (*[]*Network, error) {
	// build key and URL
	url := c.baseURL + "/api/v1/networks/"

	// http get the object
	var objList []*Network
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting networks. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// NetworkGet gets the network object
func (c *ContivClient) NetworkGet(tenantName string, networkName string) (*Network, error) {
	// build key and URL
	keyStr := tenantName + ":" + networkName
	url := c.baseURL + "/api/v1/networks/" + keyStr + "/"

	// http get the object
	var obj Network
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting network %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// NetworkInspect gets the networkInspect object
func (c *ContivClient) NetworkInspect(tenantName string, networkName string) (*NetworkInspect, error) {
	// build key and URL
	keyStr := tenantName + ":" + networkName
	url := c.baseURL + "/api/v1/inspect/networks/" + keyStr + "/"

	// http get the object
	var obj NetworkInspect
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting network %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// NetworkDelete deletes the network object
func (c *ContivClient) NetworkDelete(tenantName string, networkName string) error {
	// build key and URL
	keyStr := tenantName + ":" + networkName
	url := c.baseURL + "/api/v1/networks/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting network %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// PolicyPost posts the policy object
func (c *ContivClient) PolicyPost(obj *Policy) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.PolicyName
	url := c.baseURL + "/api/v1/policys/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating policy %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// PolicyList lists all policy objects
func (c *ContivClient) PolicyList() (*[]*Policy, error) {
	// build key and URL
	url := c.baseURL + "/api/v1/policys/"

	// http get the object
	var objList []*Policy
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting policys. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// PolicyGet gets the policy object
func (c *ContivClient) PolicyGet(tenantName string, policyName string) (*Policy, error) {
	// build key and URL
	keyStr := tenantName + ":" + policyName
	url := c.baseURL + "/api/v1/policys/" + keyStr + "/"

	// http get the object
	var obj Policy
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting policy %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// PolicyInspect gets the policyInspect object
func (c *ContivClient) PolicyInspect(tenantName string, policyName string) (*PolicyInspect, error) {
	// build key and URL
	keyStr := tenantName + ":" + policyName
	url := c.baseURL + "/api/v1/inspect/policys/" + keyStr + "/"

	// http get the object
	var obj PolicyInspect
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting policy %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// PolicyDelete deletes the policy object
func (c *ContivClient) PolicyDelete(tenantName string, policyName string) error {
	// build key and URL
	keyStr := tenantName + ":" + policyName
	url := c.baseURL + "/api/v1/policys/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting policy %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// RulePost posts the rule object
func (c *ContivClient) RulePost(obj *Rule) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.PolicyName + ":" + obj.RuleID
	url := c.baseURL + "/api/v1/rules/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating rule %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// RuleList lists all rule objects
func (c *ContivClient) RuleList() (*[]*Rule, error) {
	// build key and URL
	url := c.baseURL + "/api/v1/rules/"

	// http get the object
	var objList []*Rule
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting rules. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// RuleGet gets the rule object
func (c *ContivClient) RuleGet(tenantName string, policyName string, ruleId string) (*Rule, error) {
	// build key and URL
	keyStr := tenantName + ":" + policyName + ":" + ruleId
	url := c.baseURL + "/api/v1/rules/" + keyStr + "/"

	// http get the object
	var obj Rule
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting rule %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// RuleInspect gets the ruleInspect object
func (c *ContivClient) RuleInspect(tenantName string, policyName string, ruleId string) (*RuleInspect, error) {
	// build key and URL
	keyStr := tenantName + ":" + policyName + ":" + ruleId
	url := c.baseURL + "/api/v1/inspect/rules/" + keyStr + "/"

	// http get the object
	var obj RuleInspect
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting rule %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// RuleDelete deletes the rule object
func (c *ContivClient) RuleDelete(tenantName string, policyName string, ruleId string) error {
	// build key and URL
	keyStr := tenantName + ":" + policyName + ":" + ruleId
	url := c.baseURL + "/api/v1/rules/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting rule %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// ServiceLBPost posts the serviceLB object
func (c *ContivClient) ServiceLBPost(obj *ServiceLB) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.ServiceName
	url := c.baseURL + "/api/v1/serviceLBs/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating serviceLB %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// ServiceLBList lists all serviceLB objects
func (c *ContivClient) ServiceLBList() (*[]*ServiceLB, error) {
	// build key and URL
	url := c.baseURL + "/api/v1/serviceLBs/"

	// http get the object
	var objList []*ServiceLB
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting serviceLBs. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// ServiceLBGet gets the serviceLB object
func (c *ContivClient) ServiceLBGet(tenantName string, serviceName string) (*ServiceLB, error) {
	// build key and URL
	keyStr := tenantName + ":" + serviceName
	url := c.baseURL + "/api/v1/serviceLBs/" + keyStr + "/"

	// http get the object
	var obj ServiceLB
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting serviceLB %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// ServiceLBInspect gets the serviceLBInspect object
func (c *ContivClient) ServiceLBInspect(tenantName string, serviceName string) (*ServiceLBInspect, error) {
	// build key and URL
	keyStr := tenantName + ":" + serviceName
	url := c.baseURL + "/api/v1/inspect/serviceLBs/" + keyStr + "/"

	// http get the object
	var obj ServiceLBInspect
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting serviceLB %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// ServiceLBDelete deletes the serviceLB object
func (c *ContivClient) ServiceLBDelete(tenantName string, serviceName string) error {
	// build key and URL
	keyStr := tenantName + ":" + serviceName
	url := c.baseURL + "/api/v1/serviceLBs/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting serviceLB %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// TenantPost posts the tenant object
func (c *ContivClient) TenantPost(obj *Tenant) error {
	// build key and URL
	keyStr := obj.TenantName
	url := c.baseURL + "/api/v1/tenants/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating tenant %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// TenantList lists all tenant objects
func (c *ContivClient) TenantList() (*[]*Tenant, error) {
	// build key and URL
	url := c.baseURL + "/api/v1/tenants/"

	// http get the object
	var objList []*Tenant
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting tenants. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// TenantGet gets the tenant object
func (c *ContivClient) TenantGet(tenantName string) (*Tenant, error) {
	// build key and URL
	keyStr := tenantName
	url := c.baseURL + "/api/v1/tenants/" + keyStr + "/"

	// http get the object
	var obj Tenant
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting tenant %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// TenantInspect gets the tenantInspect object
func (c *ContivClient) TenantInspect(tenantName string) (*TenantInspect, error) {
	// build key and URL
	keyStr := tenantName
	url := c.baseURL + "/api/v1/inspect/tenants/" + keyStr + "/"

	// http get the object
	var obj TenantInspect
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting tenant %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// TenantDelete deletes the tenant object
func (c *ContivClient) TenantDelete(tenantName string) error {
	// build key and URL
	keyStr := tenantName
	url := c.baseURL + "/api/v1/tenants/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting tenant %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// VolumePost posts the volume object
func (c *ContivClient) VolumePost(obj *Volume) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.VolumeName
	url := c.baseURL + "/api/v1/volumes/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating volume %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// VolumeList lists all volume objects
func (c *ContivClient) VolumeList() (*[]*Volume, error) {
	// build key and URL
	url := c.baseURL + "/api/v1/volumes/"

	// http get the object
	var objList []*Volume
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting volumes. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// VolumeGet gets the volume object
func (c *ContivClient) VolumeGet(tenantName string, volumeName string) (*Volume, error) {
	// build key and URL
	keyStr := tenantName + ":" + volumeName
	url := c.baseURL + "/api/v1/volumes/" + keyStr + "/"

	// http get the object
	var obj Volume
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting volume %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// VolumeInspect gets the volumeInspect object
func (c *ContivClient) VolumeInspect(tenantName string, volumeName string) (*VolumeInspect, error) {
	// build key and URL
	keyStr := tenantName + ":" + volumeName
	url := c.baseURL + "/api/v1/inspect/volumes/" + keyStr + "/"

	// http get the object
	var obj VolumeInspect
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting volume %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// VolumeDelete deletes the volume object
func (c *ContivClient) VolumeDelete(tenantName string, volumeName string) error {
	// build key and URL
	keyStr := tenantName + ":" + volumeName
	url := c.baseURL + "/api/v1/volumes/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting volume %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// VolumeProfilePost posts the volumeProfile object
func (c *ContivClient) VolumeProfilePost(obj *VolumeProfile) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.VolumeProfileName
	url := c.baseURL + "/api/v1/volumeProfiles/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating volumeProfile %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// VolumeProfileList lists all volumeProfile objects
func (c *ContivClient) VolumeProfileList() (*[]*VolumeProfile, error) {
	// build key and URL
	url := c.baseURL + "/api/v1/volumeProfiles/"

	// http get the object
	var objList []*VolumeProfile
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting volumeProfiles. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// VolumeProfileGet gets the volumeProfile object
func (c *ContivClient) VolumeProfileGet(tenantName string, volumeProfileName string) (*VolumeProfile, error) {
	// build key and URL
	keyStr := tenantName + ":" + volumeProfileName
	url := c.baseURL + "/api/v1/volumeProfiles/" + keyStr + "/"

	// http get the object
	var obj VolumeProfile
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting volumeProfile %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// VolumeProfileInspect gets the volumeProfileInspect object
func (c *ContivClient) VolumeProfileInspect(tenantName string, volumeProfileName string) (*VolumeProfileInspect, error) {
	// build key and URL
	keyStr := tenantName + ":" + volumeProfileName
	url := c.baseURL + "/api/v1/inspect/volumeProfiles/" + keyStr + "/"

	// http get the object
	var obj VolumeProfileInspect
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting volumeProfile %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// VolumeProfileDelete deletes the volumeProfile object
func (c *ContivClient) VolumeProfileDelete(tenantName string, volumeProfileName string) error {
	// build key and URL
	keyStr := tenantName + ":" + volumeProfileName
	url := c.baseURL + "/api/v1/volumeProfiles/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting volumeProfile %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}
