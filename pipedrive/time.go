package pipedrive

// http://fuckinggodateformat.com/
import "time"

type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.String()
}

func (t Timestamp) Format() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.Format("2006-01-02")
}

func (t Timestamp) FormatFull() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.Format("2006-01-02 15:04:05")
}
