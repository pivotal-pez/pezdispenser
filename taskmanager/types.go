package taskmanager

import (
	"time"

	"github.com/pivotal-pez/pezdispenser/service/_integrations"
	"labix.org/v2/mgo/bson"
)

type (
	//Task - a task object
	Task struct {
		ID         bson.ObjectId          `bson:"_id"`
		Timestamp  time.Time              `bson:"timestamp"`
		Status     string                 `bson:"status"`
		Profile    ProfileType            `bson:"profile"`
		CallerName string                 `bson:"caller_name"`
		MetaData   map[string]interface{} `bson:"metadata"`
		Lock       bool                   `bson:"lock"`
	}

	//TaskManager - manages task interactions crud stuff
	TaskManager struct {
		taskCollection integrations.Collection
	}

	//ProfileType - indicator of the purpose of the task to be performed
	ProfileType string
)
