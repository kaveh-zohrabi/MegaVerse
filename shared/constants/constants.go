package constants

const (
	// Service names
	ServiceAPIGateway    = "api-gateway"
	ServiceAuthService   = "auth-service"
	ServiceUserService   = "user-service"
	ServiceSocialService = "social-service"
	ServiceMessaging     = "messaging-service"
	ServiceAIService     = "ai-service"

	// Default ports
	PortAPIGateway    = 8080
	PortAuthService   = 8081
	PortUserService   = 8082
	PortSocialService = 8083
	PortMessaging     = 8084
	PortAIService     = 8085

	// Database
	DefaultDBHost     = "localhost"
	DefaultDBPort     = 3306
	DefaultDBName     = "megaverse"
	DefaultDBUser     = "megaverse"
	DefaultDBPassword = "password"

	// Redis
	DefaultRedisHost = "localhost"
	DefaultRedisPort = 6379

	// JWT
	DefaultJWTExpiration        = 86400      // 24 hours
	DefaultJWTRefreshExpiration = 604800     // 7 days

	// Rate limiting
	DefaultRateLimit     = 1000
	DefaultRateWindowSec = 60

	// Roles
	RoleAdmin = "admin"
	RoleUser  = "user"
	RoleMod   = "moderator"
)
