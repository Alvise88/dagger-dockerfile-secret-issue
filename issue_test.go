package issue

import (
	"context"
	"testing"

	"dagger.io/dagger"
	"dagger.io/dagger/dag"
)

func TestTrySecret(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer dag.Close()

	sshKey := "white"

	_, err := dag.Container().Build(dag.Host().Directory("./data"), dagger.ContainerBuildOpts{
		Secrets: []*dagger.Secret{dag.SetSecret("ssh-key", sshKey)},
	}).Sync(ctx)

	if err != nil {
		t.Error(err)
	}
}
