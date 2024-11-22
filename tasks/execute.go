package tasks

import (
	"fmt"
	"github.com/vmihailenco/taskq/v3"
	"github.com/xnok/tasq-renovate/commands"
	"log/slog"
	"os/exec"
)

var ExecuteTask = taskq.RegisterTask(&taskq.TaskOptions{
	Name:    "execute",
	Handler: ExecuteTaskHandler,
})

var ShellExecuteCommandFunc = func(repo string) commands.Executor {
	return exec.Command("renovate",
		"--autodiscover=false",
		repo,
	)
}

func ExecuteTaskHandler(msg *taskq.Message) error {
	repo, ok := msg.Args[0].(string)
	if !ok {
		return fmt.Errorf("invalid repo argument")
	}

	slog.Info("Running execute task...",
		slog.String("repository", repo))

	cmd := ShellExecuteCommandFunc(repo)

	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("execute task failed: %v", err)
	}

	return nil
}
