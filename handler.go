// gotz/handler.go
package main

import (
    "bufio"
    "net"
    "time"
)

// handles the connection timeout and reading buffer
func TzHandler(cx net.Conn) {
    cx.SetReadDeadline(time.Now().Add(time.Second * 20)) // 20 second timeout
    defer cx.Close() 
    for {
        request, prefix, err := bufio.NewReader(cx).ReadLine()
        if err != nil && !prefix {
            logError.Println(err)
            break
        }

        reTime := TzConvert(string(request))
        cx.Write([]byte(reTime))

    }
}


