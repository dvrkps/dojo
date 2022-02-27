package mock

import (
	"testing"

	"github.com/dvrkps/dojo/ddd/user"
)

func TestServiceInterfaceImplementation(t *testing.T) {
	var i interface{} = new(Service)
	_, ok := i.(user.Service)
	if !ok {
		t.Fatalf("%T not implement interface", i)
	}
}
