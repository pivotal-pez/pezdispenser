package pezauth

import (
	"errors"

	goauth2 "golang.org/x/oauth2"
)

var (
	ClientID               string
	ClientSecret           string
	ErrCouldNotGetUserGUID = errors.New("query failed. unable to find matching user guid.")
	//Vars for my oauth calls
	Scopes              = []string{"https://www.googleapis.com/auth/plus.me", "https://www.googleapis.com/auth/userinfo.email"}
	AuthFailureResponse = []byte(`{"error": "not logged in as a valid user, or the access token is expired"}`)
	allowedDomains      = []string{
		"pivotal.io",
	}
	OauthConfig     *goauth2.Config
	userObjectCache = make(map[string]map[string]interface{})
	//Authentication Handler vars
	ErrInvalidCallerEmail = errors.New("Invalid user token for your requested action")
	//ErrUnparsableHash - an error for a hash that is not formed properly
	ErrUnparsableHash = errors.New("Could not parse the hash or hash was nil")
	//ErrEmptyKeyResponse - an error for a invalid or empty key
	ErrEmptyKeyResponse = errors.New("The key could not be found or was not valid")
	//ErrNoMatchInStore - error when there is no matching org in the datastore
	ErrNoMatchInStore = errors.New("Could not find a matching user org or connection failure")
	//ErrCanNotCreateOrg - error when we can not create an org
	ErrCanNotCreateOrg = errors.New("Could not create a new org")
	//ErrCanNotAddOrgRec - error when we can not add a new org record to the datastore
	ErrCanNotAddOrgRec = errors.New("Could not add a new org record")
	//ErrCantCallAcrossUsers - error when a user is trying to update a user record other than their own
	ErrCantCallAcrossUsers = errors.New("user calling another users endpoint")
	//UserMatch exported vars
	ErrNotValidActionForUser = errors.New("not a valid user to perform this action")
)

//Constants to construct my oauth calls
const (
	sessionName   = "pivotalpezauthservicesession"
	sessionSecret = "shhh.donttellanyone"
	//FailureStatus - failure response status from our unauthenticated rest endpoints
	FailureStatus = 403
	//SuccessStatus - success response status from our authenticated rest endpoints
	SuccessStatus = 200
	//HMFieldActive - name of metadata hash field containing active status
	HMFieldActive = "active"
	//HMFieldDetails - name of metadata hash field containing user and key details
	HMFieldDetails = "details"
	//EmailFieldName - fieldname for email
	EmailFieldName = "email"
	//GUIDLength - length of valid key
	GUIDLength = 36
	//HeaderKeyName - header keyname for api-key value
	HeaderKeyName = "X-API-KEY"
	//ErrInvalidKeyFormatMsg - error msg for invalid key
	ErrInvalidKeyFormatMsg = "Invalid key format"
	//DefaultSpaceName - default space name created for each org
	DefaultSpaceName = "development"
)
