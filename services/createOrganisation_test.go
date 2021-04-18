package services

import (
	"context"
	"fmt"
	createOrganisationService "github.com/myuser/myrepo/proto/helloworld"
	"testing"
)

func TestServer2_createOrganisation(t *testing.T) {
	server:=Server2{}
	res,err:=server.createOrganisation(context.Background(),&createOrganisationService.OrganisationRequest{
		Name:    "Amity University",
		Desc:    "A college",
		Email:   "amity.edu@gmail.com",
		Contact: "9989898989",
		Image:   "wqwewewew",
		Video:   "wewewwewe",
	})
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(res)
}
