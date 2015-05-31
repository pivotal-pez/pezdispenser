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
		userGUID string
	)
	record = new(PivotOrg)

	if _, err = s.cfClient.QueryAPIInfo(); err == nil {

		if userGUID, err = s.cfClient.QueryUserGUID(s.username); err != nil {
			s.log.Println("QueryUserGUID failed w/ error: ", err)
			err = ErrCouldNotGetUserGUID
		} else {
			record, err = s.runOrgCreateCallchain(userGUID)
		}
	}
	return
}

//RollbackCreate - will rollback anything created by a failed SafeCreate call, so we dont have orphaned/incomplete records
func (s *orgManager) RollbackCreate(orgGUID string) (err error) {
	s.log.Println("rolling back changes")
	if err = s.store.Remove(bson.M{EmailFieldName: s.username}); err == nil {

		if err = s.cfClient.RemoveOrg(orgGUID); err == nil {
			s.log.Println("org at guid deleted: ", orgGUID)

		} else {
			s.log.Println("org at guid could not be deleted: ", orgGUID, err)
		}

	} else {
		s.log.Println("rollback mongodb error: ", err.Error())
	}
	return
}

func (s *orgManager) runOrgCreateCallchain(userGUID string) (record *PivotOrg, err error) {
	var (
		orgName = getOrgNameFromEmail(s.username)
		orgGUID string
	)
	record = new(PivotOrg)
	c := goutil.NewChain(nil)
	c.CallP(c.Returns(&orgGUID, &err), s.cfClient.AddOrg, orgName)
	c.Call(s.cfClient.AddRole, cloudfoundryclient.OrgEndpoint, orgGUID, cloudfoundryclient.RoleTypeManager, userGUID)
	c.Call(s.cfClient.AddRole, cloudfoundryclient.OrgEndpoint, orgGUID, cloudfoundryclient.RoleTypeUser, userGUID)
	c.Call(s.cfClient.AddSpace, DefaultSpaceName, orgGUID)

	if record, err = s.upsert(orgGUID); err != nil {
		c.Error = err
	}

	if c.Error != nil {
		s.log.Println("we experienced a failure, should roll back changes", c.Error)
		err = c.Error

		if orgGUID != "" {
			s.RollbackCreate(orgGUID)

		} else {
			s.log.Println("nothing to rollback")
		}
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
