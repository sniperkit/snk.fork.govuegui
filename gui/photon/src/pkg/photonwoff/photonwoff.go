/*
Sniperkit-Bot
- Status: analyzed
*/

package photonwoff

import (
	"net/http"
)

// This code is generated by createhandlerlib
// github.com/as27/createhandlerlib

// Version of the library
const Version = "0.1.2"

// Handler can be used from http Server
func Handler(w http.ResponseWriter, r *http.Request) {

	w.Write(libBytes)
}
