# Form Validations

- [# Basics](#-basics)
  - [# Generate a Validation](#-generate-a-validation)
- [# How to Use](#-how-to-use)
- [# Lists of rules](#-lists-of-rules)

---

{#-basics}

## [#](#-basics) Basics

The validation will help you to filter out invalid inputs or to limit someone from passing too much data.

This page will help you to write a simple validation and re-use it wherever in your lucid project.

{#-generate-a-validation}

### [#](#-generate-a-validation) Generate a Validation

Here's how to generate a validation file

```bash
./run make:validation sign_up

Created validation, located at:
 > ~/lucid-setup/src/lucid/app/validations/sign_up.go
```

After generating, the file should look like this

```go
// app/validations/sign_up.go
type SignUpValidator struct {
    Rules *must.SetOfRules
}

func SignUp() *SignUpValidator {
    return &SignupValidator{
        Rules: &must.SetOfRules{
            "email": {
                &must.Required{},
                &must.Email{},
            },
            "name": {
                &must.Required{},
                &must.Min{Value:1},
                &must.Max{Value:100},
            },
            "password": {
                &must.Required{},
                &must.Min{Value: 8},
                &must.StrictPassword{
                    WithSpecialChar: true,
                    WithUpperCase:   true,
                    WithLowerCase:   true,
                    WithDigit:       true,
                },
            },
            "confirmed_password": {
                &must.Required{},
                &must.Matches{TargetField: "password"},
            },
        },
    }
}

func (v SignupValidator) Create() *must.SetOfRules {
    return v.Rules
}
```

{#-how-to-use}

## [#](#-how-to-use) How to Use

Most of the time you're going to use this in your handler, an example can be seen below

```go
// app/handlers/sign_up_handler.go
1. func SignUpHandler(T engines.EngineContract) *errors.AppError {
2.    // ...
3.
4.    if validator := request.Validator(validations.SignUp().Create()); validator != nil {
5.        return response.Json(map[string]interface{}{
6.            "fails": validator.ValidationError,
7.        }, http.StatusUnauthorized)
8.    }
9. }
```

At `line 4` we're calling this `validations.SignUp().Create()`, what it does is to initialize the set of rules and then we're calling the `Create()` to return all the rules to us.

Then passing that rules to our [`request.Validator`](http://localhost:8332/handlers#-request--response)

At `line 5` if the vlaidation is not `nil`, then return [`json response`](http://localhost:8332/handlers#-request--response) with the validation error in it.

The response will look like this

```json
{
  fails: {
    email: 'johndoe@gmail is not a valid email address!',
    password: 'password should contain at least 1 special character!'
  }
}
```

{#-lists-of-rules}

## [#](#-lists-of-rules) Lists of rules

Name | Sample Error Messages
-----|-----
Email | johndoe@gmail is not a valid email address!
Max | name is set to maximum of 1 length
Min | name is set to minimum of 100 length
Required | email is required!
StrictPassword | password should contain at least 1 special character!<br>password should contain at least 1 upper case character!<br>password should contain at least 1 lower case character!<br>password should contain at least 1 digit!
Matches | confirm_password did not match with password
