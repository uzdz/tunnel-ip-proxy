package config

// -------------- 标准参数

// 是否需要验证 0需要 1不需要
var NoAuth int64

// 用户访问限制
var UserConnectLimit map[string]string

// 授权用户列表（IP反转MAP）
var AuthListWithIpMap map[string]string

// 授权用户列表（用户名反转MAP）
var AuthListWithUNameMap map[string]string

// -------------- 标准代理参数

// IP切换时间间隔（秒）
var IpInterval int64

// 当日IP是否允许重复 0允许 1不允许
var IpRepeat int64
