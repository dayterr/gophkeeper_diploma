package tuiagent

import (
	"github.com/rivo/tview"
	"net/http"
)

type TUIClient struct {
	HTTPSender *http.Client
	TUIApp *tview.Application
	Address string
}
