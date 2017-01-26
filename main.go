package main

import (
	"fmt"

	netpolicy "github.com/FlorianOtel/network-policies"
)

func main() {
	fmt.Println("===> VSDAction")
	for k, v := range netpolicy.VSDAction {
		fmt.Println("Key: ", k, "Value: ", v)
	}

	fmt.Println("===> VSDProtocol")
	for k, v := range netpolicy.VSDProtocol {
		fmt.Println("Key: ", k, "Value: ", v)
	}

	fmt.Println("===> VSDScope")
	for k, v := range netpolicy.VSDSrcScope {
		fmt.Println("Key: ", k, "Value: ", v)
	}

	fmt.Println("===> VSDDstScope")
	for k, v := range netpolicy.VSDDstScope {
		fmt.Println("Key: ", k, "Value: ", v)
	}

	fmt.Println("===> VSDPolicyType")
	for k, v := range netpolicy.VSDPolicyType {
		fmt.Println("Key: ", k, "Value: ", v)
	}

	fmt.Println("===> All traffic")
	fmt.Printf("%#v\n", netpolicy.MatchAllTrafficSpec)

	fmt.Println("===> Allow All Traffic Policy Element")
	fmt.Printf("%#v\n", netpolicy.AllowAllPolicyElem)

	fmt.Println("===> Deny All Traffic Policy Element")
	fmt.Printf("%#v\n", netpolicy.DenyAllPolicyElem)

}
