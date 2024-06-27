package model

import (
	"fmt"
)

type Role uint8

func (r Role) String() string {
	switch r {
	case RoleSuperAdmin:
		return "超級管理者"
	case RoleVenusManager:
		return "場館管理者"
	case RoleBoxManager:
		return "牌盒管理者"
	}
	return fmt.Sprintf("未定義的角色(%d)", r)
}

// in 常見的放左邊比較早 return
func (r Role) in(roles ...Role) bool {
	for _, role := range roles {
		if r == role {
			return true
		}
	}
	return false
}

// func (r Role) notIn(roles ...Role) bool {
// 	for _, role := range roles {
// 		if r == role {
// 			return false
// 		}
// 	}
// 	return true
// }

const (
	RoleSuperAdmin Role = iota
	RoleVenusManager
	RoleBoxManager
)

type Permission int8

const (
	PermissionCreateBox Permission = iota
	PermissionCreateUser
	PermissionUpdateUser
)

func (p Permission) String() string {
	switch p {
	case PermissionCreateUser:
		return "用戶創建"
	case PermissionUpdateUser:
		return "用戶更新"
	}
	return fmt.Sprintf("未定義的權限(%d)", p)
}

type PermissionChecker struct {
	permission Permission
	operator   Operator
	verifier   Operator
	checkFunc  []func(*PermissionChecker)
	err        error
}

type Operator struct {
	UserRole    Role
	UserID      int32
	ProductID   int32
	VenueID     int32
	AreaID      int32
	UserAccount string
	ProductName string
}

func (r Role) HasOperatePermission(p Permission) bool {
	switch p {
	case PermissionCreateUser, PermissionUpdateUser:
		return r.in(RoleVenusManager, RoleSuperAdmin)
	}
	return false
}

/*
func (r Role) HasOperatePermission(p Permission) bool {
	switch p {
	case PermissionCreateArea, PermissionUpdateArea:
		return r == RoleVenusManager
	case PermissionReadArea:
		return r != RoleSuperAdmin
	case PermissionCreateVenue, PermissionUpdateVenue:
		return r == RoleSuperAdmin
	case PermissionReadVenue:
		return r.in(RoleVenusManager, RoleSuperAdmin)
	case PermissionCreateUser, PermissionUpdateUser:
		return r.in(RoleVenusManager, RoleSuperAdmin)
	case PermissionReadUser:
		return true
	case PermissionReadBox:
		return r != RoleSuperAdmin
	case PermissionComfirmShuffleJob:
		return r == RoleOperator
	case PermissionUpdateBox:
		return r.in(RoleVenusManager, RoleBoxManager)
	case PermissionCreateBox:
		return r == RoleOperator
	case PermissionAddTransportJob:
		return r == RoleOperator
	case PermissionComfirmTransportJob:
		return r == RoleOperator
	case PermissionTemporaryToReady:
		return r == RoleSupervisor
	case PermissionAddRecycleJob:
		return r == RoleOperator
	case PermissionComfirmRecycleJob:
		return r == RoleOperator
	case PermissionAddShuffleJob:
		return r == RoleOperator
	case PermissionResetBox:
		return r == RoleVenusManager
	case PermissionResetBoxForMaintenance:
		return r == RoleVenusManager
	}
	return false
}

func (r Role) HasVerifyPermission(p Permission) bool {
	switch p { //nolint:exhaustive
	case PermissionReadBox:
		return r != RoleSuperAdmin
	case PermissionComfirmShuffleJob:
		return r.in(RoleBoxManager, RoleVenusManager)
	case PermissionCreateBox:
		return r.in(RoleBoxManager, RoleVenusManager)
	case PermissionAddTransportJob:
		return r.in(RoleBoxManager, RoleVenusManager)
	case PermissionComfirmTransportJob:
		return r == RoleSupervisor
	case PermissionAddRecycleJob:
		return r == RoleSupervisor
	case PermissionComfirmRecycleJob:
		return r.in(RoleBoxManager, RoleVenusManager)
	case PermissionAddShuffleJob:
		return r.in(RoleBoxManager, RoleVenusManager)
	}
	return false
}
*/
