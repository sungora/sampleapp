package sample

import (
	"github.com/sungora/app/lg"
	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/workflow"

	"sample/modules/sample/apiv1"
	"sample/modules/sample/config"
	"sample/modules/sample/page"
	"sample/modules/sample/worker/workfour"
	"sample/modules/sample/worker/workone"
	"sample/modules/sample/worker/worktwo"
)

const ModName string = "sample"

func Init() (code int) {

	// config
	if 0 < config.Init(ModName) {
		return
	}

	// router
	servhttp.MountRoutes("/", page.Routes)
	servhttp.MountRoutes("/api/v1", apiv1.Routes)

	// workers
	workflow.TaskAddCron(&workone.SampleTaskOne{})
	workflow.TaskAddCron(&worktwo.SampleTaskTwo{})
	workflow.TaskAddCron(&workfour.SampleTaskFour{})

	// log
	lg.SetMessages(map[int]string{
		1000: "Message format Fmt from 1000",
		1001: "Message format Fmt from 1001",
		1002: "Message format Fmt from 1002",
		1003: "Message format Fmt from 1003",
		1004: "Message format Fmt from 1004",
		1005: "Message format Fmt from 1005",
	})

	return
}
