// 有限状态机
package fsm

import (
	"fmt"
	"log"
	"sync"
)

/** 四个概念
* 1. 状态 State
* 2. 事件 Event
* 3. 动作 Action
* 4. 转换 Transition
**/
type State uint8

type Event string

type Action func(from State, event Event, to State) error

type Transition struct {
	From      State          `desc:"旧状态"`
	Event     Event          `desc:"事件"`
	To        State          `desc:"新状态"`
	Action    Action         `desc:"动作"`
	Processor EventProcessor `desc:"处理器"`
}

// StateGraph 状态机图表
type StateGraph struct {
	name        string                         // 状态机名称
	start       State                          // 开始状态
	end         []State                        // 结束状态
	states      map[State]string               // 状态集合
	transitions map[State]map[Event]Transition // 转变器集合
}

func (g *StateGraph) IsEnd(state State) bool {
	for _, end := range g.end {
		if state == end {
			return true
		}
	}
	return false
}

/** 监听处理器接口
* 1. ExitOldState 退出旧状态: 状态离开就状态之前执行
* 2. EnterNewState 进入新状态: 状态进入新状态之后执行
**/
type EventProcessor interface {
	ExitOldState(from, to State) error
	EnterNewState(to State, event Event) error
}

// 每个状态机都需要定义一个默认的处理器 Processor，并且每个转变器 Transition 也可以自定义自己的处理器，注意，状态机和转变器的 处理器不是覆盖关系，而是先后执行的关系。
type StateMachine struct {
	locker    sync.Mutex     // 排他锁
	Processor EventProcessor // 默认处理器
	Graph     *StateGraph    // 状态机图表
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		Graph: &StateGraph{
			states:      make(map[State]string),
			transitions: make(map[State]map[Event]Transition),
		},
	}
}

func (s *StateMachine) SetName(name string) *StateMachine {
	s.Graph.name = name
	return s
}

func (s *StateMachine) SetStart(start State) *StateMachine {
	s.Graph.start = start
	return s
}

func (s *StateMachine) SetEnd(end []State) *StateMachine {
	s.Graph.end = end
	return s
}

func (s *StateMachine) SetStates(states map[State]string) *StateMachine {
	s.Graph.states = states
	return s
}

func (s *StateMachine) SetTransitions(transitions map[State]map[Event]Transition) *StateMachine {
	s.Graph.transitions = transitions
	return s
}

func (s *StateMachine) GetStateDesc(state State) string {
	return fmt.Sprintf("%s(%d)", s.Graph.states[state], state)
}

/** 状态机 StateMachine 核心方法
* 1. 状态及事件检测
* 2. 执行状态机的处理器的 ExitOldState 方法
* 3. 检查转变器是否定义了处理器，如果定义了，执行该处理器的 ExitOldState 方法
* 4. 执行转变器定义的 Action
* 5. 执行状态机的处理器的 EnterNewState 方法
* 6. 检查转变器是否定义了处理器，如果定义了，执行该处理器的 EnterNewState 方法
* 7. 执行完毕
**/
func (s *StateMachine) Run(from State, event Event) (State, error) {
	log.Printf("状态流转开始，旧状态：%s，事件：%s\n", s.GetStateDesc(from), event)
	// 检查旧状态是否存在
	if _, ok := s.Graph.states[from]; !ok {
		return 0, fmt.Errorf("旧状态不存在：%d", from)
	}
	// 检查旧状态是否已到最终状态
	if s.Graph.IsEnd(from) {
		return 0, fmt.Errorf("已到最终状态，无法流转")
	}
	// 检查状态与事件是否匹配
	transition, ok := s.Graph.transitions[from][event]
	if !ok {
		return 0, fmt.Errorf("未设置事件转换器")
	}
	to := transition.To
	// 加锁
	s.locker.Lock()
	// 执行完成后解锁
	defer s.locker.Unlock()

	// 执行状态机处理器，退出旧状态
	_ = s.Processor.ExitOldState(from, to)
	// 如果当前转变器设置了处理器，则执行处理器的退出旧状态
	if transition.Processor != nil {
		_ = transition.Processor.ExitOldState(from, to)
	}
	// 执行转变器动作
	_ = transition.Action(from, event, to)
	// 执行转变器处理器，进入新状态的方法
	_ = s.Processor.EnterNewState(to, event)
	// 如果当前转变器设置了处理器，则执行处理器的进入新状态的方法
	if transition.Processor != nil {
		_ = transition.Processor.EnterNewState(to, event)
	}
	return to, nil
}

