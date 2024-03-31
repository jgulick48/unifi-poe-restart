package unificlient

import "time"

type LoginResponse struct {
	UniqueId           string `json:"unique_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Alias              string `json:"alias"`
	FullName           string `json:"full_name"`
	Email              string `json:"email"`
	EmailStatus        string `json:"email_status"`
	EmailIsNull        bool   `json:"email_is_null"`
	Phone              string `json:"phone"`
	AvatarRelativePath string `json:"avatar_relative_path"`
	AvatarRpath2       string `json:"avatar_rpath2"`
	Status             string `json:"status"`
	EmployeeNumber     string `json:"employee_number"`
	CreateTime         int    `json:"create_time"`
	LoginTime          int    `json:"login_time"`
	Extras             struct {
	} `json:"extras"`
	Username          string `json:"username"`
	LocalAccountExist bool   `json:"local_account_exist"`
	PasswordRevision  int    `json:"password_revision"`
	SsoAccount        string `json:"sso_account"`
	SsoUuid           string `json:"sso_uuid"`
	SsoUsername       string `json:"sso_username"`
	SsoPicture        string `json:"sso_picture"`
	UidSsoId          string `json:"uid_sso_id"`
	UidSsoAccount     string `json:"uid_sso_account"`
	UidAccountStatus  string `json:"uid_account_status"`
	Groups            []struct {
		UniqueId   string        `json:"unique_id"`
		Name       string        `json:"name"`
		UpId       string        `json:"up_id"`
		UpIds      []interface{} `json:"up_ids"`
		SystemName string        `json:"system_name"`
		CreateTime time.Time     `json:"create_time"`
	} `json:"groups"`
	Roles []struct {
		UniqueId   string    `json:"unique_id"`
		Name       string    `json:"name"`
		SystemRole bool      `json:"system_role"`
		SystemKey  string    `json:"system_key"`
		Level      int       `json:"level"`
		CreateTime time.Time `json:"create_time"`
		UpdateTime time.Time `json:"update_time"`
		IsPrivate  bool      `json:"is_private"`
	} `json:"roles"`
	Permissions struct {
		AccessManagement         []string `json:"access.management"`
		CalculusManagement       []string `json:"calculus.management"`
		ConnectManagement        []string `json:"connect.management"`
		DriveManagement          []string `json:"drive.management"`
		InnerspaceManagement     []string `json:"innerspace.management"`
		LedManagement            []string `json:"led.management"`
		NetworkManagement        []string `json:"network.management"`
		OlympusManagement        []string `json:"olympus.management"`
		ProtectManagement        []string `json:"protect.management"`
		SystemManagementLocation []string `json:"system.management.location"`
		SystemManagementUser     []string `json:"system.management.user"`
		TalkManagement           []string `json:"talk.management"`
	} `json:"permissions"`
	Scopes             []string    `json:"scopes"`
	CloudAccessGranted bool        `json:"cloud_access_granted"`
	OnlyLocalAccount   bool        `json:"only_local_account"`
	UpdateTime         int         `json:"update_time"`
	Avatar             interface{} `json:"avatar"`
	NfcToken           string      `json:"nfc_token"`
	NfcDisplayId       string      `json:"nfc_display_id"`
	NfcCardType        string      `json:"nfc_card_type"`
	NfcCardStatus      string      `json:"nfc_card_status"`
	Role               string      `json:"role"`
	RoleId             string      `json:"roleId"`
	Id                 string      `json:"id"`
	IsOwner            bool        `json:"isOwner"`
	IsSuperAdmin       bool        `json:"isSuperAdmin"`
	IsMember           bool        `json:"isMember"`
	MaskedEmail        string      `json:"maskedEmail"`
	AccessMask         int         `json:"accessMask"`
	PermissionMask     int         `json:"permissionMask"`
	UcorePermission    struct {
		HasUpdateAndInstallPermission bool `json:"hasUpdateAndInstallPermission"`
	} `json:"ucorePermission"`
	DeviceToken string `json:"deviceToken"`
	SsoAuth     struct {
	} `json:"ssoAuth"`
}

type SwitchCommand struct {
	Mac     string `json:"mac"`
	PortIDX int    `json:"port_idx"`
	Cmd     string `json:"cmd"`
}
