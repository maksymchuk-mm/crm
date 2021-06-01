package errors

type CrmError interface {
	Code() int
	Error() string
	ToMap() map[string]interface{}
}
