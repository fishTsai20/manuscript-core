// Copyright 2024 The Chainbase. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"manuscript-core/client"
	"manuscript-core/commands"
	"os"
)

func main() {
	log.SetFlags(0)
	err := client.Init()
	if err != nil {
		log.Fatalf("Error starting common API: %v", err)
	}
	defer client.Shutdown()
	err = commands.Execute(os.Args[1:])
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
