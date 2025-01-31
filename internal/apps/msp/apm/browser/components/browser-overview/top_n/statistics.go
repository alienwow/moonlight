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

package top_n

import (
	"fmt"
	"strconv"

	"google.golang.org/protobuf/types/known/structpb"

	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/components/topn"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cptype"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	"github.com/ping-cloudnative/moonlight/pkg/common/errors"
	"github.com/ping-cloudnative/moonlight/pkg/math"
	"github.com/ping-cloudnative/moonlight/pkg/time"
	metricpb "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/metric/pb"
)

func (p *provider) maxReqDomainTop5(sdk *cptype.SDK) ([]topn.Item, error) {
	statement := fmt.Sprintf("SELECT host::tag, count(host::tag) " +
		"FROM ta_timing " +
		"WHERE tk::tag=$terminus_key " +
		"GROUP BY host::tag " +
		"ORDER BY count(host::tag) DESC " +
		"LIMIT 5")
	params := map[string]*structpb.Value{
		"terminus_key": structpb.NewStringValue(p.InParamsPtr.TenantId),
	}

	request := &metricpb.QueryWithInfluxFormatRequest{
		Start:     strconv.FormatInt(p.InParamsPtr.StartTime, 10),
		End:       strconv.FormatInt(p.InParamsPtr.EndTime, 10),
		Statement: statement,
		Params:    params,
	}

	ctx := apis.GetContext(sdk.Ctx, func(header *transport.Header) {
		header.Set("terminus_key", p.InParamsPtr.TenantId)
	})

	response, err := p.Metric.QueryWithInfluxFormat(ctx, request)
	if err != nil {
		return nil, errors.NewInternalServerError(err)
	}
	rows := response.Results[0].Series[0].Rows
	if rows == nil || len(rows) == 0 {
		return nil, nil
	}
	var items []topn.Item
	total := math.DecimalPlacesWithDigitsNumber(rows[0].Values[1].GetNumberValue(), 2)
	for _, row := range rows {
		item := topn.Item{
			ID:    row.Values[0].GetStringValue(),
			Name:  row.Values[0].GetStringValue(),
			Value: math.DecimalPlacesWithDigitsNumber(row.Values[1].GetNumberValue(), 2),
			Unit:  "",
		}
		if item.Value == 0 {
			continue
		}
		item.Percent = math.DecimalPlacesWithDigitsNumber(item.Value/total*1e2, 2)
		items = append(items, item)
	}
	return items, nil
}

func (p *provider) maxReqPageTop5(sdk *cptype.SDK) ([]topn.Item, error) {
	statement := fmt.Sprintf("SELECT doc_path::tag, count(doc_path::tag) " +
		"FROM ta_timing " +
		"WHERE tk::tag=$terminus_key " +
		"GROUP BY doc_path::tag " +
		"ORDER BY count(doc_path::tag) DESC " +
		"LIMIT 5")
	params := map[string]*structpb.Value{
		"terminus_key": structpb.NewStringValue(p.InParamsPtr.TenantId),
	}

	request := &metricpb.QueryWithInfluxFormatRequest{
		Start:     strconv.FormatInt(p.InParamsPtr.StartTime, 10),
		End:       strconv.FormatInt(p.InParamsPtr.EndTime, 10),
		Statement: statement,
		Params:    params,
	}

	ctx := apis.GetContext(sdk.Ctx, func(header *transport.Header) {
		header.Set("terminus_key", p.InParamsPtr.TenantId)
	})

	response, err := p.Metric.QueryWithInfluxFormat(ctx, request)
	if err != nil {
		return nil, errors.NewInternalServerError(err)
	}
	rows := response.Results[0].Series[0].Rows
	if rows == nil || len(rows) == 0 {
		return nil, nil
	}
	var items []topn.Item
	total := math.DecimalPlacesWithDigitsNumber(rows[0].Values[1].GetNumberValue(), 2)
	for _, row := range rows {
		item := topn.Item{
			ID:    row.Values[0].GetStringValue(),
			Name:  row.Values[0].GetStringValue(),
			Value: math.DecimalPlacesWithDigitsNumber(row.Values[1].GetNumberValue(), 2),
			Unit:  "",
		}
		if item.Value == 0 {
			continue
		}
		item.Percent = math.DecimalPlacesWithDigitsNumber(item.Value/total*1e2, 2)
		items = append(items, item)
	}
	return items, nil
}

