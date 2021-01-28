package Kcpnet

/*
	external struct collection for server or client for special using.
*/

type ExternalCollection struct {
	centerSession  *CenterSessionMgr
	client         *KcpClient
	externalClient *SessionMgr
	// ...
}

func NewExternalCollection() *ExternalCollection {
	return &ExternalCollection{
		centerSession: &CenterSessionMgr{},
	}
}

func (this *ExternalCollection) GetCenterSession() *CenterSessionMgr {
	return this.centerSession
}

func (this *ExternalCollection) SetCenterClient(c *KcpClient) {
	this.client = c
}

func (this *ExternalCollection) GetCenterClient() *KcpClient {
	return this.client
}

func (this *ExternalCollection) SetExternalClient(sess *SessionMgr) {
	this.externalClient = sess
}

func (this *ExternalCollection) GetExternalClient() *SessionMgr {
	return this.externalClient
}
