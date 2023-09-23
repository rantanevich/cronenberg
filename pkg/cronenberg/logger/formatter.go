package logger

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Formatter struct{}

func (formatter *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	levelText := strings.ToUpper(entry.Level.String())[0:4]

	entry.Message = strings.TrimSuffix(entry.Message, "\n")

	fmt.Fprintf(
		b,
		entryFormat,
		entry.Time.Format(time.RFC3339),
		entry.Data["context"],
		levelText,
		entry.Message,
	)

	return b.Bytes(), nil
}

var entryFormat = "[%s] (%s) %s: %s\n"

/*
Copyright 2019 Dennis Walters

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
