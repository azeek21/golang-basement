package ex01

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func isValidTag(str string) bool {
	re := regexp.MustCompile(`^[a-zA-z]{1,}:"[a-zA-z]{1,}"$`)
	return re.Match([]byte(str))
}

func getAllTags(str string) map[string]string {
	res := map[string]string{}
	tags := strings.Split(str, " ")
	for _, t := range tags {
		if isValidTag(t) {
			tKV := strings.Split(t, ":")
			tV, err := strconv.Unquote(tKV[1])
			tK := tKV[0]
			if err == nil && len(tV) > 0 && len(tK) > 0 {
				res[tK] = tV
			}
		}
	}
	return res
}

func joinTags(tags map[string]string, sep string) []string {
	res := []string{}
	for key, val := range tags {
		res = append(res, fmt.Sprintf("%s%s%s", key, sep, val))
	}
	return res
}

func DescribeStruct(tar interface{}) string {
	res := ""
	fields := reflect.ValueOf(tar)
	types := fields.Type()
	for i := 0; i < fields.NumField(); i++ {
		tag := types.Field(i).Tag
		if len(tag) > 0 {
			tagsMap := getAllTags(string(tag))
			tagsKVArr := joinTags(tagsMap, "=")
			res += fmt.Sprintf("%s(%s):%v\n", types.Field(i).Name, strings.Join(tagsKVArr, " "), fields.Field(i).Interface())
		} else {
			res += fmt.Sprintf("%s:%v\n", types.Field(i).Name, fields.Field(i).Interface())
		}
	}
	return res
}
