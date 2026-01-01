package main

import (
	"errors"
)

type user struct {
	name        string
	phoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	// ?
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	result := make(map[string]user)
	for i := range names {
		result[names[i]] = user{name: names[i], phoneNumber: phoneNumbers[i]}
	}
	return result, nil
}
