package request

import (
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"project/model/system"
)

// CustomClaims Custom claims structure
type CustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	BufferTime  int64
	AuthorityId uint
	jwt.StandardClaims
}

// UserCache User cache structure
type UserCache struct {
	UUID        string                `redis:"uuid"`
	ID          uint                  `redis:"id"`
	DeptId      uint                  `redis:"deptId"`
	AuthorityId []uint                `redis:"authorityId"`
	Authority   []system.SysAuthority `redis:"-"`
}

// UserCacheRedis User cache structure
type UserCacheRedis struct {
	ID          uint   `redis:"id"`
	DeptId      uint   `redis:"deptId"`
	AuthorityId []byte `redis:"authorityId"`
}
