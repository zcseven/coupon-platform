// Code generated by goctl. DO NOT EDIT.
package types

type RegisterReq struct {
	Email     string `json:"email"`
	Telephone string `json:"telepohone"`
	Username  string `json:"user_name"`
}

type RegisterResp struct {
	Uid   int64  `json:"uid"`
	Email string `json:"email"`
}

type LoginReq struct {
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

type LoginResp struct {
	Uid      int64  `json:"uid"`
	Username string `json:"user_name"`
	Headpic  string `json:"head_pic"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserReq struct {
	Uid       int64  `json:"uid"`
	Telephone string `json:"telephone,optional"`
}

type UserResp struct {
	Uid                 int64   `json:"uid"`
	Username            string  `json:"user_name"`
	Headpic             string  `json:"head_pic"`
	IntegralNum         int64   `json:"integral_num"`         // 签到积分量
	MembershipLevel     int64   `json:"membership_level"`     // 会员等级
	FansNum             int64   `json:"fans_num"`             // 粉丝数量
	WithdrawalAmount    float64 `json:"withdrawal_amount"`    // 提现金额
	RedEnvelope         float64 `json:"red_envelope"`         // 未到位的红包金额
	PromotionCommission float64 `json:"promotion_commission"` // 推广提成金额
	HistoryWithdrawal   float64 `json:"history_withdrawal"`   // 历史提现金额
}

type UserRightsListReq struct {
	Uid      int64 `json:"uid"`
	Pg       int64 `json:"pg"`
	PageSize int64 `json:"page_size"`
}

type UserRightsListResp struct {
}
