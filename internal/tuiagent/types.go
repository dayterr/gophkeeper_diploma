package tuiagent

import (
	"net/http"

	"github.com/rivo/tview"
)

type TUIClient struct {
	HTTPSender *http.Client
	TUIApp     *tview.Application
	Address    string
}
