package ginx

type ListResponses struct {
	List interface{} `json:"list"`
	Page interface{} `json:"page,omitempty"`
}

type InfoResponses struct {
	Info interface{} `json:"info"`
}

type LoginResponses struct {
	User  interface{} `json:"user"`
	Token interface{} `json:"token"`
}

type RoleResponses struct {
	Info interface{} `json:"info"`
	Role interface{} `json:"role"`
}
