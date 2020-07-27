package channels

import (
	"github.ibm.com/coligo/satcon-client/client/actions"
	"github.ibm.com/coligo/satcon-client/client/types"
)

const (
	QueryChannels       = "channels"
	ChannelsVarTemplate = `{{define "vars"}}"orgId":"{{js .OrgID}}"{{end}}`
)

type ChannelsVariables struct {
	actions.GraphQLQuery
	OrgID string
}

func NewChannelsVariables(orgID string) ChannelsVariables {
	vars := ChannelsVariables{
		OrgID: orgID,
	}

	vars.Type = actions.QueryTypeQuery
	vars.QueryName = QueryChannels
	vars.Args = map[string]string{
		"orgId": "String!",
	}
	vars.Returns = []string{
		"uuid",
		"orgId",
		"name",
		"created",
	}

	return vars
}

type ChannelsResponse struct {
	Data *ChannelsResponseData `json:"data,omitempty"`
}

type ChannelsResponseData struct {
	Channels types.ChannelList `json:"channels,omitempty"`
}

func (c *Client) Channels(orgID, token string) (types.ChannelList, error) {
	var response ChannelsResponse

	vars := NewChannelsVariables(orgID)

	err := c.DoQuery(ChannelsVarTemplate, vars, nil, &response, token)

	if err != nil {
		return nil, err
	}

	if response.Data != nil {
		return response.Data.Channels, err
	}

	return nil, err
}