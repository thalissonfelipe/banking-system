package account

import (
	"os"
	"testing"

	"github.com/thalissonfelipe/banking/pkg/tests/dockertest"
)

var pgDocker *dockertest.PostgresDocker

func TestMain(m *testing.M) {
	pgDocker = dockertest.SetupTest("../migrations")

	exitCode := m.Run()

	defer dockertest.RemoveContainer(pgDocker)

	os.Exit(exitCode)
}