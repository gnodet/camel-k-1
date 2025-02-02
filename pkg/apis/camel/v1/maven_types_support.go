/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"encoding/xml"
)

func (in *MavenArtifact) GetDependencyID() string {
	mvn := "mvn:" + in.GroupID + ":" + in.ArtifactID
	if in.Classifier != "" {
		if in.Version != "" {
			mvn = mvn + ":" + in.Type + ":" + in.Version + ":" + in.Classifier
		} else {
			mvn = mvn + ":" + in.Type + "::" + in.Classifier
		}
	} else {
		if in.Type != "" {
			mvn = mvn + ":" + in.Type
		}
		if in.Version != "" {
			mvn = mvn + ":" + in.Version
		}
	}
	return mvn
}

//nolint:musttag // the name of the xml is dynamic
type propertiesEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (m Properties) AddAll(properties map[string]string) {
	for k, v := range properties {
		m.Add(k, v)
	}
}

func (m Properties) Add(key string, value string) {
	m[key] = value
}

func (m Properties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		//nolint:musttag
		if err := e.Encode(propertiesEntry{XMLName: xml.Name{Local: k}, Value: v}); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

func (m PluginProperties) AddAll(properties map[string]string) {
	for k, v := range properties {
		m.Add(k, v)
	}
}

func (m PluginProperties) Add(key string, value string) {
	m[key] = StringOrProperties{Value: value}
}

func (m PluginProperties) AddProperties(key string, properties map[string]string) {
	m[key] = StringOrProperties{Properties: properties}
}

func (m PluginProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for k, v := range m {
		if v.Value != "" {
			//nolint:musttag
			if err := e.Encode(propertiesEntry{XMLName: xml.Name{Local: k}, Value: v.Value}); err != nil {
				return err
			}
		}

		if len(v.Properties) > 0 {
			nestedPropertyStart := xml.StartElement{Name: xml.Name{Local: k}}
			if err := e.EncodeToken(nestedPropertyStart); err != nil {
				return err
			}

			for key, value := range v.Properties {
				//nolint:musttag
				if err := e.Encode(propertiesEntry{XMLName: xml.Name{Local: key}, Value: value}); err != nil {
					return err
				}
			}

			if err := e.EncodeToken(nestedPropertyStart.End()); err != nil {
				return err
			}

		}
	}

	return e.EncodeToken(start.End())
}
