// package main

// import (
// 	"encoding/base64"
// 	"fmt"
// )

// func main() {
// 	// body := "ymSDL_v2d0p2ogtYYNgMZbOKXYHzFZSfNqaccjnH6B-5B78o8iVas8UOUh2myns-MQ_LLmNQtxiraZtRYlJaPqYha7xhdPgF9q7tlbMC6IyJW4JcaMWrkisJ3ego2K-0MlDfW-Qk5SqR7i_q9YFs8FgHLkstXD9dOE4xBviwum9eCzhWXrPZ8JkgVei6vz5vQTnhTmE9FLxihjWVtY243BwXjRFLcReNhARBfRidRGjR2CAvZp1N4Ebo_A_xKLn2c0QRwnu3eXK4WMzD6SztIeRlfeaFSZo6c88"
// 	// body := "eJxNkbtSwzAQRf9FNYr1sB05XYaKKkwGKkwhS+vH4NhCDwaTyb8j5GRIqXP37l7tnpGx8/fyshhAO0TQA7LwGcD5Z2nlyUV2rlE/exifdI12NWI0r0hFa/RQoymcGrCH9jjPpyjSiGalgpGTWuI7OlU/jNrCtO/ARfL2fufa6zD6CNkde7zWR0wukVvp4TqYClFVeVFwwUmafg16lRtQfMtVg1vCAedCKCwUKBwdjRZtQRq+XW0x7K0lFWVeVQk78H6YOrfm1oMzo1weg41h0mdqlMoGdwQ5pre3ARIb5dQF2UGiPz1WMf5feDd0U2JABKtoUZYaOCtL2bKmLVqut43OGdV6DeDlcoyd4LY5UB9Pq58RxjFlmK9rT9Ih+JuWY0IxYXHo5f9+r3aMx+u9N7sss6CNcV9q4+dgT8OkN2rKpBmyL5rNVoPNjIXUFl1+AeHtqDk="
// 	body := "eJyrVkotKsovUrKqVkrOT0lVsjIyMDTQUcpNLS5OTAdylYpSC0tTi0sU8vJLFBLLEjNzEpNyUpVqawEgFROD"
// 	newBody, err := base64.StdEncoding.DecodeString(body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(string(newBody))
// }

package main

import (
    "encoding/base64"
    "fmt"
)

func main() {
    // base64 编码的字符串
    encodedString := "eJyrVkotKsovUrKqVkrOT0lVsjIyMDTQUcpNLS5OTAdylYpSC0tTi0sU8vJLFBLLEjNzEpNyUpVqawEgFROD"

    // 解码字符串
    decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
    if err != nil {
        fmt.Println("解码错误:", err)
        return
    }

    // 打印解码后的字符串
    fmt.Println("解码后的字符串:", string(decodedBytes))
}