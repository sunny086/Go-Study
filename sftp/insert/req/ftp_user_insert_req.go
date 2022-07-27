package req

type FtpUserInsertReq struct {
	Status         int              `json:"status"`
	Username       string           `json:"username"`
	Password       string           `json:"password"`
	HomeDir        string           `json:"home_dir"`
	VirtualFolders []VirtualFolders `json:"virtual_folders"`
	Permissions    interface{}      `json:"permissions"`
	Filters        Filters          `json:"filters"`
}
type Filesystem struct {
	Provider int `json:"provider"`
}
type VirtualFolders struct {
	Name        string     `json:"name"`
	MappedPath  string     `json:"mapped_path"`
	Users       []string   `json:"users"`
	Filesystem  Filesystem `json:"filesystem"`
	VirtualPath string     `json:"virtual_path"`
}

type Filters struct {
	AllowedIP []string `json:"allowed_ip"`
}
