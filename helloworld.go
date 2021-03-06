//
// Sample hello world in go that has base36 version string
//
//
// Copyright (C) 2016 Ryan A. Chapman. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
//   1. Redistributions of source code must retain the above copyright notice,
//      this list of conditions and the following disclaimer.
//
//   2. Redistributions in binary form must reproduce the above copyright notice,
//      this list of conditions and the following disclaimer in the documentation
//      and/or other materials provided with the distribution.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
// INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
// FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE AUTHORS
// OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
// EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
// PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
// OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
// OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
// ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package main

import (
    "fmt"
    "strconv"
    "strings"
    "time"
)

const VERSION = "1.0"

func versionString() string {
    buildNum := strings.ToUpper(strconv.FormatInt(BUILDTIMESTAMP, 36))
    buildDate := time.Unix(BUILDTIMESTAMP, 0).Format(time.UnixDate)
    v := fmt.Sprintf("helloworld %s (build %v, %v by %v@%v)", VERSION, buildNum, buildDate, BUILDUSER, BUILDHOST)
    return v
}


func main() {
	fmt.Printf("%s\n\n", versionString())
	fmt.Printf("Hello World\n")
}
