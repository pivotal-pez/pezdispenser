package pezauth

import (
	"fmt"
	"log"
	"strings"

	"github.com/martini-contrib/oauth2"
	"github.com/pivotalservices/pezdispenser/cloudfoundryclient"
	"gopkg.in/mgo.v2/bson"
)

type (
	//OrgManager - interface to the org creation functionality
	OrgManager interface {
		Show() (result *PivotOrg, err error)
		SafeCreate() (record *PivotOrg, err error)
	}
	orgManager struct {
		username string
		userGUID string
		log      *log.Logger
		tokens   oauth2.Tokens
		store    Persistence
		cfClient cloudfoundryclient.CloudFoundryClient
		apiInfo  map[string]interface{}
	}
)

//NewOrg - creates a new org manager
var NewOrg = func(username string, log *log.Logger, tokens oauth2.Tokens, store Persistence, authClient AuthRequestCreator) OrgManager {
	s := &orgManager{
		username: username,
		log:      log,
		tokens:   tokens,
		store:    store,
		cfClient: cloudfoundryclient.NewCloudFoundryClient(authClient, log),
	}
	return s
}

//Show - show if the user exists and already has an org or not.
func (s *orgManager) Show() (result *PivotOrg, err error) {
	result = new(PivotOrg)
	userInfo := GetUserInfo(s.tokens)
	NewUserMatch().UserInfo(userInfo).UserName(s.username).OnSuccess(func() {
		s.log.Println("getting userInfo: ", userInfo)
		s.log.Println("result value: ", result)
		err = s.store.FindOne(bson.M{EmailFieldName: s.username}, result)
		s.log.Println("response: ", result, err)
	}).OnFailure(func() {
		s.log.Println(ErrCantCallAcrossUsers.Error())
		err = ErrCantCallAcrossUsers
	}).Run()
	return
}

//SafeCreate - will create an org for the given user
func (s *orgManager) SafeCreate() (record *PivotOrg, err error) {
	var (
		orgName  = getOrgNameFromEmail(s.username)
		orgGUID  string
		userGUID string
	)
	s.cfClient.QueryAPIInfo()

	if userGUID, err = s.cfClient.QueryUserGUID(s.username); err == nil && userGUID != "" {
		s.log.Println("found user guid")

		if orgGUID, err = s.cfClient.AddOrg(orgName); err == nil {
			s.log.Println("we created the org successfully")

			if err = s.cfClient.AddRole(cloudfoundryclient.OrgEndpoint, orgGUID, cloudfoundryclient.RoleTypeManager, userGUID); err == nil {

				if err = s.cfClient.AddRole(cloudfoundryclient.OrgEndpoint, orgGUID, cloudfoundryclient.RoleTypeUser, userGUID); err == nil {
					record, err = s.upsert(orgGUID)
				}
			}
		}
	}

	if err != nil {
		s.log.Println("call to create org api failed")
		record = new(PivotOrg)
	}
	return
}

func (s *orgManager) upsert(orgGUID string) (record *PivotOrg, err error) {
	orgname := getOrgNameFromEmail(s.username)
	record = &PivotOrg{
		Email:   s.username,
		OrgName: orgname,
		OrgGUID: orgGUID,
	}
	err = s.store.Upsert(bson.M{EmailFieldName: s.username}, record)
	return
}

func getOrgNameFromEmail(email string) (orgName string) {
	username := strings.Split(email, "@")[0]
	orgName = fmt.Sprintf("pivot-%s", username)
	return
}
