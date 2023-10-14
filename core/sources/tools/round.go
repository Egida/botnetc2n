package tools

import (
	"fmt"
	"strings"
	"time"
)

//this is used for when the time.Since function is wanted
//this will auto resolve the timestamp into its current format
//with the format of f you can display if you want the full time display
func ResolveTimeStamp(t time.Time, f bool) string {
	//detects if the session has been open for longer than one hour
	if strings.Count(time.Since(t).Round(0).String(), "d") > 0 {
		if f {
			//displays the full time measurement
			return fmt.Sprintf("%.0f Days", time.Since(t).Hours()/24)
		}
		return fmt.Sprintf("%.0fdays", time.Since(t).Hours()/24)
	} else if strings.Count(time.Since(t).Round(0).String(), "h") > 0 {
		if f {
			//displays the full time measurement
			return fmt.Sprintf("%.0f Hours", time.Since(t).Hours())
		}
		return fmt.Sprintf("%.0fhrs", time.Since(t).Hours())
	} else if strings.Count(time.Since(t).Round(0).String(), "m") > 0 && !strings.Contains(time.Since(t).Round(0).String(), "ms") {
		if f {
			//displays the full time measurement
			return fmt.Sprintf("%.0f Minutes", time.Since(t).Minutes())
		}
		return fmt.Sprintf("%.0fmins", time.Since(t).Minutes())
	} else if strings.Count(time.Since(t).Round(0).String(), "s") > 0 {
		if f {
			//displays the full time measurement
			return fmt.Sprintf("%.0f Seconds", time.Since(t).Seconds())
		}
		return fmt.Sprintf("%.0fsecs", time.Since(t).Seconds())
	}

	return "0secs"
}

//this is used for when the time.until function is wanted
//this will auto resolve the timestamp into its current format
//with the format of f you can display if you want the full time display
func ResolveTimeStampUnix(t time.Time, f bool) string {
	//detects if the session has been open for longer than one hour
	if strings.Count(time.Until(t).Round(0).String(), "h") > 0 && int(time.Until(t).Round(0).Hours()) > 24 {
		if f {
			//displays the full time measurement
			return fmt.Sprintf("%.0f Days", time.Until(t).Hours()/24)
		}
		return fmt.Sprintf("%.0fdays", time.Until(t).Hours()/24)
	} else if strings.Count(time.Until(t).Round(0).String(), "h") > 0 {
		if f {
			//displays the full time measurement
			return fmt.Sprintf("%.0f Hours", time.Until(t).Hours())
		}
		return fmt.Sprintf("%.0fhrs", time.Until(t).Hours())
	} else if strings.Count(time.Until(t).Round(0).String(), "m") > 0 && !strings.Contains(time.Until(t).Round(0).String(), "ms") {
		if f {
			//displays the full time measurement
			return fmt.Sprintf("%.0f Minutes", time.Until(t).Minutes())
		}
		return fmt.Sprintf("%.0fmins", time.Until(t).Minutes())
	} else if strings.Count(time.Until(t).Round(0).String(), "s") > 0 {
		if f {
			//displays the full time measurement
			return fmt.Sprintf("%.0f Seconds", time.Until(t).Seconds())
		}
		return fmt.Sprintf("%.0fsecs", time.Until(t).Seconds())
	}

	return "0secs"
}

func ResolveString(str string) time.Duration {
	switch strings.ToLower(str) {
	case "s", "seconds", "second", "secs":
		return time.Second
	case "m", "minutes", "mins", "minute":
		return time.Minute
	case "ms", "milliseconds", "millis", "millisecond":
		return time.Millisecond
	default:
		return time.Second
	}
}