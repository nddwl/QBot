package utils

import (
	"QBot/utils/ecode"
	"encoding/json"
	"fmt"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"regexp"
	"strings"
)

func Find(s string, prefix string, max, min int) (kv []string, ok bool) {
	max += 1
	min += 1
	if len(s) > 30 {
		return
	}
	if !strings.HasPrefix(s, prefix) {
		return
	}
	kv = strings.SplitN(s, " ", max)
	if len(kv) < min {
		return
	}
	if len(kv) < max {
		kv = append(kv, make([]string, max-len(kv))...)
	}
	ok = true
	return
}

func SendErr(ctx *zero.Ctx, err error) {
	e := ecode.Cause(err)
	ctx.Send(message.Text(fmt.Sprintf("%s", e.Error())))
}

func Json(obj interface{}, b []byte) (err error) {
	err = json.Unmarshal(b, &obj)
	if err != nil {
		return ecode.JsonUnmarshalErr
	}
	return nil
}
func JsonRaw(obj interface{}, raw json.RawMessage) (err error) {
	b, err := raw.MarshalJSON()
	if err != nil {
		return ecode.JsonUnmarshalErr
	}
	err = Json(&obj, b)
	return
}

func CqImage(s string) (string, bool) {
	re := regexp.MustCompile(`,url=(.+)[,\]]`)
	if !re.MatchString(s) {
		return "", false
	}
	return re.FindStringSubmatch(s)[1], true
}
