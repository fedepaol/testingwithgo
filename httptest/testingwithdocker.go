import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
)

// START OMIT
var db *sql.DB

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	// pulls an image, creates a container based on it and runs it
	container, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
	// exponential backoff-retry, because the application in the container might
	// not be ready to accept connections yet
	if err := pool.Retry(func() error {
		//...
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql", container.GetPort("3306/tcp")))
		//...
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	//run the tests
	code := m.Run()
	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(container); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

// END OMIT
