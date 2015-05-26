package pezauth

import (
	"fmt"
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
)

//Constants to construct routes with
const (
	UserParam     = "user"
	APIVersion1   = "v1"
	AuthGroup     = "auth"
	OrgGroup      = "org"
	APIKeys       = "/api-keys"
	ValidKeyCheck = "/valid-key"
	StaticPath    = "public"
)

//formatted strings based on constants, to be used in URLs
var (
	APIKey        = fmt.Sprintf("/api-key/:%s", UserParam)
	OrgUser       = fmt.Sprintf("/user/:%s", UserParam)
	URLAuthBaseV1 = fmt.Sprintf("/%s/%s", APIVersion1, AuthGroup)
	URLOrgBaseV1  = fmt.Sprintf("/%s/%s", APIVersion1, OrgGroup)
)

//Response - generic response object
type Response struct {
	Payload  interface{}
	APIKey   string
	ErrorMsg string
}

//InitRoutes - initialize the mappings for controllers against valid routes
func InitRoutes(m *martini.ClassicMartini, redisConn Doer, mongoConn mongoCollection, authClient AuthRequestCreator) {
	setOauthConfig()
	keyGen := NewKeyGen(redisConn, &GUIDMake{})
	m.Use(render.Renderer())
	m.Use(martini.Static(StaticPath))
	m.Use(oauth2.Google(OauthConfig))
	authKey := NewAuthKeyV1(keyGen)

	m.Get("/info", authKey.Get())
	m.Get(ValidKeyCheck, NewValidateV1(keyGen).Get())

	m.Get("/me", oauth2.LoginRequired, DomainCheck, NewMeController().Get())

	m.Get("/", oauth2.LoginRequired, DomainCheck, func(params martini.Params, log *log.Logger, r render.Render, tokens oauth2.Tokens) {
		userInfo := GetUserInfo(tokens)
		r.HTML(SuccessStatus, "index", userInfo)
	})

	m.Group(URLAuthBaseV1, func(r martini.Router) {
		r.Put(APIKey, authKey.Put())
		r.Get(APIKey, authKey.Get())
		r.Delete(APIKey, authKey.Delete())
	}, oauth2.LoginRequired, DomainCheck)

	m.Group(URLOrgBaseV1, func(r martini.Router) {
		pcfOrg := NewOrgController(newMongoCollectionWrapper(mongoConn), authClient)
		r.Put(OrgUser, pcfOrg.Put())
		r.Get(OrgUser, pcfOrg.Get())
	}, oauth2.LoginRequired, DomainCheck)
}
