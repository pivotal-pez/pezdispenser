package pezauth

import (
	"fmt"
	"log"
	"strings"

	"github.com/martini-contrib/oauth2"
	"github.com/pivotalservices/pezdispenser/cloudfoundryclient"
	"github.com/xchapter7x/goutil"
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

func (s *orgManager) RollbackCreate() (err error) {
	s.log.Println("rolling back changes")
	return
}

//SafeCreate - will create an org for the given user
func (s *orgManager) SafeCreate() (record *PivotOrg, err error) {
	var (
		orgName  = getOrgNameFromEmail(s.username)
		orgGUID  string
		userGUID string
	)
	c := goutil.NewChain(nil)
	c.Call(s.cfClient.QueryAPIInfo)
	c.CallP(c.Returns(&userGUID, &err), s.cfClient.QueryUserGUID, s.username)
	c.CallP(c.Returns(&orgGUID, &err), s.cfClient.AddOrg, orgName)
	c.Call(s.cfClient.AddRole, cloudfoundryclient.OrgEndpoint, orgGUID, cloudfoundryclient.RoleTypeManager, userGUID)
	c.Call(s.cfClient.AddRole, cloudfoundryclient.OrgEndpoint, orgGUID, cloudfoundryclient.RoleTypeUser, userGUID)
	c.CallP(c.Returns(record, &err), s.upsert, orgGUID)

	if c.Error != nil {
		s.log.Println("we experienced a failure, should roll back changes", c.Error)
		err = c.Error
		record = new(PivotOrg)
		s.RollbackCreate()
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
