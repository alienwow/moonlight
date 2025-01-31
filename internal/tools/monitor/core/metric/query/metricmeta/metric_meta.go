// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metricmeta

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/go-redis/redis"

	"github.com/ping-cloudnative/moonlight-utils/providers/i18n"
	"github.com/ping-cloudnative/moonlight/proto-go/core/monitor/metric/pb"
)

type MetricMetaProvider interface {
	MetricMeta(langCodes i18n.LanguageCodes, i i18n.I18n, scope, scopeID string, names ...string) (map[string]*pb.MetricMeta, error)
}

func (m *Manager) getMetricMetaProviders() (list []MetricMetaProvider) {
	for _, fn := range m.metricProviders {
		list = append(list, fn())
	}
	return list
}

func (m *Manager) MetricNames(langCodes i18n.LanguageCodes, scope, scopeID string) (names []*pb.NameDefine, err error) {
	metrics, err := m.MetricMeta(langCodes, scope, scopeID)
	if err != nil {
		return nil, err
	}
	for _, m := range metrics {
		names = append(names, m.Name)
	}
	return names, nil
}

func (m *Manager) MetricMeta(langCodes i18n.LanguageCodes, scope, scopeID string, names ...string) ([]*pb.MetricMeta, error) {
	metricMetas, err := m.getMetricMeta(langCodes, scope, scopeID, names...)
	if err != nil {
		return nil, err
	}
	var list []*pb.MetricMeta
	for _, item := range metricMetas {
		list = append(list, item)
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].Name.Name == list[j].Name.Name {
			return list[i].Name.Key == list[j].Name.Key
		}
		return list[i].Name.Name < list[j].Name.Name
	})
	return list, nil
}

func (m *Manager) GetSingleMetricsMeta(langCodes i18n.LanguageCodes, scope string, scopeID string, metric string) (*pb.MetricMeta, error) {
	metricMetas, err := m.getMetricMeta(langCodes, scope, scopeID, metric)
	if err != nil {
		return nil, err
	}
	metricMeta, ok := metricMetas[metric]
	if !ok {
		return nil, fmt.Errorf("not found metric %q", metric)
	}
	return metricMeta, nil
}

func (m *Manager) GetMetricMetaByCache(scope, scopeID string, names ...string) ([]*pb.MetricMeta, error) {
	var result []*pb.MetricMeta
	for _, metric := range names {
		key := fmt.Sprintf("metric_meta_%s_%s_%s", scope, scopeID, metric)
		meta, err := m.redis.Get(key).Result()

		if err != nil && err != redis.Nil {
			return nil, err
		}
		var r []*pb.MetricMeta
		if len(meta) > 0 {
			err := json.NewDecoder(strings.NewReader(meta)).Decode(&r)
			if err != nil {
				return nil, err
			}
			if len(r) > 0 {
				result = append(result, r...)
				continue
			}
		}
		metricMetas, err := m.getMetricMeta(nil, scope, scopeID, metric)
		if err != nil {
			return nil, err
		}
		for _, item := range metricMetas {
			r = append(r, item)
		}
		if len(r) <= 0 {
			continue
		}

		sb := &strings.Builder{}
		err = json.NewEncoder(sb).Encode(r)
		if err != nil {
			return nil, err
		}

		_, err = m.redis.Set(key, sb.String(), m.metricMetaCacheExpiration).Result()
		if err != nil {
			return nil, err
		}
		result = append(result, r...)
	}

	return result, nil
}

func (m *Manager) getMetricMeta(langCodes i18n.LanguageCodes, scope, scopeID string, names ...string) (map[string]*pb.MetricMeta, error) {
	mp := m.getMetricMetaProviders()
	ms := make(map[string]*pb.MetricMeta)
	for _, p := range mp {
		m, err := p.MetricMeta(langCodes, m.i18n, scope, scopeID, names...)
		if err != nil {
			return nil, err
		}
		ms = appendMetricMeta(ms, m)
	}
	return ms, nil
}

func appendMetricMeta(metric1, metric2 map[string]*pb.MetricMeta) map[string]*pb.MetricMeta {
	for n1, m1 := range metric1 {
		if m2, ok := metric2[n1]; ok {
			m1.Name = m2.Name
			m1.Labels = appendLabels(m1.Labels, m2.Labels)
			m1.Fields = appendFields(m1.Fields, m2.Fields)
			m1.Tags = appendTags(m1.Tags, m2.Tags)
		}
	}
	for n2, m2 := range metric2 {
		if _, ok := metric1[n2]; !ok {
			metric1[n2] = m2
		}
	}
	return metric1
}

func appendTags(a, b map[string]*pb.TagDefine) map[string]*pb.TagDefine {
	if a == nil {
		return b
	}
	if b != nil {
		for k, v := range b {
			// if _, ok := a[k]; ok {
			// 	continue
			// }
			a[k] = v
		}
	}
	return a
}

func appendFields(a, b map[string]*pb.FieldDefine) map[string]*pb.FieldDefine {
	if a == nil {
		return b
	}
	if b != nil {
		for k, v := range b {
			// if _, ok := a[k]; ok {
			// 	continue
			// }
			a[k] = v
		}
	}
	return a
}

func appendLabels(a, b map[string]string) map[string]string {
	if a == nil {
		return b
	}
	if b != nil {
		for k, v := range b {
			// if _, ok := a[k]; ok {
			// 	continue
			// }
			a[k] = v
		}
	}
	return a
}

func transMetricMetas(langCodes i18n.LanguageCodes, i i18n.I18n, metas map[string]*pb.MetricMeta) map[string]*pb.MetricMeta {
	for _, item := range metas {
		t := i.Translator(item.Name.Key)
		item.Name.Name = t.Text(langCodes, item.Name.Name)
		for _, f := range item.Fields {
			f.Name = t.Text(langCodes, f.Name)
			for _, v := range f.Values {
				v.Name = t.Text(langCodes, v.Name)
			}
		}
		for _, tag := range item.Tags {
			tag.Name = t.Text(langCodes, tag.Name)
			for _, v := range tag.Values {
				v.Name = t.Text(langCodes, v.Name)
			}
		}
	}
	return metas
}

func (m *Manager) RegeistMetricMeta(scope, scopeID, group string, metrics ...*pb.MetricMeta) error {
	return m.regeistMetricMeta(scope, scopeID, group, metrics...)
}

func (m *Manager) UnregeistMetricMeta(scope, scopeID, group string, metrics ...string) error {
	return m.unregeistMetricMeta(scope, scopeID, group, metrics...)
}
