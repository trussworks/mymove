// Code generated by go-swagger; DO NOT EDIT.

package payment_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// UpdatePaymentServiceItemStatusURL generates an URL for the update payment service item status operation
type UpdatePaymentServiceItemStatusURL struct {
	MoveTaskOrderID      string
	PaymentServiceItemID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UpdatePaymentServiceItemStatusURL) WithBasePath(bp string) *UpdatePaymentServiceItemStatusURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *UpdatePaymentServiceItemStatusURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *UpdatePaymentServiceItemStatusURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/move-task-orders/{moveTaskOrderID}/payment-service-items/{paymentServiceItemID}/status"

	moveTaskOrderID := o.MoveTaskOrderID
	if moveTaskOrderID != "" {
		_path = strings.Replace(_path, "{moveTaskOrderID}", moveTaskOrderID, -1)
	} else {
		return nil, errors.New("moveTaskOrderId is required on UpdatePaymentServiceItemStatusURL")
	}

	paymentServiceItemID := o.PaymentServiceItemID
	if paymentServiceItemID != "" {
		_path = strings.Replace(_path, "{paymentServiceItemID}", paymentServiceItemID, -1)
	} else {
		return nil, errors.New("paymentServiceItemId is required on UpdatePaymentServiceItemStatusURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/ghc/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *UpdatePaymentServiceItemStatusURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *UpdatePaymentServiceItemStatusURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *UpdatePaymentServiceItemStatusURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on UpdatePaymentServiceItemStatusURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on UpdatePaymentServiceItemStatusURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *UpdatePaymentServiceItemStatusURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}