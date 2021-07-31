package config

type Server struct {
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
	AutoCode AutoCode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	Mysql    Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Timer    Timer    `mapstructure:"timer" json:"timer" yaml:"timer"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Local    Local    `mapstructure:"local" json:"local" yaml:"local"`
	Captcha  Captcha  `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	JWT      JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Email    Email    `mapstructure:"email" json:"email" yaml:"email"`
	Casbin   Casbin   `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
}
