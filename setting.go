package poego

type GetChannelResp struct {
	TchannelData struct {
		MinSeq          string `json:"minSeq"`
		Channel         string `json:"channel"`
		ChannelHash     string `json:"channelHash"`
		BoxName         string `json:"boxName"`
		BaseHost        string `json:"baseHost"`
		TargetURL       string `json:"targetUrl"`
		EnableWebsocket bool   `json:"enableWebsocket"`
	} `json:"tchannelData"`
}
