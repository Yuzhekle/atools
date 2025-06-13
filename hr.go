package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(GenerateHRSplOrderNo([]float64{112.2}))
	fmt.Println(GenerateHRSplOrderNo("3344"))
	wds := WeekDates{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(wds.Value())
	fmt.Println(TInt(127).ToWeekDates())
}

const (
	HR_TYPE_ZHIQIAN = iota + 1 // 1 直签
	HR_TYPE_MANUAL             // 2 手工
	HR_TYPE_B2B                // 3 B2B
)

// T（BCDEFHIJKOPQRST）
const (
	HR_SPL_ORDER_PREFIX = "T" // hold房供应商订单号前缀
	HR_RANDOM_CHARS     = "BCDEFHIJKOPQRST"
)

// var HrRandomChar = []string{"B", "C", "D", "E", "F", "H", "I", "J", "K", "O", "P", "Q", "R", "S", "T"}

// hold房类型
type HrType int

func (ht HrType) String() string {
	switch ht {
	case HR_TYPE_ZHIQIAN:
		return "直签"
	case HR_TYPE_B2B:
		return "B2B"
	case HR_TYPE_MANUAL:
		return "手工"
	default:
		return "未知类型"
	}
}

// hold房邮件状态
type HrEmailStatus int

const (
	HR_EMAIL_STATUS_PENDING     HrEmailStatus = iota // 0 初始状态
	HR_EMAIL_STATUS_SENDED                           // 1 已发预定邮件
	HR_EMAIL_STATUS_CANCEL                           // 2 已发取消邮件
	HR_EMAIL_STATUS_CHANGE_NAME                      // 3 已发改名邮件
	HR_EMAIL_STATUS_CHANGE_DATE                      // 4 已发改期邮件
	HR_EMAIL_STATUS_FAILED                           // 5 邮件发送失败
)

func (es HrEmailStatus) String() string {
	switch es {
	case HR_EMAIL_STATUS_PENDING:
		return "邮件状态未设置"
	case HR_EMAIL_STATUS_SENDED:
		return "已发预定邮件"
	case HR_EMAIL_STATUS_CANCEL:
		return "已发取消邮件"
	case HR_EMAIL_STATUS_CHANGE_NAME:
		return "已发改名邮件"
	case HR_EMAIL_STATUS_CHANGE_DATE:
		return "已发改期邮件"
	case HR_EMAIL_STATUS_FAILED:
		return "邮件发送失败"
	default:
		return "未知邮件状态"
	}
}

// hold房酒店状态
type HrHotelStatus int

const (
	HR_HOTEL_STATUS_PENDING            HrHotelStatus = iota // 0 初始状态
	HR_HOTEL_STATUS_CANCEL                                  // 1 酒店已取消
	HR_HOTEL_STATUS_CONFIRM                                 // 2 酒店已预定
	HR_HOTEL_STATUS_FAILED                                  // 3 酒店预定失败
	HR_HOTEL_STATUS_CHANGE_NAME                             // 4 酒店已改名
	HR_HOTEL_STATUS_CHANGE_DATE                             // 5 酒店已改期
	HR_HOTEL_STATUS_WAIT_REPLY                              // 6 酒店待回复
	HR_HOTEL_STATUS_REJECT_CHANGE_NAME                      // 7 酒店拒绝改名
	HR_HOTEL_STATUS_REJECT_CHANGE_DATE                      // 8 酒店拒绝改期
	HR_HOTEL_STATUS_REJECT_CANCEL                           // 9 酒店拒绝取消
)

func (hs HrHotelStatus) String() string {
	switch hs {
	case HR_HOTEL_STATUS_PENDING:
		return "酒店状态未设置"
	case HR_HOTEL_STATUS_CANCEL:
		return "酒店已取消"
	case HR_HOTEL_STATUS_CONFIRM:
		return "酒店已预定"
	case HR_HOTEL_STATUS_FAILED:
		return "酒店预定失败"
	case HR_HOTEL_STATUS_CHANGE_NAME:
		return "酒店已改名"
	case HR_HOTEL_STATUS_CHANGE_DATE:
		return "酒店已改期"
	case HR_HOTEL_STATUS_WAIT_REPLY:
		return "酒店待回复"
	case HR_HOTEL_STATUS_REJECT_CHANGE_NAME:
		return "酒店拒绝改名"
	case HR_HOTEL_STATUS_REJECT_CHANGE_DATE:
		return "酒店拒绝改期"
	case HR_HOTEL_STATUS_REJECT_CANCEL:
		return "酒店拒绝取消"
	default:
		return "未知酒店状态"
	}
}

// hold房状态
type HrStatus int

const (
	HR_STATUS_OFF      HrStatus = iota // 0 未上架
	HR_STATUS_ON                       // 1 已上架
	HR_STATUS_OCCUPIED                 // 2 已占用
	HR_STATUS_EXPIRED                  // 3 已过期
)

func (hs HrStatus) String() string {
	switch hs {
	case HR_STATUS_OFF:
		return "未上架"
	case HR_STATUS_ON:
		return "已上架"
	case HR_STATUS_OCCUPIED:
		return "已占用"
	case HR_STATUS_EXPIRED:
		return "已过期"
	default:
		return "未知状态"
	}
}

// hold房供应商订单状态
type HrOrderStatus int

const (
	HR_SUPPLIER_ORDER_STATUS_PENDING HrOrderStatus = iota + 1 // 1 待确认
	HR_SUPPLIER_ORDER_STATUS_CONFIRM                          // 2 已确认
	HR_SUPPLIER_ORDER_STATUS_CANCEL                           // 3 已取消
)

func (hs HrOrderStatus) String() string {
	switch hs {
	case HR_SUPPLIER_ORDER_STATUS_PENDING:
		return "待确认"
	case HR_SUPPLIER_ORDER_STATUS_CANCEL:
		return "已取消"
	case HR_SUPPLIER_ORDER_STATUS_CONFIRM:
		return "已确认"
	default:
		return "未知状态"
	}
}

// 生成 hold房订单号
func GenerateHRSplOrderNo(oriOrderNo any) string {
	// 随机种子
	seed := time.Now().UnixNano()
	// 随机源
	source := rand.NewSource(seed)
	// 创建 rand 实例
	random := rand.New(source)
	randomChar := HR_RANDOM_CHARS[random.Intn(len(HR_RANDOM_CHARS))]
	return fmt.Sprintf("%s%s%v", HR_SPL_ORDER_PREFIX, string(randomChar), oriOrderNo)
}

type WeekDate int

type WeekDates []WeekDate

func (w WeekDates) Value() int {
	var result int
	for _, v := range w {
		result |= 1 << uint(v)
	}
	return result
}

type TInt int

func (t TInt) ToWeekDates() WeekDates {
	var result WeekDates
	for i := 0; i < 7; i++ {
		if t&(1<<uint(i)) > 0 {
			result = append(result, WeekDate(i))
		}
	}
	return result
}
