// Copyright (c) 2026 malasahjagotwin. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !go1.11
// +build !go1.11

package http2

import (
	"net/textproto"

	"github.com/malasahjagotwin/sawithttp/httptrace"
)

func traceHasWroteHeaderField(trace *httptrace.ClientTrace) bool { return false }

func traceWroteHeaderField(trace *httptrace.ClientTrace, k, v string) {}

func traceGot1xxResponseFunc(trace *httptrace.ClientTrace) func(int, textproto.MIMEHeader) error {
	return nil
}
