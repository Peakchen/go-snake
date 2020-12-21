package utils

import satoriUUID "github.com/satori/go.uuid"

//github.com/satori/go.uuid:   e1b1b802-f342-41af-b887-c09772fbf9a3
//get new uuid
func GetUUID() string {
	return satoriUUID.NewV4().String()
}