// 状态：待支付，待确认，已支付，已取消
// 事件：支付，支付确认，取消
/** 1. 主订单状态流转
* 1.3 待支付 -(支付)-> 待确认
* 1.1 待支付 -(支付确认)-> 已支付
* 1.2 待支付 -(取消)-> 已取消
* 1.4 待确认 -(支付确认)-> 待确认
**/

// State 主订单状态
const (
	// StateWaitPay 待支付
	StateWaitPay State = iota
	// StateWaitConfirm 待确认
	StateWaitConfirm
	// StatePayied 已支付
	StatePayied
	// StateCanceled 已取消
	StateCanceled
)

// Event 主订单事件
const (
	// EventPay 支付
	EventPay Event = "pay"
	// EventPayConfirm 支付确认
	EventPayConfirm Event = "pay_confirm"
	// EventCancel 取消
	EventCancel Event = "cancel"
)

func MainPay(from State, event Event, to State) error {
	log.Printf("主订单支付，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func MainPayConfirm(from State, event Event, to State) error {
	log.Printf("主订单支付确认，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func MainCancel(from State, event Event, to State) error {
	log.Printf("主订单取消，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

// 主订单转变器
var transitions = map[State]map[Event]Transition{
	StateWaitPay: {
		EventPay: {
			From:   StateWaitPay,
			Action: MainPay,
			To:     StateWaitConfirm,
			Event:  EventPay,
		},
		EventCancel: {
			From:   StateWaitPay,
			Action: MainCancel,
			To:     StateCanceled,
			Event:  EventCancel,
		},
		EventPayConfirm: {
			From:   StateWaitPay,
			Action: MainPayConfirm,
			To:     StatePayied,
			Event:  EventPayConfirm,
		},
	},
	StateWaitConfirm: {
		EventPayConfirm: {
			From:   StateWaitConfirm,
			Action: MainPayConfirm,
			To:     StatePayied,
			Event:  EventPayConfirm,
		},
	},
}

var mainStates = map[State]string{
	StateWaitPay:     "wait_pay",
	StateWaitConfirm: "wait_confirm",
	StatePayied:      "payied",
	StateCanceled:    "canceled",
}

// 主订单状态机
var mainStateMachine = NewStateMachine().
	SetName("主订单状态机").
	SetEnd([]State{StateCanceled}).
	SetStart(StateWaitPay).
	SetTransitions(transitions).
	SetStates(mainStates)

// 状态：待支付，待确认，待发货，待收货，售后中-退款，售后中-退货退款， 已取消，已签收，已完成
// 事件：支付，支付确认，发货，签收，申请退款，申请退货退款，取消，取消售后, 售后完成，订单完成
/** 2. 子订单状态流转
* 2.1 待支付 -(支付)-> 待确认
* 2.2 待支付 -(取消)-> 已取消
* 2.4 待支付 -(支付确认)-> 待发货
* 2.3 待确认 -(支付确认)-> 待发货
* 2.5 待发货 -(发货)-> 待收货
* 2.6 待发货 -(申请退款)-> 售后中-退款
* 2.7 售后中-退款 -(售后完成)-> 已完成
* 2.7 售后中-退款 -(取消售后)-> 待发货
* 2.8 待收货 -(签收)-> 已签收
* 2.9 已签收 -(订单完成)-> 已完成
* 2.10 已签收 -(申请退货退款)-> 售后中-退货退款
* 2.11 售后中-退货退款 -(售后完成)-> 已完成
* 2.12 售后中-退货退款 -(取消售后)-> 已签收
**/

// State 子订单状态
const (
	// StateSubWaitPay 待支付
	StateSubWaitPay State = iota
	// StateSubWaitConfirm 待确认
	StateSubWaitConfirm
	// StateSubWaitShip 待发货
	StateSubWaitShip
	// StateSubWaitReceive 待收货
	StateSubWaitReceive
	// StateSubAfterSaleRefund 售后中-退款
	StateSubAfterSaleRefund
	// StateSubAfterSaleRefund 售后中-退货退款
	StateSubAfterSaleRefundAndReturn
	// StateSubCanceled 已取消
	StateSubCanceled
	// StateSubReceived 已签收
	StateSubReceived
	// StateSubCompleted 已完成
	StateSubCompleted
)

// Event 子订单事件
const (
	// EventSubPay 支付
	EventSubPay Event = "pay"
	// EventSubPayConfirm 支付确认
	EventSubPayConfirm Event = "pay_confirm"
	// EventSubShip 发货
	EventSubShip Event = "ship"
	// EventSubReceive 签收
	EventSubReceive Event = "receive"
	// EventSubRefund 申请退款
	EventSubRefund Event = "refund"
	// EventSubRefundAndReturn 申请退货退款
	EventSubRefundAndReturn Event = "refund_return"
	// EventSubCancel 取消
	EventSubCancel Event = "cancel"
	// EventSubCancelAfterSale 取消售后
	EventSubCancelAfterSale Event = "cancel_after_sale"
	// EventSubAfterSaleComplete 售后完成
	EventSubAfterSaleComplete Event = "after_sale_complete"
	// EventSubComplete 订单完成
	EventSubComplete Event = "complete"
)

func SubPay(from State, event Event, to State) error {
	log.Printf("子订单支付，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func SubPayConfirm(from State, event Event, to State) error {
	log.Printf("子订单支付确认，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func SubShip(from State, event Event, to State) error {
	log.Printf("子订单发货，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func SubReceive(from State, event Event, to State) error {
	log.Printf("子订单签收，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func SubRefund(from State, event Event, to State) error {
	log.Printf("子订单退款，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func SubRefundAndReturn(from State, event Event, to State) error {
	log.Printf("子订单退货退款，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func SubCancel(from State, event Event, to State) error {
	log.Printf("子订单取消，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func SubCancelAfterSale(from State, event Event, to State) error {
	log.Printf("子订单取消售后，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func SubAfterSaleComplete(from State, event Event, to State) error {
	log.Printf("子订单售后完成，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func SubComplete(from State, event Event, to State) error {
	log.Printf("子订单完成，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

// 子订单转变器
var subTransitions = map[State]map[Event]Transition{
	StateSubWaitPay: {
		EventSubPay: {
			From:   StateSubWaitPay,
			Action: SubPay,
			To:     StateSubWaitConfirm,
			Event:  EventSubPay,
		},
		EventSubCancel: {
			From:   StateSubWaitPay,
			Action: SubCancel,
			To:     StateSubCanceled,
			Event:  EventSubCancel,
		},
		EventSubPayConfirm: {
			From:   StateSubWaitPay,
			Action: SubPayConfirm,
			To:     StateSubWaitShip,
			Event:  EventSubPayConfirm,
		},
	},
	StateSubWaitConfirm: {
		EventSubPayConfirm: {
			From:   StateSubWaitConfirm,
			Action: SubPayConfirm,
			To:     StateSubWaitShip,
			Event:  EventSubPayConfirm,
		},
	},
	StateSubWaitShip: {
		EventSubShip: {
			From:   StateSubWaitShip,
			Action: SubShip,
			To:     StateSubWaitReceive,
			Event:  EventSubShip,
		},
		EventSubRefund: {
			From:   StateSubWaitShip,
			Action: SubRefund,
			To:     StateSubAfterSaleRefund,
			Event:  EventSubRefund,
		},
	},
	StateSubWaitReceive: {
		EventSubReceive: {
			From:   StateSubWaitReceive,
			Action: SubReceive,
			To:     StateSubReceived,
			Event:  EventSubReceive,
		},
	},
	StateSubAfterSaleRefund: {
		EventSubAfterSaleComplete: {
			From:   StateSubAfterSaleRefund,
			Action: SubAfterSaleComplete,
			To:     StateSubCompleted,
			Event:  EventSubAfterSaleComplete,
		},
		EventSubCancelAfterSale: {
			From:   StateSubAfterSaleRefund,
			Action: SubCancelAfterSale,
			To:     StateSubWaitShip,
			Event:  EventSubCancelAfterSale,
		},
	},
	StateSubAfterSaleRefundAndReturn: {
		EventSubAfterSaleComplete: {
			From:   StateSubAfterSaleRefundAndReturn,
			Action: SubAfterSaleComplete,
			To:     StateSubCompleted,
			Event:  EventSubAfterSaleComplete,
		},
		EventSubCancelAfterSale: {
			From:   StateSubAfterSaleRefundAndReturn,
			Action: SubCancelAfterSale,
			To:     StateSubReceived,
			Event:  EventSubCancelAfterSale,
		},
	},
	StateSubReceived: {
		EventSubComplete: {
			From:   StateSubReceived,
			Action: SubComplete,
			To:     StateSubCompleted,
			Event:  EventSubComplete,
		},
		EventSubRefundAndReturn: {
			From:   StateSubReceived,
			Action: SubRefundAndReturn,
			To:     StateSubAfterSaleRefundAndReturn,
			Event:  EventSubRefundAndReturn,
		},
	},
}

var subStates = map[State]string{
	StateSubWaitPay:                  "wait_pay",
	StateSubWaitConfirm:              "wait_confirm",
	StateSubWaitShip:                 "wait_ship",
	StateSubWaitReceive:              "wait_receive",
	StateSubAfterSaleRefund:          "after_sale_refund",
	StateSubAfterSaleRefundAndReturn: "after_sale_refund_return",
	StateSubCanceled:                 "canceled",
	StateSubReceived:                 "received",
	StateSubCompleted:                "completed",
}

// 子订单状态机
var subStateMachine = NewStateMachine().
	SetName("子订单状态机").
	SetEnd([]State{StateSubCompleted, StateSubCanceled}).
	SetStart(StateSubWaitPay).
	SetTransitions(subTransitions).
	SetStates(subStates)

// 状态：待审批，已驳回，已通过，已取消， 退货中，待收货，退款中，已完成
// 事件：驳回，通过，取消，发货，签收，退款完成，等待用户寄回
/** 3. 售后状态流转
* 3.1 待审批 -(通过)-> 已通过
* 3.2 待审批 -(驳回)-> 已驳回
* 3.3 待审批 -(取消)-> 已取消
* 3.4 已通过 -(等待用户寄回)-> 退货中
* 3.5 已通过 -(提交退款申请[未发货订单])-> 退款中
* 3.6 已通过 -(取消)-> 已取消
* 3.7 退货中 -(发货)-> 待收货
* 3.8 待收货 -(签收)-> 退款中
* 3.9 退款中 -(退款完成)-> 已完成
**/

// State 售后状态
const (
	// StateAfterSaleWaitReview 待审批
	StateAfterSaleWaitReview State = iota
	// StateAfterSaleReject 已驳回
	StateAfterSaleReject
	// StateAfterSalePass 已通过
	StateAfterSalePass
	// StateAfterSaleCancel 已取消
	StateAfterSaleCancel
	// StateAfterSaleReturn 退货中
	StateAfterSaleReturn
	// StateAfterSaleWaitReceive 待收货
	StateAfterSaleWaitReceive
	// StateAfterSaleRefund 退款中
	StateAfterSaleRefund
	// StateAfterSaleComplete 已完成
	StateAfterSaleComplete
)

// Event 售后事件
const (
	// EventAfterSaleReject 驳回
	EventAfterSaleReject Event = "reject"
	// EventAfterSalePass 通过
	EventAfterSalePass Event = "pass"
	// EventAfterSaleCancel 取消
	EventAfterSaleCancel Event = "cancel"
	// EventAfterSaleShip 发货
	EventAfterSaleShip Event = "ship"
	// EventAfterSaleReceive 签收
	EventAfterSaleReceive Event = "receive"
	// EventAfterSaleRefund 退款完成
	EventAfterSaleRefund Event = "refund"
	// EventAfterSaleReturn 等待用户寄回
	EventAfterSaleReturn Event = "return"
	// EventRefundReq 退款申请
	EventRefundReq Event = "refund_req"
)

func AfterSaleReject(from State, event Event, to State) error {
	log.Printf("售后驳回，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func AfterSalePass(from State, event Event, to State) error {
	log.Printf("售后通过，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func AfterSaleCancel(from State, event Event, to State) error {
	log.Printf("售后取消，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func AfterSaleShip(from State, event Event, to State) error {
	log.Printf("售后发货，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func AfterSaleReceive(from State, event Event, to State) error {
	log.Printf("售后签收，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func AfterSaleRefund(from State, event Event, to State) error {
	log.Printf("售后退款，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func AfterSaleReturn(from State, event Event, to State) error {
	log.Printf("售后退货，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func AfterSaleComplete(from State, event Event, to State) error {
	log.Printf("售后完成，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

func AfterSaleRefundReq(from State, event Event, to State) error {
	log.Printf("售后退款申请，旧状态：%s，新状态：%s\n", from, to)
	return nil
}

// 售后转变器
var afterSaleTransitions = map[State]map[Event]Transition{
	StateAfterSaleWaitReview: {
		EventAfterSaleReject: {
			From:   StateAfterSaleWaitReview,
			Action: AfterSaleReject,
			To:     StateAfterSaleReject,
			Event:  EventAfterSaleReject,
		},
		EventAfterSalePass: {
			From:   StateAfterSaleWaitReview,
			Action: AfterSalePass,
			To:     StateAfterSalePass,
			Event:  EventAfterSalePass,
		},
	},
	StateAfterSalePass: {
		EventAfterSaleReturn: {
			From:   StateAfterSalePass,
			Action: AfterSaleReturn,
			To:     StateAfterSaleReturn,
			Event:  EventAfterSaleReturn,
		},
		EventRefundReq: {
			From:   StateAfterSalePass,
			Action: AfterSaleRefundReq,
			To:     StateAfterSaleRefund,
			Event:  EventRefundReq,
		},
		EventAfterSaleCancel: {
			From:   StateAfterSalePass,
			Action: AfterSaleCancel,
			To:     StateAfterSaleCancel,
			Event:  EventAfterSaleCancel,
		},
	},
	StateAfterSaleReturn: {
		EventAfterSaleShip: {
			From:   StateAfterSaleReturn,
			Action: AfterSaleShip,
			To:     StateAfterSaleWaitReceive,
			Event:  EventAfterSaleShip,
		},
	},
	StateAfterSaleWaitReceive: {
		EventAfterSaleReceive: {
			From:   StateAfterSaleWaitReceive,
			Action: AfterSaleReceive,
			To:     StateAfterSaleRefund,
			Event:  EventAfterSaleReceive,
		},
	},
	StateAfterSaleRefund: {
		EventAfterSaleRefund: {
			From:   StateAfterSaleRefund,
			Action: AfterSaleRefund,
			To:     StateAfterSaleComplete,
			Event:  EventAfterSaleRefund,
		},
	},
}

// 状态机状态
var afterSaleStates = map[State]string{
	StateAfterSaleWaitReview:  "wait_review",
	StateAfterSaleReject:      "reject",
	StateAfterSalePass:        "pass",
	StateAfterSaleCancel:      "cancel",
	StateAfterSaleReturn:      "return",
	StateAfterSaleWaitReceive: "wait_receive",
	StateAfterSaleRefund:      "refund",
	StateAfterSaleComplete:    "complete",
}

// 售后状态机
var afterSaleStateMachine = NewStateMachine().
	SetName("售后状态机").
	SetEnd([]State{StateAfterSaleComplete, StateAfterSaleCancel, StateAfterSaleReject}).
	SetStart(StateAfterSaleWaitReview).
	SetTransitions(afterSaleTransitions).
	SetStates(afterSaleStates)
