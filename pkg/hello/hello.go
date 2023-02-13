package hello

import (
	"github.com/tangx-labs/go-internal-demo/internal/student"
	"github.com/tangx-labs/go-internal-demo/pkg/internal/master"
)

func Hello() {
	student.HelloStudent()

	master.HelloMaster()
}
