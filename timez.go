// gotz/timez.go
package main

import (
    "time"
)

func TzConvert(tzname string) string {
    local := time.Now().UTC()

    location, err := time.LoadLocation(tzname)

    if err == nil {
        local := local.In(location)
        return local.Format(time.UnixDate) + "\n"
    }

    return local.Format(time.UnixDate) + "\n"
}
