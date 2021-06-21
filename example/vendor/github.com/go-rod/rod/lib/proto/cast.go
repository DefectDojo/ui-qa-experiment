// This file is generated by "./lib/proto/generate"

package proto

/*

Cast

A domain for interacting with Cast, Presentation API, and Remote Playback API
functionalities.

*/

// CastSink ...
type CastSink struct {

	// Name ...
	Name string `json:"name"`

	// ID ...
	ID string `json:"id"`

	// Session (optional) Text describing the current session. Present only if there is an active
	// session on the sink.
	Session string `json:"session,omitempty"`
}

// CastEnable Starts observing for sinks that can be used for tab mirroring, and if set,
// sinks compatible with |presentationUrl| as well. When sinks are found, a
// |sinksUpdated| event is fired.
// Also starts observing for issue messages. When an issue is added or removed,
// an |issueUpdated| event is fired.
type CastEnable struct {

	// PresentationURL (optional) ...
	PresentationURL string `json:"presentationUrl,omitempty"`
}

// ProtoReq name
func (m CastEnable) ProtoReq() string { return "Cast.enable" }

// Call sends the request
func (m CastEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// CastDisable Stops observing for sinks and issues.
type CastDisable struct {
}

// ProtoReq name
func (m CastDisable) ProtoReq() string { return "Cast.disable" }

// Call sends the request
func (m CastDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// CastSetSinkToUse Sets a sink to be used when the web page requests the browser to choose a
// sink via Presentation API, Remote Playback API, or Cast SDK.
type CastSetSinkToUse struct {

	// SinkName ...
	SinkName string `json:"sinkName"`
}

// ProtoReq name
func (m CastSetSinkToUse) ProtoReq() string { return "Cast.setSinkToUse" }

// Call sends the request
func (m CastSetSinkToUse) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// CastStartTabMirroring Starts mirroring the tab to the sink.
type CastStartTabMirroring struct {

	// SinkName ...
	SinkName string `json:"sinkName"`
}

// ProtoReq name
func (m CastStartTabMirroring) ProtoReq() string { return "Cast.startTabMirroring" }

// Call sends the request
func (m CastStartTabMirroring) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// CastStopCasting Stops the active Cast session on the sink.
type CastStopCasting struct {

	// SinkName ...
	SinkName string `json:"sinkName"`
}

// ProtoReq name
func (m CastStopCasting) ProtoReq() string { return "Cast.stopCasting" }

// Call sends the request
func (m CastStopCasting) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// CastSinksUpdated This is fired whenever the list of available sinks changes. A sink is a
// device or a software surface that you can cast to.
type CastSinksUpdated struct {

	// Sinks ...
	Sinks []*CastSink `json:"sinks"`
}

// ProtoEvent name
func (evt CastSinksUpdated) ProtoEvent() string {
	return "Cast.sinksUpdated"
}

// CastIssueUpdated This is fired whenever the outstanding issue/error message changes.
// |issueMessage| is empty if there is no issue.
type CastIssueUpdated struct {

	// IssueMessage ...
	IssueMessage string `json:"issueMessage"`
}

// ProtoEvent name
func (evt CastIssueUpdated) ProtoEvent() string {
	return "Cast.issueUpdated"
}
