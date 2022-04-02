// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: channel.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyChannelEnabledStatusRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyChannelEnabledStatusResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateNotifyChannelEnabledRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateNotifyChannelEnabledResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyChannelEnabledRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyChannelEnabledResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteNotifyChannelRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteNotifyChannelResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyChannelRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyChannelResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateNotifyChannelRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateNotifyChannelResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyChannelsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyChannelsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateNotifyChannelRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateNotifyChannelResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyChannelTypesRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyChannelTypesResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NotifyChannelTypeResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NotifyChannelType)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NotifyChannelProviderType)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NotifyChannel)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyChannelsEnabledRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyChannelsEnabledResponse)(nil)

// GetNotifyChannelEnabledStatusRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyChannelEnabledStatusRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}

// GetNotifyChannelEnabledStatusResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyChannelEnabledStatusResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "hasEnable":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.HasEnable = val
			case "enableChannelName":
				m.EnableChannelName = vals[0]
			}
		}
	}
	return nil
}

// UpdateNotifyChannelEnabledRequest implement urlenc.URLValuesUnmarshaler.
func (m *UpdateNotifyChannelEnabledRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "enable":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Enable = val
			}
		}
	}
	return nil
}

// UpdateNotifyChannelEnabledResponse implement urlenc.URLValuesUnmarshaler.
func (m *UpdateNotifyChannelEnabledResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "enable":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Enable = val
			}
		}
	}
	return nil
}

// GetNotifyChannelEnabledRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyChannelEnabledRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scopeId":
				m.ScopeId = vals[0]
			case "scopeType":
				m.ScopeType = vals[0]
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}

// GetNotifyChannelEnabledResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyChannelEnabledResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.Id = vals[0]
			case "data.name":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.Name = vals[0]
			case "data.type":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.Type == nil {
					m.Data.Type = &NotifyChannelType{}
				}
			case "data.type.name":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.Type == nil {
					m.Data.Type = &NotifyChannelType{}
				}
				m.Data.Type.Name = vals[0]
			case "data.type.displayName":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.Type == nil {
					m.Data.Type = &NotifyChannelType{}
				}
				m.Data.Type.DisplayName = vals[0]
			case "data.scopeId":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.ScopeId = vals[0]
			case "data.scopeType":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.ScopeType = vals[0]
			case "data.creatorName":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.CreatorName = vals[0]
			case "data.createAt":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.CreateAt = vals[0]
			case "data.updateAt":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.UpdateAt = vals[0]
			case "data.channelProviderType":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.ChannelProviderType == nil {
					m.Data.ChannelProviderType = &NotifyChannelProviderType{}
				}
			case "data.channelProviderType.name":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.ChannelProviderType == nil {
					m.Data.ChannelProviderType = &NotifyChannelProviderType{}
				}
				m.Data.ChannelProviderType.Name = vals[0]
			case "data.channelProviderType.displayName":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.ChannelProviderType == nil {
					m.Data.ChannelProviderType = &NotifyChannelProviderType{}
				}
				m.Data.ChannelProviderType.DisplayName = vals[0]
			case "data.enable":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.Enable = val
			}
		}
	}
	return nil
}

// DeleteNotifyChannelRequest implement urlenc.URLValuesUnmarshaler.
func (m *DeleteNotifyChannelRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			}
		}
	}
	return nil
}

// DeleteNotifyChannelResponse implement urlenc.URLValuesUnmarshaler.
func (m *DeleteNotifyChannelResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			}
		}
	}
	return nil
}

// GetNotifyChannelRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyChannelRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			}
		}
	}
	return nil
}

// GetNotifyChannelResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyChannelResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.Id = vals[0]
			case "data.name":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.Name = vals[0]
			case "data.type":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.Type == nil {
					m.Data.Type = &NotifyChannelType{}
				}
			case "data.type.name":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.Type == nil {
					m.Data.Type = &NotifyChannelType{}
				}
				m.Data.Type.Name = vals[0]
			case "data.type.displayName":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.Type == nil {
					m.Data.Type = &NotifyChannelType{}
				}
				m.Data.Type.DisplayName = vals[0]
			case "data.scopeId":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.ScopeId = vals[0]
			case "data.scopeType":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.ScopeType = vals[0]
			case "data.creatorName":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.CreatorName = vals[0]
			case "data.createAt":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.CreateAt = vals[0]
			case "data.updateAt":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.UpdateAt = vals[0]
			case "data.channelProviderType":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.ChannelProviderType == nil {
					m.Data.ChannelProviderType = &NotifyChannelProviderType{}
				}
			case "data.channelProviderType.name":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.ChannelProviderType == nil {
					m.Data.ChannelProviderType = &NotifyChannelProviderType{}
				}
				m.Data.ChannelProviderType.Name = vals[0]
			case "data.channelProviderType.displayName":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.ChannelProviderType == nil {
					m.Data.ChannelProviderType = &NotifyChannelProviderType{}
				}
				m.Data.ChannelProviderType.DisplayName = vals[0]
			case "data.enable":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.Enable = val
			}
		}
	}
	return nil
}

// UpdateNotifyChannelRequest implement urlenc.URLValuesUnmarshaler.
func (m *UpdateNotifyChannelRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "name":
				m.Name = vals[0]
			case "type":
				m.Type = vals[0]
			case "channelProviderType":
				m.ChannelProviderType = vals[0]
			case "enable":
				m.Enable = vals[0]
			case "scopeId":
				m.ScopeId = vals[0]
			case "scopeType":
				m.ScopeType = vals[0]
			}
		}
	}
	return nil
}

// UpdateNotifyChannelResponse implement urlenc.URLValuesUnmarshaler.
func (m *UpdateNotifyChannelResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.Id = vals[0]
			case "data.name":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.Name = vals[0]
			case "data.type":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.Type == nil {
					m.Data.Type = &NotifyChannelType{}
				}
			case "data.type.name":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.Type == nil {
					m.Data.Type = &NotifyChannelType{}
				}
				m.Data.Type.Name = vals[0]
			case "data.type.displayName":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.Type == nil {
					m.Data.Type = &NotifyChannelType{}
				}
				m.Data.Type.DisplayName = vals[0]
			case "data.scopeId":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.ScopeId = vals[0]
			case "data.scopeType":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.ScopeType = vals[0]
			case "data.creatorName":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.CreatorName = vals[0]
			case "data.createAt":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.CreateAt = vals[0]
			case "data.updateAt":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.UpdateAt = vals[0]
			case "data.channelProviderType":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.ChannelProviderType == nil {
					m.Data.ChannelProviderType = &NotifyChannelProviderType{}
				}
			case "data.channelProviderType.name":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.ChannelProviderType == nil {
					m.Data.ChannelProviderType = &NotifyChannelProviderType{}
				}
				m.Data.ChannelProviderType.Name = vals[0]
			case "data.channelProviderType.displayName":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.ChannelProviderType == nil {
					m.Data.ChannelProviderType = &NotifyChannelProviderType{}
				}
				m.Data.ChannelProviderType.DisplayName = vals[0]
			case "data.enable":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.Enable = val
			}
		}
	}
	return nil
}

// GetNotifyChannelsRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyChannelsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "pageNo":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageNo = val
			case "pageSize":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageSize = val
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}

// GetNotifyChannelsResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyChannelsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "pageNo":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageNo = val
			case "pageSize":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageSize = val
			case "total":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Total = val
			}
		}
	}
	return nil
}

// CreateNotifyChannelRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateNotifyChannelRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "type":
				m.Type = vals[0]
			case "channelProviderType":
				m.ChannelProviderType = vals[0]
			}
		}
	}
	return nil
}

// CreateNotifyChannelResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateNotifyChannelResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.Id = vals[0]
			case "data.name":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.Name = vals[0]
			case "data.type":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.Type == nil {
					m.Data.Type = &NotifyChannelType{}
				}
			case "data.type.name":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.Type == nil {
					m.Data.Type = &NotifyChannelType{}
				}
				m.Data.Type.Name = vals[0]
			case "data.type.displayName":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.Type == nil {
					m.Data.Type = &NotifyChannelType{}
				}
				m.Data.Type.DisplayName = vals[0]
			case "data.scopeId":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.ScopeId = vals[0]
			case "data.scopeType":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.ScopeType = vals[0]
			case "data.creatorName":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.CreatorName = vals[0]
			case "data.createAt":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.CreateAt = vals[0]
			case "data.updateAt":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				m.Data.UpdateAt = vals[0]
			case "data.channelProviderType":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.ChannelProviderType == nil {
					m.Data.ChannelProviderType = &NotifyChannelProviderType{}
				}
			case "data.channelProviderType.name":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.ChannelProviderType == nil {
					m.Data.ChannelProviderType = &NotifyChannelProviderType{}
				}
				m.Data.ChannelProviderType.Name = vals[0]
			case "data.channelProviderType.displayName":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				if m.Data.ChannelProviderType == nil {
					m.Data.ChannelProviderType = &NotifyChannelProviderType{}
				}
				m.Data.ChannelProviderType.DisplayName = vals[0]
			case "data.enable":
				if m.Data == nil {
					m.Data = &NotifyChannel{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.Enable = val
			}
		}
	}
	return nil
}

// GetNotifyChannelTypesRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyChannelTypesRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetNotifyChannelTypesResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyChannelTypesResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// NotifyChannelTypeResponse implement urlenc.URLValuesUnmarshaler.
func (m *NotifyChannelTypeResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "displayName":
				m.DisplayName = vals[0]
			}
		}
	}
	return nil
}

// NotifyChannelType implement urlenc.URLValuesUnmarshaler.
func (m *NotifyChannelType) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "displayName":
				m.DisplayName = vals[0]
			}
		}
	}
	return nil
}

// NotifyChannelProviderType implement urlenc.URLValuesUnmarshaler.
func (m *NotifyChannelProviderType) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "displayName":
				m.DisplayName = vals[0]
			}
		}
	}
	return nil
}

// NotifyChannel implement urlenc.URLValuesUnmarshaler.
func (m *NotifyChannel) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "name":
				m.Name = vals[0]
			case "type":
				if m.Type == nil {
					m.Type = &NotifyChannelType{}
				}
			case "type.name":
				if m.Type == nil {
					m.Type = &NotifyChannelType{}
				}
				m.Type.Name = vals[0]
			case "type.displayName":
				if m.Type == nil {
					m.Type = &NotifyChannelType{}
				}
				m.Type.DisplayName = vals[0]
			case "scopeId":
				m.ScopeId = vals[0]
			case "scopeType":
				m.ScopeType = vals[0]
			case "creatorName":
				m.CreatorName = vals[0]
			case "createAt":
				m.CreateAt = vals[0]
			case "updateAt":
				m.UpdateAt = vals[0]
			case "channelProviderType":
				if m.ChannelProviderType == nil {
					m.ChannelProviderType = &NotifyChannelProviderType{}
				}
			case "channelProviderType.name":
				if m.ChannelProviderType == nil {
					m.ChannelProviderType = &NotifyChannelProviderType{}
				}
				m.ChannelProviderType.Name = vals[0]
			case "channelProviderType.displayName":
				if m.ChannelProviderType == nil {
					m.ChannelProviderType = &NotifyChannelProviderType{}
				}
				m.ChannelProviderType.DisplayName = vals[0]
			case "enable":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Enable = val
			}
		}
	}
	return nil
}

// GetNotifyChannelsEnabledRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyChannelsEnabledRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetNotifyChannelsEnabledResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyChannelsEnabledResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}