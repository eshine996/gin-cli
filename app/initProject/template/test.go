package template

var TestServiceUserTestTmp = `package service

import (
	"context"
	"fmt"
	"{{projectName}}/service"
	"testing"
)

func TestUser(t *testing.T) {
	out, err := service.User().GetUserByUserName(context.TODO(), "admin")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(out)
}`
