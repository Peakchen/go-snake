package utils

import "github.com/google/uuid"

//github.com/google/go.uuid:   f2789e88-b2ba-433c-8cc4-26c493f719df
//get new uuid
func GetGoogleUUID() string {
	return uuid.New().String()
}
