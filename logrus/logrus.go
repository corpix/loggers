package logrus

// The MIT License (MIT)
//
// Copyright © 2017 Dmitry Moskowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"github.com/sirupsen/logrus"
)

// Logrus is a logrus for logger that implements
// io.Writer interface.
type Logrus struct {
	*logrus.Logger
}

// Write slice of bytes into the logger and return number of written
// bytes and error value of present.
func (l *Logrus) Write(buf []byte) (int, error) {
	n := len(buf) + 1
	l.Printf("%s\n", buf)
	return n, nil
}

// Level returns a current logger level number.
func (l *Logrus) Level() interface{} {
	return l.Logger.Level
}

// New wraps logrus logger with binding.
func New(l *logrus.Logger) *Logrus {
	return &Logrus{l}
}
