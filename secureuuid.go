/*
 * Copyright (c) 2024 flowerinsnow
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

/*
 * This file uses the library google/uuid (https://github.com/google/uuid)
 * BSD 3-Clause "New" or "Revised" License
 * Copyright 2009 The Go Authors.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *    * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *    * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *    * Neither the name of Google LLC nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/*
 * This file uses the library google/uuid (https://github.com/google/uuid)
 * BSD 3-Clause "New" or "Revised" License
 * Copyright (c) 2009,2014 Google Inc. All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *   - Redistributions of source code must retain the above copyright
 *
 * notice, this list of conditions and the following disclaimer.
 *   - Redistributions in binary form must reproduce the above
 *
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *   - Neither the name of Google Inc. nor the names of its
 *
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */
package main

import (
	"flag"
	"fmt"
	"github.com/flowerinsnowdh/secureuuid/common"
	"github.com/google/uuid"
	"os"
	"strings"
)

var (
	versionFlag    bool
	licenseFlag    bool
	randomFlag     bool
	uuidValue      string
	compactFlag    bool
	standardFlag   bool
	noNewLine      bool
	namespaceValue string
	md5Value       string
	sha1Value      string
)

func init() {
	flag.BoolVar(&versionFlag, "v", false, "Print secureuuid version")
	flag.BoolVar(&versionFlag, "version", false, "Print secureuuid version")
	flag.BoolVar(&licenseFlag, "l", false, "Print license information")
	flag.BoolVar(&licenseFlag, "license", false, "Print license information")
	flag.BoolVar(&randomFlag, "r", false, "Use a random UUIDv4")
	flag.BoolVar(&randomFlag, "random", false, "Use a random UUIDv4")
	flag.StringVar(&uuidValue, "u", "", "Use the specified UUID")
	flag.StringVar(&uuidValue, "uuid", "", "Use the specified UUID")
	flag.StringVar(&namespaceValue, "n", "", "Use the specified UUID namespace")
	flag.StringVar(&namespaceValue, "namespace", "", "Use the specified namespace")
	flag.StringVar(&md5Value, "md5", "", "Use the specified UUIDv3 name")
	flag.StringVar(&sha1Value, "sha1", "", "Use the specified UUIDv5 name")
	flag.BoolVar(&compactFlag, "c", false, "Print compact UUID (xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx)")
	flag.BoolVar(&compactFlag, "compact", false, "Print compact UUID (xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx)")
	flag.BoolVar(&standardFlag, "s", false, "Print standard UUID (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx))")
	flag.BoolVar(&standardFlag, "standard", false, "Print standard UUID (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx))")
	flag.BoolVar(&noNewLine, "no-newline", false, "Print without newline")
}

func main() {
	flag.Parse()

	if versionFlag {
		fmt.Println("secureuuid v" + common.VERSION)
		return
	}
	if licenseFlag {
		fmt.Println(common.NOTICE)
		return
	}

	if common.CountTrue(
		randomFlag,
		uuidValue != "",
		md5Value != "",
		sha1Value != "") != 1 {
		_, _ = fmt.Fprintln(os.Stderr, "At least one and no more than one of -random, -uuid, -md5, or -sha1 must be specified.")
		os.Exit(1)
	}

	if (md5Value != "" || sha1Value != "") && namespaceValue == "" {
		_, _ = fmt.Fprintln(os.Stderr, "-namespace must be specified if -md5 or -sha1 is specified.")
		os.Exit(1)
	}

	var namespace uuid.UUID
	if namespaceValue != "" {
		var err error
		if strings.EqualFold(namespaceValue, "dns") {
			namespace, _ = uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
		} else if strings.EqualFold(namespaceValue, "url") {
			namespace, _ = uuid.Parse("6ba7b811-9dad-11d1-80b4-00c04fd430c8")
		} else if strings.EqualFold(namespaceValue, "oid") {
			namespace, _ = uuid.Parse("6ba7b812-9dad-11d1-80b4-00c04fd430c8")
		} else if strings.EqualFold(namespaceValue, "x500") {
			namespace, _ = uuid.Parse("6ba7b814-9dad-11d1-80b4-00c04fd430c8")
		} else if namespace, err = uuid.Parse(namespaceValue); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Invalid namespace: "+err.Error())
			os.Exit(124)
		}
	}

	var u uuid.UUID

	if randomFlag {
		u = uuid.New() // new UUIDv4
	} else if md5Value != "" {
		u = uuid.NewMD5(namespace, []byte(md5Value))
	} else if sha1Value != "" {
		u = uuid.NewSHA1(namespace, []byte(sha1Value))
	} else if err := uuid.Validate(uuidValue); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Invalid UUID: "+err.Error())
		os.Exit(125)
	} else {
		u, _ = uuid.Parse(uuidValue)
	}

	var result = u.String()

	if compactFlag {
		fmt.Print(strings.ReplaceAll(result, "-", ""))
	}

	if standardFlag {
		fmt.Print(result)
	}

	if !noNewLine {
		fmt.Println()
	}
}
