package helpers

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cast"
)

const (
	// See http://golang.org/pkg/time/#Parse
	timeFormat              = "2006-01-02 15:04:05.999"
	dateFormat              = "2006-01-02"
	dateFormatyyyMMdd       = "20060102"
	dateFormatyyyMMddHHmmSS = "20060102150405"
	dateFormatHH            = "15:04"
)

// RandomString get random string from given string length
// commonly used to generate requestId in json response
func RandomNumber(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return fmt.Sprintf("%s", string(result))
}

// RandomString get random string from given string length
// commonly used to generate requestId in json response
func RandomString(strlen int, id string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return fmt.Sprintf("%s%s", id, string(result))
}

/**
 * InArray like PHP in_array function
 * @param  {[type]} val   interface{}  [description]
 * @param  {[type]} array interface{}) (exists       bool, index int [description]
 * @return {[type]}       [description]
 * source: http://codereview.stackexchange.com/questions/60074/in-array-in-go
 */
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	if array == nil {
		return false, -1
	}
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func MakeTimestamp(now time.Time) int64 {
	return now.UnixNano() / int64(time.Millisecond)
}

func Masking(name string) string {
	r, _ := regexp.Compile(`[^ ]+`)
	res := r.FindAllString(name, -1)

	var fullWord = []string{}

	for _, n := range res {
		x := 0
		var emptySlice = []string{}

		if len(n) <= 3 {
			var z = ""
			for x < len(n) {
				yo := string([]rune(n)[x])
				z = yo
				emptySlice = append(emptySlice, z)
				x++
			}
		} else {
			var beginMask = len(n) - 3
			for x < len(n) {
				var z = ""
				yo := string([]rune(n)[x])
				if x >= beginMask {
					z = yo
				} else {
					z = "*"
				}
				emptySlice = append(emptySlice, z)
				x++
			}
		}
		justString := strings.Join(emptySlice, "")
		fullWord = append(fullWord, justString)
	}

	fixMasking := strings.Join(fullWord, " ")
	return fixMasking
}

func Substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
func TimeToSecond(t time.Duration) int64 {
	return t.Nanoseconds() / 1000000
}
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

//conver
func ConvertTo62(value string) string {
	if value != "" && strings.HasPrefix(value, "0") {
		temp := value[1:]
		return "62" + temp
	}
	return value

}
func ConvertTo0(value string) string {
	if value != "" && strings.HasPrefix(value, "62") {
		temp := value[2:]
		return "0" + temp
	}
	return value
}

func TimeStamp() string {
	t := time.Now()
	res := t.Format(dateFormatyyyMMddHHmmSS)
	return res
}

func Time() string {
	t := time.Now()
	res := t.Format(timeFormat)
	return res
}
func TimeStampSimple() string {
	t := time.Now()
	res := t.Format(dateFormatyyyMMdd)
	return res
}

func InTimeSpan(tstart, tend, tcheck string) bool {
	loc, _ := time.LoadLocation("Asia/Jakarta")

	check, _ := time.ParseInLocation(dateFormatHH, tcheck, loc)
	start, _ := time.ParseInLocation(dateFormatHH, tstart, loc)
	end, _ := time.ParseInLocation(dateFormatHH, tend, loc)

	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}

func TimeHour() string {
	t := time.Now()
	res := t.Format(dateFormatHH)
	return res
}

func VisitFile(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() && info.Name() != "." {
			*files = append(*files, path)
		}
		return nil
	}
}

func ConvertParam(params ...interface{}) string {
	result := ""
	for _, param := range params {
		result += cast.ToString(param) + ","
	}
	return result
}
