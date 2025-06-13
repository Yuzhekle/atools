package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

func init() {
	run()
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("UnixNano", GenerateRandomNumberFromUUID(7))
	}
	println()
	println()
	fmt.Println("Unix", time.Now().Unix()/1000)
	fmt.Println("UnixNano", time.Now().UnixNano())
	fmt.Println("UnixNano", time.Now().UnixNano())
	fmt.Println("UnixNano", time.Now().UnixNano())
	fmt.Println("Micro", time.Now().UnixMicro())
	fmt.Println("mill", time.Now().UnixMilli())
	fmt.Println("Mill", time.Now().UnixNano()/1e6)
	fmt.Println("Second", time.Now().UnixNano()/1e9)

	fmt.Println("UnixNano", time.Now().UnixNano())
	for {
	}
}

// 生成一个测试用例，生成一个随机数，长度为 n
func GenerateRandomNumber(n int) int {
	// 获取当前 Unix 时间戳（毫秒）
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)

	// 将时间戳转换为字符串并取后 n 位
	timeStr := strconv.FormatInt(currentTime, 10)
	if len(timeStr) > n {
		timeStr = timeStr[len(timeStr)-n:] // 取最后 n 位数字
	}
	fmt.Println("timeStr", timeStr)

	// 使用随机数增加不可预测性
	// 随机种子
	seed := time.Now().UnixNano()
	// 随机源
	source := rand.NewSource(uint64(seed))
	// 创建 rand 实例
	random := rand.New(source)
	randomSuffix := random.Intn(100) // 获取 0 到 99 的随机数

	// 计算 10 的 n 次方
	powN := int(math.Pow(10, float64(n)))
	// 将时间戳的最后 n 位数字与随机数拼接，并确保返回的结果为 n 位数字
	result := (currentTime%int64(powN) + int64(randomSuffix)) % int64(powN)

	return int(result)
}


// GenerateRandomNumber 生成 n 位随机数，增加唯一性
func GenerateRandomNumber2(n int) int {
	// 获取当前 Unix 时间戳（纳秒）
	currentTime := time.Now().UnixNano()

	// 使用随机数增强唯一性
	seed := time.Now().UnixNano()
	source := rand.NewSource(uint64(seed))
	random := rand.New(source)

	// 获取一个更大的随机数范围
	randomSuffix := random.Intn(10000) // 增加随机数范围，0-999

	// 计算 10 的 n 次方
	powN := int64(1)
	for i := 0; i < n; i++ {
		powN *= 10
	}

	// 确保生成的是 n 位数字
	result := (currentTime%powN + int64(randomSuffix)) % powN

	return int(result)
}

// GenerateRandomNumberFromUUID 通过 UUID 生成纯数字的随机字符串
func GenerateRandomNumberFromUUID(length int) string {
	// 生成一个新的 UUID
	uuidValue := uuid.New()

	// 将 UUID 转换为一个大整数
	num := new(big.Int)
	num.SetBytes(uuidValue[:])

	// 将大整数转换为字符串
	randomStr := num.String()

	// 截取最后的 'length' 位数字
	if len(randomStr) > length {
		randomStr = randomStr[len(randomStr)-length:]
	}

	return randomStr
}

func run() {
	go func() {
		for {
			Update()
			time.Sleep(1 * time.Second)
		}
	}()
}

func Update() {
	fmt.Println("update")
}
