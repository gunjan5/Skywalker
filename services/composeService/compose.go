package composeService

/* Copyright (C) 2015  Gunjan Patel
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation,
 * either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
 * the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details. You should have received a copy of
 * the GNU Affero General Public License along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import (
	_ "log"

	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/project"

	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/exec"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
)

func check(e error) error {
	if e != nil {
		fmt.Println(e)
		return e
	}
	return nil
}

// ComposeService provides operations on strings.
type ComposeServiceI interface {
	Compose(string) error
}

type ComposeService struct{}

func (ComposeService) Compose(s string) error {
	project, err := docker.NewProject(&docker.Context{
		Context: project.Context{
			ComposeFile: "docker-compose.yml",
			ProjectName: "my-compose",
		},
	})

	check(err)
	fmt.Println(s)

	project.Up()

	return nil
}

func MakeComposeEndpoint(svc ComposeServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(composeRequest)
		e := svc.Compose(req.S)
		if e != nil {
			return composeResponse{e.Error()}, nil
		}
		return composeResponse{""}, nil
	}
}

func DecodeComposeRequest(r *http.Request) (interface{}, error) {
	var request composeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type composeRequest struct {
	S string `json:"s"`
}

type composeResponse struct {
	Err string `json:"err,omitempty"`
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
