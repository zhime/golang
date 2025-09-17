package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	// bcrypt成本因子，值越高安全性越高但性能越慢
	// 推荐值：12
	defaultCost = 12
)

// HashPassword 使用bcrypt加密密码
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("密码不能为空")
	}

	// 生成密码哈希
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), defaultCost)
	if err != nil {
		return "", fmt.Errorf("密码加密失败: %w", err)
	}

	return string(hashedBytes), nil
}

// VerifyPassword 验证密码是否正确
func VerifyPassword(hashedPassword, password string) bool {
	if hashedPassword == "" || password == "" {
		return false
	}

	// 比较密码
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// PasswordStrength 检查密码强度
type PasswordStrength struct {
	IsValid        bool     `json:"is_valid"`
	MinLength      bool     `json:"min_length"`      // 最小长度8位
	HasLower       bool     `json:"has_lower"`       // 包含小写字母
	HasUpper       bool     `json:"has_upper"`       // 包含大写字母
	HasDigit       bool     `json:"has_digit"`       // 包含数字
	HasSpecial     bool     `json:"has_special"`     // 包含特殊字符
	Suggestions    []string `json:"suggestions"`     // 改进建议
}

// CheckPasswordStrength 检查密码强度
func CheckPasswordStrength(password string) PasswordStrength {
	strength := PasswordStrength{
		Suggestions: make([]string, 0),
	}

	// 检查最小长度
	if len(password) >= 8 {
		strength.MinLength = true
	} else {
		strength.Suggestions = append(strength.Suggestions, "密码长度至少8位")
	}

	// 检查字符类型
	for _, char := range password {
		switch {
		case char >= 'a' && char <= 'z':
			strength.HasLower = true
		case char >= 'A' && char <= 'Z':
			strength.HasUpper = true
		case char >= '0' && char <= '9':
			strength.HasDigit = true
		case char >= 33 && char <= 126 && !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z') && !(char >= '0' && char <= '9'):
			strength.HasSpecial = true
		}
	}

	// 添加建议
	if !strength.HasLower {
		strength.Suggestions = append(strength.Suggestions, "密码应包含小写字母")
	}
	if !strength.HasUpper {
		strength.Suggestions = append(strength.Suggestions, "密码应包含大写字母")
	}
	if !strength.HasDigit {
		strength.Suggestions = append(strength.Suggestions, "密码应包含数字")
	}
	if !strength.HasSpecial {
		strength.Suggestions = append(strength.Suggestions, "密码应包含特殊字符")
	}

	// 判断密码是否有效（满足所有基本要求）
	strength.IsValid = strength.MinLength && strength.HasLower && strength.HasUpper && strength.HasDigit && strength.HasSpecial

	return strength
}

// IsPasswordValid 简单的密码有效性检查
func IsPasswordValid(password string) bool {
	strength := CheckPasswordStrength(password)
	return strength.IsValid
}