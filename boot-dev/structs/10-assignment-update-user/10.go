package _0_assignment_update_user

type User struct {
	Name string
	Membership
}
type Membership struct {
	MessageCharLimit int
	Type             string
}

func newUser(name string, membershipType string) User {
	// ?
	user := User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: 100,
		},
	}

	if membershipType == "premium" {

		user.MessageCharLimit = 1000
	}
	return user
}
