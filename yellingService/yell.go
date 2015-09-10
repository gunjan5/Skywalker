package yellingService

/* Copyright (C) 2015  Gunjan Patel
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation,
 * either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
 * the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details. You should have received a copy of
 * the GNU Affero General Public License along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import (
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

// YellingService provides operations on strings.
type YellingServiceI interface {
	Yell(string) error
}

type YellingService struct{}

func (YellingService) Yell(s string) error {
	if s == "" {
		cmd := exec.Command("say", "It's an empty string dummy!")
		err := cmd.Run()
		check(err)
		return ErrEmpty
	}
	cmd := exec.Command("say", s)
	err := cmd.Run()

	return check(err)
}

func MakeYellEndpoint(svc YellingServiceI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(yellRequest)
		e := svc.Yell(req.S)
		return yellResponse{e.Error()}, nil
	}
}

func DecodeYellRequest(r *http.Request) (interface{}, error) {
	var request yellRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type yellRequest struct {
	S string `json:"s"`
}

type yellResponse struct {
	Err string `json:"err,omitempty"`
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
