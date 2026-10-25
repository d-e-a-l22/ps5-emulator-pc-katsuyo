package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.7" )

func SzeMiePNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ik5i3WiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzFzYErSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJR45cnlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uMZYBTOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AlJqfrmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Be1P9qP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K4kVIvAGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzJFvkEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Psz1JsChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJJZiq1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIAbF7SnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGY4waG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3STLjRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iWUN55BSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKiSwVKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L44ou5brWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbWS1TqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yuEEa9hBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZW1CnNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8j6xe418Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bzIFabDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0SlvUZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JD6i0zZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fuACutmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCPUmOBBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YKjD7g3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bf0DVbYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1SuLlrzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func luVlOGLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2zCsSEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjdOs0U2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tA3odLknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRBqeNovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vL1d0jMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6OT97CcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func el66i7xgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdMVUJO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yoegw3gPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CxcWbmYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OdN3wn3WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EdnjCJbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfZPYDNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGkCnOC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06owumqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YrmfPTenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A34Ap3NsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7VvieWb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SzUdIAIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxHloz06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVY3OC4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6W5jWDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LarmlrsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDAnaMH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Reec8XzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdSY1FmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4W5RkpdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpv9JQmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8D6XHVuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2jqVIaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qftBdM9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uItTVOiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lyc2ts6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wF7qmP6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CnCN56y0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKcPfTmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2e8JGcn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNTu6KXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHjlXQn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mCc77zdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lc5t395qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKuRTHPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CoBbj2rIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiPmzWrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IpV00LsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dT2cPETQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svmgbKSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zb7cmLj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUUXieGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMKVL1hKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqPIKdgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFg4aepLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8UewRpqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZTOmJhjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UcOZllzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func os90GGC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2T1QJUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vN4P6SE7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdJeb5emWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCmibvNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hcNtjqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JunqmIpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rYhWEhcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miVWCDkwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJNW7iqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXktZGztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6WCJ8Q6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5sTcORpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPFHkQxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1p83Rs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRJF46VvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func apXjSByhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nCb6LDiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0iTHmDc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d55bJxH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7xMX3ZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vGh1r46sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVHXR3tDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sK5pGAIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jd0ao6RDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YpbqSZDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UMoLhETiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rzl2qjS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izcUx35oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5LtYOQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bIb7hJZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSJ8c7aeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thHqgPEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MsCywmiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hBdTFCatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQJJHFwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NPIEagFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0TPjEcsiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90K75kvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5bHWDS80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CAYWaGUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3udQqyf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cyTJyajKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBFd19OOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8vZ4E2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XtmXP9d1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQzm5ClFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BLhfHH6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9NKLjFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2f4etquIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pz17HM9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ckLOyqJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rEnPqnhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGDh5SwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N8CBFfs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTGBkk1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPC8fiKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Om4rWaGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9QrTKZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jAA70x1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0mWZvUtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PYa7zfszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XflIzDGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNJtZYWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DYmtBFDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2x2pbgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zILXbuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VneHW3ZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mrMOd1CqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EN2bOlTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ej5m26nwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDz5gTiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i0Bv6DGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aN6Sb6vxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e39H6yoCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWb31kPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ZYA0QvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T1HPqvUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xq8ZbwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z8zo31mgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqKNpaEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTwBuEaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZUFywr2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPSF2lO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZBoys9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1VBCa4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oc5ubWBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEVQgxU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5TbKT4iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAeP75I2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xvefZPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jtoXDdxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a91Anry9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzatqOutWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func exRhaWmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gIh1LHspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dr3QsrW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IhQwTzGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHlW3wJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KaGkfuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFV3QqHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9Cc5KrhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMMLTxFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEycTmmcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iHRAVVP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95eZPzUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIGDPGpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTLH65GVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swEDDuIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kjqe90nqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vo36dBOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mNZVdmj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2nLjezvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wr25sTtnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DF5gFrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9koWUZ6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aurrfyR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NvUn07TbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXKgOH9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7JuimN11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hr7iJ5N8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5RFH96Q7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZ55zB0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w8j1ON8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jIQ4x5oZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ti5Rvh6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3I6O9PLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 194PYVgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WI6EE9lBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVBvCwzLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gsfmWADFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fL65UYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzm69dnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwCHsIYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ansNT99IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uY6zjerbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxhRKUxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNv8jOGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLKbjiCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FyMZLgBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbtDGRMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 94gc6PmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1erCDMLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhPuYOOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXPpl8Y1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzDFjN3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IqoRs074Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtIuQa7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrVjpgkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BOm6UMPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRJugC12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBDsXDmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yGLfS6jbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQYlHv2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNTwetJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2R1J0LWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZVmr7MsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ij8HZcg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQZatzceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3F2bUteGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JT9Hq69VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFpxzZfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5uo4mCUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WXlpGEokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TykFyQ6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kaHjZQbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHpZK6ZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AM8dqeHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxAyrQYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZNyxWswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tc93LgZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqY6qjIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5PQ4ycRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REc0Ui0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MIj9EoGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ljhgzs9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nU2tKlt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1qYbOvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFrqhraaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s12VIOLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wgvM0R3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5NCMCqLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZgllrKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pyMYp7MvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Np8XzItcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OCv8NSNZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZTGxew7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mamFKUQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7vefPdXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWetMdbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAjWN58RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjlBml6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzWuRO8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tdY5KLj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uxAqDGKJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFq5BNp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQVEglI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmQaFiWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9EJHV6GsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fq6HeqAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2mXx5GTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psXjXBs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9wGiPp5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c7BBHNY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrrouXROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAw2hpuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func myl6abPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDZ4d3eFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XMBehhRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zExYcocsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TDEVYEuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTFENoSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gfi6U0jmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoNx0bHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8gnvWjGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GtXHYpB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tbm81EKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i0oryHgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41qMeyxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdrrUBIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sI0UayH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCaokiabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4L9IyQUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFZ9IWXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nw9fihrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0p9cm9Z0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vQ7fTYEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CGmEw2M2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5h40g3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYrLIIonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsV8HFYwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JWukQCOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lB5bG8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func amyeTezGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOXPuoTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHyyWXwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RT2Y8cSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1vdiLj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwNIvdYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f3BAvHQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R44hpn4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func prvYHbVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uvfw7oPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLNqUWKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVb82IreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nRbYWe51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQWQXYziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kc9S4IheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tRxrXBg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CvEUoyr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UMvSi6QmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enR8rjxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wzs9KamiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eXJ3hp5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7vX6ndwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LcLoM1V7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4161udJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwVl4TR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pjAD0cIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pprCUXG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 139XpMU4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdcNEjpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLHHVpZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOx748YDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRDNZ4bxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2HkHSPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbiJW2mBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwazupXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n35HHtDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6lq4MUYEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEoUOjlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bSRr35AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tlMKTtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kI8ZP2arWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FqJGh8X2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIq9GCeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ovaQYTuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QBimBzOJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tyj2UTXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ZmTqeQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vOHaAvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UU772iX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8plL9SuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yv7s0P69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eAHXKrmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQRnhHTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Xb5LW7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HxtZzmXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5exRqhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBjEVqltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0NKMQMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AMQYqXsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ny0zq4E0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rmCysPJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCzj2gEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WsxKvBFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DMeS1BBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYKFnSiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mx40WJKJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xE2MdrcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ux3o5FkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5mbxneNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YH7u0okxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oGWwKWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGw4lmiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vme3mCsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AjDClMtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jkWYPYhsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CULd9M8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9ECYK54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U6lxLeRBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nry7K0zQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TJPHP44TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Piz2B6RrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wO6ixiHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVJeCHeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k94GnwZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LBtjU34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xHM0evuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Def6jURsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OcFBnRH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b6EuvBRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjkyOuyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUPSmr0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3FvK13qIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pEDCpV4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ci8PFVLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhLf768WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yKltrpcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNyqJmG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0nsBeHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThWcjZc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F078A341Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4y4Vi02CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8UZ3RhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06LrUqfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L416ToO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2YClxQgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z10gKLIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUGgIhtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1okxKyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DAdhaMx7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5q7wpoIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ekL8VX1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FoJGEdMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mcgpP1jHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9qRNPBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZkNXcN5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dahiNzkdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWYg0ZiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzpSsOQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CyEoGUPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4kMBO8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C25Zudf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Z6CsFRsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tfpQVX11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5xRBgcFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nn6NvfZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLM64EeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBbO1aYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hTlunqvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PurKXrBBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NIKqIdxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIDwtXHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mvh6AIK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0joETBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RNArmLAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SLWBTAhpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2BolA5RhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXdSYM4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdvygvLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bbXqi0WqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lRzmdTUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1N5WnQg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oBx2e3d3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ercYtbdLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIpwhKf7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miqtKpOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n2a6Q8UeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsZfaUBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y5krsWiiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2iCzVLfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func llAndl9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJJM3I7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRgXjBKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AkdVCdHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDSfmbWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H40qJ5lSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hN1lPVYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6qAqM1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SyVYrKs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSW549laWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgtI6zzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eRvKzW6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PyMksMsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RA6kNZfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eARRGWVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGQYlFizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mhkb40dqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMsjrKbXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iM2MbV93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zwW9bdPJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ya1GXxH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Slt44ehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4GeK4kG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8bhmL0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Adm6vU7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3upJIIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vCqvMDDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vEi41Q40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oO5ebG8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VobZpOP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yZKYwrnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I45oup1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZQj0SNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LH3gwCjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suEjtCy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9h1lPAm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4bAm5EwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func euxJMj3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RAoQ4kQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxtHAV1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wo2H5ztHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNWtBctOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U4x9YW2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OFBYvzW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FX4vpLibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvXPVEg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBf7M00CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ekYPQjaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsTThwoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdLIv7JVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVrmxdhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xzpfpJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I24xaBgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

