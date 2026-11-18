// Copyright (c) 2026 malasahjagotwin. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cookiejar_test

import "github.com/malasahjagotwin/sawithttp/cookiejar"

type dummypsl struct {
	List cookiejar.PublicSuffixList
}

func (dummypsl) PublicSuffix(domain string) string {
	return domain
}

func (dummypsl) String() string {
	return "dummy"
}

var publicsuffix = dummypsl{}
