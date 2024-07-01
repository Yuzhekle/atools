package main

import (
	"strconv"
	"strings"
)

func main() {
	id1, s := ParsePlanCode("v1737863462296514562_34_2")
	println(id1, s)

	id2, s := ParsePlanCode("1737863462296514562")
	println(id2, s)
}

func ParsePlanCode(ratePlanCode string) (rateId int64, supplierId int) {
	if strings.HasPrefix(ratePlanCode, "v") {
		ratePlanCode = strings.TrimPrefix(ratePlanCode, "v")
		strs := strings.Split(ratePlanCode, "_")
		rateId = 0
		supplierId, _ = strconv.Atoi(strs[1])
	} else {
		// 转换ratePlanCode
		rateId1, _ := strconv.Atoi(ratePlanCode)
		rateId = int64(rateId1)
	}
	return
}
