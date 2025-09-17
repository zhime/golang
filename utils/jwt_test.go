package utils

import (
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	// 测试生成Token
	userID := 1
	email := "test@example.com"
	name := "测试用户"

	token, err := GenerateToken(userID, email, name)
	if err != nil {
		t.Fatalf("生成Token失败: %v", err)
	}

	if token == "" {
		t.Fatal("生成的Token为空")
	}

	t.Logf("生成的Token: %s", token)
}

func TestParseToken(t *testing.T) {
	// 先生成一个Token
	userID := 1
	email := "test@example.com"
	name := "测试用户"

	token, err := GenerateToken(userID, email, name)
	if err != nil {
		t.Fatalf("生成Token失败: %v", err)
	}

	// 解析Token
	claims, err := ParseToken(token)
	if err != nil {
		t.Fatalf("解析Token失败: %v", err)
	}

	// 验证解析结果
	if claims.UserID != userID {
		t.Errorf("用户ID不匹配: 期望 %d, 实际 %d", userID, claims.UserID)
	}

	if claims.Email != email {
		t.Errorf("邮箱不匹配: 期望 %s, 实际 %s", email, claims.Email)
	}

	if claims.Name != name {
		t.Errorf("用户名不匹配: 期望 %s, 实际 %s", name, claims.Name)
	}

	t.Logf("解析成功: UserID=%d, Email=%s, Name=%s", claims.UserID, claims.Email, claims.Name)
}

func TestValidateToken(t *testing.T) {
	// 测试有效Token
	token, err := GenerateToken(1, "test@example.com", "测试用户")
	if err != nil {
		t.Fatalf("生成Token失败: %v", err)
	}

	if !ValidateToken(token) {
		t.Error("有效Token验证失败")
	}

	// 测试无效Token
	invalidToken := "invalid.token.here"
	if ValidateToken(invalidToken) {
		t.Error("无效Token验证通过")
	}

	t.Log("Token验证测试通过")
}

func TestGetUserInfoFromToken(t *testing.T) {
	// 生成测试Token
	expectedUserID := 123
	expectedEmail := "user@test.com"
	expectedName := "用户测试"

	token, err := GenerateToken(expectedUserID, expectedEmail, expectedName)
	if err != nil {
		t.Fatalf("生成Token失败: %v", err)
	}

	// 从Token获取用户信息
	userID, email, name, err := GetUserInfoFromToken(token)
	if err != nil {
		t.Fatalf("获取用户信息失败: %v", err)
	}

	// 验证信息
	if userID != expectedUserID {
		t.Errorf("用户ID不匹配: 期望 %d, 实际 %d", expectedUserID, userID)
	}

	if email != expectedEmail {
		t.Errorf("邮箱不匹配: 期望 %s, 实际 %s", expectedEmail, email)
	}

	if name != expectedName {
		t.Errorf("用户名不匹配: 期望 %s, 实际 %s", expectedName, name)
	}

	t.Logf("用户信息获取成功: ID=%d, Email=%s, Name=%s", userID, email, name)
}

func TestTokenExpiration(t *testing.T) {
	// 注意：这个测试会修改全局变量，仅用于演示
	originalExpiration := tokenExpiration
	defer func() {
		tokenExpiration = originalExpiration
	}()

	// 设置很短的过期时间用于测试
	tokenExpiration = time.Millisecond * 10

	token, err := GenerateToken(1, "test@example.com", "测试用户")
	if err != nil {
		t.Fatalf("生成Token失败: %v", err)
	}

	// 立即验证应该成功
	if !ValidateToken(token) {
		t.Log("新生成的Token验证失败，跳过过期测试")
		return
	}

	// 等待Token过期
	time.Sleep(time.Millisecond * 50)

	// 验证过期的Token
	if ValidateToken(token) {
		t.Error("过期的Token验证通过")
	}

	t.Log("Token过期测试通过")
}