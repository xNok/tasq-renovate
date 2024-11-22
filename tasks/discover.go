package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vmihailenco/taskq/v3"
	"github.com/xnok/tasq-renovate/broker"
	"github.com/xnok/tasq-renovate/commands"
	"log"
	"log/slog"
	"os"
	"os/exec"
)

var DiscoverTask = taskq.RegisterTask(&taskq.TaskOptions{
	Name:    "discover",
	Handler: DiscoveryTaskHandler,
})

var ShellDiscoverCommandFunc = func(file string) commands.Executor {
	return exec.Command("renovate",
		"--autodiscover",
		fmt.Sprintf("--write-discovered-repos=%s", file))
}

func DiscoveryTaskHandler(ctx context.Context) error {
	slog.Info("Running discover task...")

	// Temp file for the discovery output
	file, err := os.CreateTemp("", "prefix")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Run Renovate CLI with writeDiscoveredRepos
	cmd := ShellDiscoverCommandFunc(file.Name())

	_, err = cmd.Output()
	if err != nil {
		return err
	}

	// Process the output JSON file
	data, err := os.ReadFile(file.Name())
	if err != nil {
		return fmt.Errorf("error reading discovered repos file: %v", err)
	}

	var discoveredRepos []string
	if err := json.Unmarshal(data, &discoveredRepos); err != nil {
		return fmt.Errorf("error unmarshaling discovered repos: %v", err)
	}

	// Create "execute" tasks for each discovered repo
	for _, repo := range discoveredRepos {
		err = broker.MainQueue.Add(ExecuteTask.WithArgs(ctx, repo))
		if err != nil {
			slog.Error("Failed to register execute task for discovered repo",
				slog.String("repository", repo),
				slog.String("err", err.Error()),
			)
		}
	}

	return nil
}
