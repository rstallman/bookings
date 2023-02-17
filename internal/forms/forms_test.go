package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForms_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()

	if !isValid {
		t.Error("got invalid when should have been valid")
	}

}

func TestForms_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("forms shows valid when required field missing")
	}

	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "b")
	postData.Add("c", "c")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postData
	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}

}

func TestForms_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	if form.Has("a") {
		t.Error("forms shows has field a when no field provided")
	}

	postData := url.Values{}
	postData.Add("a", "a")

	form = New(postData)

	if !form.Has("a") {
		t.Error("forms shows no field a when it has")
	}

}

func TestForms_Minlength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}


	postedValues := url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)
	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("form shows min length  of 100 met when data is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")
	form = New(postedValues)
	form.MinLength("another_field", 1)

	if !form.Valid() {
		t.Error("form shows min length  of 1 is not met when it is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error, but  got one")
	}

}

func TestForms_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")

	if form.Valid() {
		t.Error("forms shows valid email for non-existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "me@email.com")
	form = New(postedValues)
	form.IsEmail("email")

	if !form.Valid() {
		t.Error("forms shows invalid email for valid email")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "x")
	form = New(postedValues)
	form.IsEmail("email")

	if form.Valid() {
		t.Error("forms shows valid email for invalid email")
	}
}
