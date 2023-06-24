package main

import "net/http"

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "terminal", nil); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	cardHolder := r.Form.Get("cardholder-name")
	email := r.Form.Get("email")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_currency")

	data := make(map[string]interface{})
	data["cardHolder"] = cardHolder
	data["email"] = email
	data["pI"] = paymentIntent
	data["pM"] = paymentMethod
	data["pA"] = paymentAmount
	data["pC"] = paymentCurrency
	if err := app.renderTemplate(w, r, "succeeded", &TemplateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}
