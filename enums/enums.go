package enums

// PairType 键值对类型定义
type PairType string

// 用于标识 form 表单参数类型
const (
	PairTypeText PairType = "text"
	PairTypeFile PairType = "file"
)

// EntityType 标识实体类型
type EntityType string

const (
	Form    EntityType = "multipart/form-data"
	RawJson EntityType = "application/json"
)
