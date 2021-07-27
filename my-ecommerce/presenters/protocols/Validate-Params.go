package protocols

type ValidateParam interface {
	ValidateRequestParams(params interface{}) error
}
