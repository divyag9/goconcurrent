package concurrent

import (
	"errors"
	"testing"
)

var commandTests = []struct {
	command string
	err     error
}{
	{
		"C://work//src//github.com//divyag9//encryptbackup//encryptbackup -sd D://Users//dmuppaneni//Documents//source -td D://Users//dmuppaneni//Documents//target -sgpkey D://Users//dmuppaneni//Documents//pubring.gpg -midkey D://Users//dmuppaneni//Documents//pubring.gpg",
		nil,
	},
	{
		"C://work//src//github.com//divyag9//encryptbackup//encrypt -sd D://Users//dmuppaneni//Documents//source -td D://Users//dmuppaneni//Documents//target -sgpkey D://Users//dmuppaneni//Documents//pubring.gpg -midkey D://Users//dmuppaneni//Documents//pubring.gpg",
		errors.New("exec: \"C://work//src//github.com//divyag9//encryptbackup//encrypt\": file does not exist"),
	},
}

func TestExecuteCommand(t *testing.T) {
	for _, commandTests := range commandTests {
		_, err := ExecuteCommand(commandTests.command)
		if err != commandTests.err && err.Error() != commandTests.err.Error() {
			t.Errorf("Returned: %v. Expected: %v", err, commandTests.err)
		}
	}
}
