// START 2OMIT
package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/nomad/api"
)

func main() {
	c, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	eval, _, err := c.Jobs().Register(testJob(), nil)
	if err != nil {
		log.Fatal(err)
	}
	if eval == "" {
		log.Fatalf("missing eval id")
	}
}

// END 2OMIT

// START 1OMIT
func testJob() *api.Job {
	task := api.NewTask("task1", "docker").
		SetConfig("image", "some/container_image").
		Require(&api.Resources{CPU: 100, MemoryMB: 256, IOPS: 0}).
		SetLogConfig(&api.LogConfig{MaxFiles: 1, MaxFileSizeMB: 2})

	task.Env = map[string]string{"MY_KEY": "1234"}

	group := api.NewTaskGroup("group1", 1).
		AddTask(task).
		RequireDisk(&api.EphemeralDisk{
			SizeMB: 25,
		})

	job := api.NewBatchJob("job1", "test", "region1", 1).
		AddDatacenter("dc1").
		AddTaskGroup(group)

	return job
}

// END 1OMIT
