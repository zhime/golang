package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "testPassword123"
	
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("密码哈希失败: %v", err)
	}

	if hashedPassword == "" {
		t.Fatal("哈希密码为空")
	}

	if hashedPassword == password {
		t.Error("哈希密码与原密码相同")
	}

	t.Logf("密码哈希成功: %s", hashedPassword)
}

func TestVerifyPassword(t *testing.T) {
	password := "testPassword123"
	wrongPassword := "wrongPassword456"

	// 先哈希密码
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("密码哈希失败: %v", err)
	}

	// 验证正确密码
	if !VerifyPassword(hashedPassword, password) {
		t.Error("正确密码验证失败")
	}

	// 验证错误密码
	if VerifyPassword(hashedPassword, wrongPassword) {
		t.Error("错误密码验证通过")
	}

	t.Log("密码验证测试通过")
}

func TestHashPasswordEmpty(t *testing.T) {
	// 测试空密码
	_, err := HashPassword("")
	if err == nil {
		t.Error("空密码应该返回错误")
	}

	t.Log("空密码测试通过")
}

func TestVerifyPasswordEmpty(t *testing.T) {
	// 测试空参数
	if VerifyPassword("", "password") {
		t.Error("空哈希值应该返回false")
	}

	if VerifyPassword("hashedPassword", "") {
		t.Error("空密码应该返回false")
	}

	if VerifyPassword("", "") {
		t.Error("空参数应该返回false")
	}

	t.Log("空参数验证测试通过")
}

func TestCheckPasswordStrength(t *testing.T) {
	tests := []struct {
		password    string
		expectedValid bool
		description string
	}{
		{"Abc123!@", true, "强密码"},
		{"password", false, "只有小写字母"},
		{"PASSWORD", false, "只有大写字母"},
		{"Password", false, "没有数字和特殊字符"},
		{"Password123", false, "没有特殊字符"},
		{"abc123", false, "太短且没有大写字母和特殊字符"},
		{"", false, "空密码"},
		{"Abc1", false, "太短"},
		{"VeryLongPasswordWith123", false, "没有特殊字符"},
		{"Abcd1234!", true, "符合所有要求"},
	}

	for _, test := range tests {
		strength := CheckPasswordStrength(test.password)
		if strength.IsValid != test.expectedValid {
			t.Errorf("密码 '%s' (%s): 期望有效性 %v, 实际 %v", 
				test.password, test.description, test.expectedValid, strength.IsValid)
		}

		t.Logf("密码 '%s' 强度检查: 有效=%v, 建议=%v", 
			test.password, strength.IsValid, strength.Suggestions)
	}
}

func TestIsPasswordValid(t *testing.T) {
	validPasswords := []string{
		"Abcd1234!",
		"MyPass123@",
		"Secure99#",
	}

	invalidPasswords := []string{
		"password",
		"123456",
		"ABCDEF",
		"Abc123",
		"short",
		"",
	}

	// 测试有效密码
	for _, password := range validPasswords {
		if !IsPasswordValid(password) {
			t.Errorf("有效密码 '%s' 被判定为无效", password)
		}
	}

	// 测试无效密码
	for _, password := range invalidPasswords {
		if IsPasswordValid(password) {
			t.Errorf("无效密码 '%s' 被判定为有效", password)
		}
	}

	t.Log("密码有效性测试通过")
}

func TestPasswordStrengthDetails(t *testing.T) {
	password := "TestPass123!"
	strength := CheckPasswordStrength(password)

	// 检查各项指标
	if !strength.MinLength {
		t.Error("应该满足最小长度要求")
	}

	if !strength.HasLower {
		t.Error("应该包含小写字母")
	}

	if !strength.HasUpper {
		t.Error("应该包含大写字母")
	}

	if !strength.HasDigit {
		t.Error("应该包含数字")
	}

	if !strength.HasSpecial {
		t.Error("应该包含特殊字符")
	}

	if !strength.IsValid {
		t.Error("应该是有效密码")
	}

	if len(strength.Suggestions) > 0 {
		t.Errorf("不应该有改进建议，但有: %v", strength.Suggestions)
	}

	t.Log("密码强度详情测试通过")
}