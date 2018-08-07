package true_vendor_sdk

import (
	"net/http"
	"github.com/alexmay23/httpshared/shared"
	"github.com/alexmay23/httputils"
)

type VendorCallbackHandler struct {
	useCase VendorUseCase
}

func NewVendorCallbackHandler(useCase VendorUseCase)*VendorCallbackHandler{
	return &VendorCallbackHandler{useCase:useCase}
}



func(self *VendorCallbackHandler)writeMapOrError(w http.ResponseWriter, m map[string]interface{}, err error){
	if err != nil{
	 	shared.NewServerError(400, "undefined", err.Error(), "VENDOR_ERROR").Write(w)
	}else{
		httputils.JSON(w, m, 200)
	}
}


func(self *VendorCallbackHandler)ReserveVerificationHandler(w http.ResponseWriter, r *http.Request){
	body, err := httputils.GetValidatedBody(r, reserveVerificationValidatorMap())
	if err != nil{
		err.(httputils.ServerError).Write(w)
	}
	token, err := self.useCase.ReservationVerification(body["id"].(string), body["user_id"].(string))
	self.writeMapOrError(w, map[string]interface{}{"token":token}, err)
}

func(self *VendorCallbackHandler)ReserveCallbackHandler(w http.ResponseWriter, r *http.Request){
	body, err := httputils.GetValidatedBody(r, idValidatorMap())
	if err != nil{
		err.(httputils.ServerError).Write(w)
	}
	err = self.useCase.ReservedCallback(body["id"].(string))
	self.writeMapOrError(w, shared.OKJSON, err)
}

func(self *VendorCallbackHandler)ResolveCallbackHandler(w http.ResponseWriter, r *http.Request){
	body, err := httputils.GetValidatedBody(r, idValidatorMap())
	if err != nil{
		err.(httputils.ServerError).Write(w)
	}
	err = self.useCase.ResolvingCallback(body["id"].(string))
	self.writeMapOrError(w, shared.OKJSON, err)
}

func(self *VendorCallbackHandler)CancelReserveHandler(w http.ResponseWriter, r *http.Request){
	body, err := httputils.GetValidatedBody(r, idValidatorMap())
	if err != nil{
		err.(httputils.ServerError).Write(w)
	}
	err = self.useCase.CancelReservationCallback(body["id"].(string))
	self.writeMapOrError(w, shared.OKJSON, err)
}

func (self *VendorCallbackHandler)ResolveVerification(w http.ResponseWriter, r *http.Request){
	body, err := httputils.GetValidatedBody(r, resolveVerificationValidatorMap())
	if err != nil{
		err.(httputils.ServerError).Write(w)
	}
	token, err := self.useCase.ResolvingVerification(body["id"].(string), body["token"].(string))
	self.writeMapOrError(w, map[string]interface{}{"token":token}, err)
}