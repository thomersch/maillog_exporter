// Code generated by "stringer -type=event"; DO NOT EDIT

package collector

import "fmt"

const _event_name = "eventEmptyconnectSMTPdisconnectSMTPconnectIMAPdisconnectIMAPincomingMailoutgoingMailunknownUsergreylistedgreylistPassspamBlocked"

var _event_index = [...]uint8{0, 10, 21, 35, 46, 60, 72, 84, 95, 105, 117, 128}

func (i event) String() string {
	if i < 0 || i >= event(len(_event_index)-1) {
		return fmt.Sprintf("event(%d)", i)
	}
	return _event_name[_event_index[i]:_event_index[i+1]]
}
