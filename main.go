package main

/*
   Copyright (C) 2015  Gunjan Patel

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import (
	_ "encoding/json"
	_ "errors"
	"log"
	"net/http"
	_ "strings"

	"golang.org/x/net/context"

	_ "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	yell "github.com/gunjan5/Skywalker/yellingService"
)

func main() {
	ctx := context.Background()
	svc := stringService{}

	countHandler := httptransport.Server{
		Context:            ctx,
		Endpoint:           yell.makeCountEndpoint(svc),
		DecodeRequestFunc:  yell.decodeCountRequest,
		EncodeResponseFunc: yell.encodeResponse,
	}

	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
