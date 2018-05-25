# zendeskticketattachments
Simple module to download all of the attachments from a given Zendesk ticket

#### Example Usage

```go
package main

import (
	"github.com/jakeloredo/zendeskticketattachments"
)

func main() {
	// Setup Zendesk struct
	zd := zendeskticketattachments.Zendesk{
		Username:  "your_username@your_domain.com",
		Password:  "your_password_or_token",
		Subdomain: "your_subdomain"} // This would mean the full url is https://your_subdomain.zendesk.com/

	// Download all attachments from ticket number 123456
	zd.DownloadAttachments("123456")
}

```
