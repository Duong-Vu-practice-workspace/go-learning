package __struct_methods

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (auInfo authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", auInfo.username, auInfo.password)
}
