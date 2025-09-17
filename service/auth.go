package service

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/zhime/golang/model"
	"github.com/zhime/golang/utils"
)

// AuthService 认证服务
type AuthService struct {
	users         []model.User           // 内存存储用户数据
	tokenBlacklist map[string]time.Time  // Token黑名单
	mutex         sync.RWMutex           // 读写锁
}

// 全局认证服务实例
var authService *AuthService
var once sync.Once

// GetAuthService 获取认证服务单例
func GetAuthService() *AuthService {
	once.Do(func() {
		authService = &AuthService{
			users:         make([]model.User, 0),
			tokenBlacklist: make(map[string]time.Time),
		}
		// 初始化测试用户数据
		authService.initTestUsers()
	})
	return authService
}

// 初始化测试用户数据
func (s *AuthService) initTestUsers() {
	// 创建默认管理员用户
	hashedPassword, _ := utils.HashPassword("admin123")
	adminUser := model.User{
		ID:        1,
		Name:      "张三",
		Email:     "admin@qq.com",
		Password:  hashedPassword,
		Age:       18,
		Phone:     132000000000,
		Address:   "杭州",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 创建普通测试用户
	hashedPassword2, _ := utils.HashPassword("user123")
	testUser := model.User{
		ID:        2,
		Name:      "李四",
		Email:     "user@qq.com",
		Password:  hashedPassword2,
		Age:       25,
		Phone:     138000000000,
		Address:   "北京",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.users = append(s.users, adminUser, testUser)
}

// Login 用户登录
func (s *AuthService) Login(loginReq *model.LoginRequest) (*model.LoginResponse, error) {
	if loginReq.Username == "" || loginReq.Password == "" {
		return nil, errors.New("用户名和密码不能为空")
	}

	// 查找用户
	user, err := s.findUserByEmailOrName(loginReq.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	if !utils.VerifyPassword(user.Password, loginReq.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成JWT Token
	token, err := utils.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return nil, errors.New("登录失败，请稍后重试")
	}

	// 返回登录响应
	return &model.LoginResponse{
		Token: token,
		User:  user,
	}, nil
}

// Logout 用户注销
func (s *AuthService) Logout(token string) error {
	if token == "" {
		return errors.New("Token不能为空")
	}

	// 验证Token
	if !utils.ValidateToken(token) {
		return errors.New("无效的Token")
	}

	// 将Token加入黑名单
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	s.tokenBlacklist[token] = time.Now().Add(time.Hour * 24) // 24小时后自动清理
	
	// 清理过期的黑名单Token
	s.cleanExpiredTokens()

	return nil
}

// IsTokenBlacklisted 检查Token是否在黑名单中
func (s *AuthService) IsTokenBlacklisted(token string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	expireTime, exists := s.tokenBlacklist[token]
	if !exists {
		return false
	}
	
	// 如果已过期，从黑名单中移除
	if time.Now().After(expireTime) {
		delete(s.tokenBlacklist, token)
		return false
	}
	
	return true
}

// GetUserByID 根据ID获取用户
func (s *AuthService) GetUserByID(userID int) (*model.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	for _, user := range s.users {
		if user.ID == userID {
			return &user, nil
		}
	}
	
	return nil, errors.New("用户不存在")
}

// Register 用户注册
func (s *AuthService) Register(regReq *model.RegisterRequest) (*model.User, error) {
	// 检查邮箱是否已存在
	if _, err := s.findUserByEmailOrName(regReq.Email); err == nil {
		return nil, errors.New("邮箱已被注册")
	}

	// 检查密码强度
	if !utils.IsPasswordValid(regReq.Password) {
		return nil, errors.New("密码强度不足：密码应包含大小写字母、数字，且长度至少8位")
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(regReq.Password)
	if err != nil {
		return nil, errors.New("注册失败，请稍后重试")
	}

	// 创建新用户
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	newUser := model.User{
		ID:        s.getNextUserID(),
		Name:      regReq.Name,
		Email:     regReq.Email,
		Password:  hashedPassword,
		Age:       regReq.Age,
		Phone:     regReq.Phone,
		Address:   regReq.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.users = append(s.users, newUser)
	return &newUser, nil
}

// ChangePassword 修改密码
func (s *AuthService) ChangePassword(userID int, changeReq *model.ChangePasswordRequest) error {
	// 查找用户
	user, err := s.GetUserByID(userID)
	if err != nil {
		return err
	}

	// 验证旧密码
	if !utils.VerifyPassword(user.Password, changeReq.OldPassword) {
		return errors.New("原密码错误")
	}

	// 检查新密码强度
	if !utils.IsPasswordValid(changeReq.NewPassword) {
		return errors.New("新密码强度不足：密码应包含大小写字母、数字，且长度至少8位")
	}

	// 加密新密码
	hashedPassword, err := utils.HashPassword(changeReq.NewPassword)
	if err != nil {
		return errors.New("修改密码失败，请稍后重试")
	}

	// 更新密码
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	for i := range s.users {
		if s.users[i].ID == userID {
			s.users[i].Password = hashedPassword
			s.users[i].UpdatedAt = time.Now()
			break
		}
	}

	return nil
}

// findUserByEmailOrName 根据邮箱或用户名查找用户
func (s *AuthService) findUserByEmailOrName(emailOrName string) (*model.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	emailOrName = strings.TrimSpace(strings.ToLower(emailOrName))
	
	for _, user := range s.users {
		if strings.ToLower(user.Email) == emailOrName || strings.ToLower(user.Name) == emailOrName {
			return &user, nil
		}
	}
	
	return nil, errors.New("用户不存在")
}

// getNextUserID 获取下一个用户ID
func (s *AuthService) getNextUserID() int {
	maxID := 0
	for _, user := range s.users {
		if user.ID > maxID {
			maxID = user.ID
		}
	}
	return maxID + 1
}

// cleanExpiredTokens 清理过期的黑名单Token
func (s *AuthService) cleanExpiredTokens() {
	now := time.Now()
	for token, expireTime := range s.tokenBlacklist {
		if now.After(expireTime) {
			delete(s.tokenBlacklist, token)
		}
	}
}

// RefreshToken 刷新Token
func (s *AuthService) RefreshToken(token string) (string, error) {
	// 解析当前token
	claims, err := utils.ParseToken(token)
	if err != nil {
		return "", errors.New("无效的Token")
	}

	// 检查token是否在黑名单中
	if s.IsTokenBlacklisted(token) {
		return "", errors.New("Token已失效")
	}

	// 验证用户是否存在
	user, err := s.GetUserByID(claims.UserID)
	if err != nil {
		return "", errors.New("用户不存在")
	}

	// 生成新的Token
	newToken, err := utils.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return "", errors.New("刷新Token失败")
	}

	// 将旧token加入黑名单
	s.mutex.Lock()
	s.tokenBlacklist[token] = time.Now().Add(time.Hour * 24)
	s.mutex.Unlock()

	return newToken, nil
}

// GetAllUsers 获取所有用户（管理员功能）
func (s *AuthService) GetAllUsers() []model.User {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	// 返回用户副本，不暴露密码
	users := make([]model.User, len(s.users))
	copy(users, s.users)
	
	for i := range users {
		users[i].Password = "" // 清空密码字段
	}
	
	return users
}