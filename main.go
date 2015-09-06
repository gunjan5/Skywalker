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
	"log"
	"net/http"

	"golang.org/x/net/context"

	httptransport "github.com/go-kit/kit/transport/http"
	//yell "github.com/gunjan5/Skywalker/yellingService"
	yell "yellingService"
)

// func check(e error) error {
// 	if e != nil {
// 		fmt.Println(e)
// 		return e
// 	}
// 	return nil
// }

func main() {
	ctx := context.Background()
	svc := yell.YellingService{}

	yellHandler := httptransport.Server{
		Context:            ctx,
		Endpoint:           yell.MakeYellEndpoint(svc),
		DecodeRequestFunc:  yell.DecodeYellRequest,
		EncodeResponseFunc: yell.EncodeResponse,
	}

	// cmd := exec.Command("say", "bloody blood sucking (blood) blaste${}rs")
	// err := cmd.Run()
	// fmt.Println(check(err))

	http.Handle("/yell", yellHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
