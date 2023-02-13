package say

import (
	"github.com/tangx-labs/go-internal-demo/internal/student" // (1)
	// "github.com/tangx-labs/go-internal-demo/pkg/internal/master" // (2)
	"github.com/tangx-labs/go-internal-demo/pkg/hello" // (3)
)

func Say() {
	student.HelloStudent() // (1)

	// master.SayMaster() // (2)

	hello.Hello() // (3)
}
