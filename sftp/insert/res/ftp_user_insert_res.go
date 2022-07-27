package res

type FtpUserInsertRes struct {
	ID                   int              `json:"id"`
	Status               int              `json:"status"`
	Username             string           `json:"username"`
	ExpirationDate       int              `json:"expiration_date"`
	HomeDir              string           `json:"home_dir"`
	UID                  int              `json:"uid"`
	Gid                  int              `json:"gid"`
	MaxSessions          int              `json:"max_sessions"`
	QuotaSize            int              `json:"quota_size"`
	QuotaFiles           int              `json:"quota_files"`
	Permissions          interface{}      `json:"permissions"`
	UploadDataTransfer   int              `json:"upload_data_transfer"`
	DownloadDataTransfer int              `json:"download_data_transfer"`
	TotalDataTransfer    int              `json:"total_data_transfer"`
	CreatedAt            int64            `json:"created_at"`
	UpdatedAt            int64            `json:"updated_at"`
	Filters              Filters          `json:"filters"`
	VirtualFolders       []VirtualFolders `json:"virtual_folders"`
	Filesystem           Filesystem       `json:"filesystem"`
}

type Hooks struct {
	ExternalAuthDisabled  bool `json:"external_auth_disabled"`
	PreLoginDisabled      bool `json:"pre_login_disabled"`
	CheckPasswordDisabled bool `json:"check_password_disabled"`
}

type TotpConfig struct {
	Secret interface{} `json:"secret"`
}
type Filters struct {
	AllowedIP  []string   `json:"allowed_ip"`
	Hooks      Hooks      `json:"hooks"`
	TotpConfig TotpConfig `json:"totp_config"`
}

type S3Config struct {
	AccessSecret interface{} `json:"access_secret"`
}

type Gcsconfig struct {
	Credentials interface{} `json:"credentials"`
}

type Azblobconfig struct {
	AccountKey interface{} `json:"account_key"`
	SasURL     interface{} `json:"sas_url"`
}

type Cryptconfig struct {
	Passphrase interface{} `json:"passphrase"`
}

type Sftpconfig struct {
	Password      interface{} `json:"password"`
	PrivateKey    interface{} `json:"private_key"`
	KeyPassphrase interface{} `json:"key_passphrase"`
}
type Filesystem struct {
	Provider     int          `json:"provider"`
	S3Config     S3Config     `json:"s3config"`
	Gcsconfig    Gcsconfig    `json:"gcsconfig"`
	Azblobconfig Azblobconfig `json:"azblobconfig"`
	Cryptconfig  Cryptconfig  `json:"cryptconfig"`
	Sftpconfig   Sftpconfig   `json:"sftpconfig"`
}
type VirtualFolders struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	MappedPath      string     `json:"mapped_path"`
	UsedQuotaSize   int        `json:"used_quota_size"`
	UsedQuotaFiles  int        `json:"used_quota_files"`
	LastQuotaUpdate int        `json:"last_quota_update"`
	Filesystem      Filesystem `json:"filesystem"`
	VirtualPath     string     `json:"virtual_path"`
	QuotaSize       int        `json:"quota_size"`
	QuotaFiles      int        `json:"quota_files"`
}
