package concurrent

import "testing"

var commandTests = []struct {
	command string
	err     error
}{
	{
		"C://work//src//github.com//divyag9//encryptbackup//encryptbackup -sd D://Users//dmuppaneni//Documents//source -td D://Users//dmuppaneni//Documents//target -sgpkey D://Users//dmuppaneni//Documents//pubring.gpg -midkey D://Users//dmuppaneni//Documents//pubring.gpg",
		nil,
	},
}

func TestExecuteCommand(t *testing.T) {
	for _, commandTests := range commandTests {
		_, err := ExecuteCommand(commandTests.command)
		if err != commandTests.err {
			t.Errorf("Returned: %v. Expected: %v", err, commandTests.err)
		}
	}
}