func (p *provider) slowReqPageTop5(sdk *cptype.SDK) ([]topn.Item, error) {
	statement := fmt.Sprintf("SELECT doc_path::tag, avg(plt::field) " +
		"FROM ta_timing " +
		"WHERE tk::tag=$terminus_key " +
		"GROUP BY doc_path::tag " +
		"ORDER BY avg(plt::field) DESC " +
		"LIMIT 5")
	params := map[string]*structpb.Value{
		"terminus_key": structpb.NewStringValue(p.InParamsPtr.TenantId),
	}

	request := &metricpb.QueryWithInfluxFormatRequest{
		Start:     strconv.FormatInt(p.InParamsPtr.StartTime, 10),
		End:       strconv.FormatInt(p.InParamsPtr.EndTime, 10),
		Statement: statement,
		Params:    params,
	}
	ctx := apis.GetContext(sdk.Ctx, func(header *transport.Header) {
		header.Set("terminus_key", p.InParamsPtr.TenantId)
	})

	response, err := p.Metric.QueryWithInfluxFormat(ctx, request)
	if err != nil {
		return nil, errors.NewInternalServerError(err)
	}
	rows := response.Results[0].Series[0].Rows
	if rows == nil || len(rows) == 0 {
		return nil, nil
	}
	var items []topn.Item
	total := math.DecimalPlacesWithDigitsNumber(rows[0].Values[1].GetNumberValue(), 2)
	for _, row := range rows {
		item := topn.Item{
			ID:    row.Values[0].GetStringValue(),
			Name:  row.Values[0].GetStringValue(),
			Value: math.DecimalPlacesWithDigitsNumber(row.Values[1].GetNumberValue(), 2),
			Unit:  "",
		}
		if item.Value == 0 {
			continue
		}
		item.Percent = math.DecimalPlacesWithDigitsNumber(item.Value/total*1e2, 2)
		v, unit := time.AutomaticConversionUnit(item.Value * 1e6)
		item.Value = v
		item.Unit = unit
		items = append(items, item)
	}
	return items, nil
}

func (p *provider) slowReqRegionTop5(sdk *cptype.SDK) ([]topn.Item, error) {
	statement := fmt.Sprintf("SELECT province::tag, avg(plt::field) " +
		"FROM ta_timing " +
		"WHERE tk::tag=$terminus_key " +
		"GROUP BY province::tag " +
		"ORDER BY avg(plt::field) DESC " +
		"LIMIT 5")
	params := map[string]*structpb.Value{
		"terminus_key": structpb.NewStringValue(p.InParamsPtr.TenantId),
	}

	request := &metricpb.QueryWithInfluxFormatRequest{
		Start:     strconv.FormatInt(p.InParamsPtr.StartTime, 10),
		End:       strconv.FormatInt(p.InParamsPtr.EndTime, 10),
		Statement: statement,
		Params:    params,
	}

	ctx := apis.GetContext(sdk.Ctx, func(header *transport.Header) {
		header.Set("terminus_key", p.InParamsPtr.TenantId)
	})

	response, err := p.Metric.QueryWithInfluxFormat(ctx, request)
	if err != nil {
		return nil, errors.NewInternalServerError(err)
	}
	rows := response.Results[0].Series[0].Rows
	if rows == nil || len(rows) == 0 {
		return nil, nil
	}
	var items []topn.Item
	total := math.DecimalPlacesWithDigitsNumber(rows[0].Values[1].GetNumberValue(), 2)
	for _, row := range rows {
		item := topn.Item{
			ID:    row.Values[0].GetStringValue(),
			Name:  row.Values[0].GetStringValue(),
			Value: math.DecimalPlacesWithDigitsNumber(row.Values[1].GetNumberValue(), 2),
			Unit:  "",
		}
		if item.Value == 0 {
			continue
		}
		item.Percent = math.DecimalPlacesWithDigitsNumber(item.Value/total*1e2, 2)
		v, unit := time.AutomaticConversionUnit(item.Value * 1e6)
		item.Value = v
		item.Unit = unit
		items = append(items, item)
	}
	return items, nil
}
