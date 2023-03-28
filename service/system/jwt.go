package system

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"project/global"
	"project/model/system"
	"project/model/system/request"
	"time"
)

type JwtService struct{}

var authorityService = AuthorityService{}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	err := global.TPA_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: JoinInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JoinInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.TPA_DB.Create(&jwtList).Error
	return
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string

func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.TPA_REDIS.Get(context.Background(), userName).Result()
	return err, redisJWT
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.TPA_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.TPA_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: GetRedisUserInfo
//@description: 从redis获取用户信息
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisUserInfo(uuid string, userInfo request.UserCache) (err error) {
	_, err = global.TPA_REDIS.Pipelined(context.Background(), func(rdb redis.Pipeliner) error {
		rdb.HSet(context.Background(), uuid, "id", userInfo.ID)
		rdb.HSet(context.Background(), uuid, "uuid", userInfo.UUID)
		rdb.HSet(context.Background(), uuid, "deptId", userInfo.DeptId)
		authorityJson, _ := json.Marshal(userInfo.AuthorityId)
		rdb.HSet(context.Background(), uuid, "authorityId", authorityJson)
		return nil
	})
	return err
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: SetRedisUserInfo
//@description: 从redis获取用户信息
//@param: uuid string
//@return: err error

func (jwtService *JwtService) GetRedisUserInfo(uuid string) (redisUserInfo request.UserCache, err error) {
	var userInfoRedis request.UserCacheRedis
	err = global.TPA_REDIS.HGetAll(context.Background(), uuid).Scan(&userInfoRedis)
	if err != nil || userInfoRedis.ID == 0 {
		return redisUserInfo, errors.New("查询用户缓存失败")
	}
	redisUserInfo.UUID = uuid
	redisUserInfo.ID = userInfoRedis.ID
	redisUserInfo.DeptId = userInfoRedis.DeptId
	var authorityId []uint
	err = json.Unmarshal(userInfoRedis.AuthorityId, &authorityId)
	redisUserInfo.AuthorityId = authorityId
	err, redisUserInfo.Authority = authorityService.GetAuthorityInfoByIDs(redisUserInfo.AuthorityId)
	if err != nil || userInfoRedis.ID == 0 {
		return redisUserInfo, errors.New("查询用户缓存失败")
	}
	return
}

//@author: [liaofan](https://github.com/liaofan-liaofan)
//@function: DelRedisUserInfo
//@description: 从redis获取用户信息
//@param: uuid string
//@return: err error

func (jwtService *JwtService) DelRedisUserInfo(uuid string) (err error) {
	return global.TPA_REDIS.HDel(context.Background(), uuid).Err()
}
