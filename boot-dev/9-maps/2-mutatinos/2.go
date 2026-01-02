package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	// ?
	item, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if !item.scheduledForDeletion {
		return false, nil
	}
	delete(users, item.name)
	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
