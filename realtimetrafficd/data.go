/*
 * Copyright (C) 2014-2017 struktur AG
 * http://www.strukturag.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"encoding/json"
)

type interfacedata struct {
	data  interface{}
	iface string
}

func (i *interfacedata) set(iface string, data interface{}) {
	i.data = map[string]interface{}{
		iface:  data,
		"name": iface,
	}
	i.iface = iface
}

func (i *interfacedata) encode() ([]byte, error) {
	return json.Marshal(i.data)
}