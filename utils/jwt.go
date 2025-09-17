package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT配置
var (
	// JWT密钥，生产环境应该从环境变量或配置文件读取
	jwtSecret = []byte("your-secret-key-here-should-be-very-long-and-random")
	// Token过期时间
	tokenExpiration = time.Hour * 24 // 24小时
)

// JWTClaims JWT声明结构
type JWTClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(userID int, email, name string) (string, error) {
	// 创建声明
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "golang-auth-system",
		},
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	// 签名token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("生成token失败: %w", err)
	}

	return tokenString, nil
}

// ParseToken 解析JWT Token
func ParseToken(tokenString string) (*JWTClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("解析token失败: %w", err)
	}

	// 验证token并提取声明
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的token")
}

// ValidateToken 验证JWT Token是否有效
func ValidateToken(tokenString string) bool {
	_, err := ParseToken(tokenString)
	return err == nil
}

// RefreshToken 刷新JWT Token（如果token在有效期内）
func RefreshToken(tokenString string) (string, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	// 检查token是否即将过期（在1小时内过期则可以刷新）
	if time.Until(claims.ExpiresAt.Time) > time.Hour {
		return "", errors.New("token还未到刷新时间")
	}

	// 生成新token
	return GenerateToken(claims.UserID, claims.Email, claims.Name)
}

// GetUserInfoFromToken 从token中提取用户信息
func GetUserInfoFromToken(tokenString string) (userID int, email, name string, err error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return 0, "", "", err
	}

	return claims.UserID, claims.Email, claims.Name, nil
}