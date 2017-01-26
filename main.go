package main

import (
	"fmt"
	"log"

	netpolicy "github.com/FlorianOtel/network-policies"

	"gopkg.in/yaml.v2"
)

func main() {

	fmt.Println("===> Allow All Traffic Policy Element")
	fmt.Printf("%#v\n", netpolicy.AllowAllPolicyElem)

	fmt.Println("===> Deny All Traffic Policy Element")
	fmt.Printf("%#v\n", netpolicy.DenyAllPolicyElem)

	aapemyml, err := yaml.Marshal(&netpolicy.AllowAllPolicyElem)

	if err != nil {
		log.Fatalf("YAML Marshalling error: %v", err)
	}

	fmt.Printf("===> YAML format for Policy Element 'Allow All':\n%s\n", string(aapemyml))

	np := netpolicy.NetworkPolicy{}

	if err := np.New("My Policy Name", "Ingress", "MyOrg", "MyDomain", 1000, &netpolicy.AllowAllPolicyElem); err != nil {
		log.Fatalf("KABOOM:", err)
	}

	npyml, err := yaml.Marshal(&np)

	if err != nil {
		log.Fatalf("YAML Marshalling error: %v", err)
	}

	fmt.Printf("===> YAML format for Network Policy:\n%s\n", string(npyml))

}
