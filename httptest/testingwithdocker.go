import (
	"log"
	"os"
	"testing"
)

// START OMIT

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")

	container, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})

	if err := pool.Retry(func() error {
		//...
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(container); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

// END OMIT
