# Prompt Credentials Provider for MoneyLover API Client

[![GitHub Releases](https://img.shields.io/github/v/release/nhatthm/moneyloverprompt)](https://github.com/nhatthm/moneyloverprompt/releases/latest)
[![Build Status](https://github.com/nhatthm/moneyloverprompt/actions/workflows/test.yaml/badge.svg)](https://github.com/nhatthm/moneyloverprompt/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/nhatthm/moneyloverprompt/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/moneyloverprompt)
[![Go Report Card](https://goreportcard.com/badge/github.com/nhatthm/httpmock)](https://goreportcard.com/report/github.com/nhatthm/httpmock)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/github.com/nhatthm/moneyloverprompt)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

`moneyloverprompt` uses [AlecAivazis/survey](https://github.com/AlecAivazis/survey) to get credentials from the prompt.

## Prerequisites

- `Go >= 1.15`

## Install

```bash
go get github.com/nhatthm/moneyloverprompt
```

## Usage

**Examples**

```go
package mypackage

import (
	"github.com/nhatthm/moneyloverapi"
	"github.com/nhatthm/moneyloverprompt/credentials"
)

func buildClient() (*moneyloverapi.Client, error) {
	c := moneyloverapi.NewClient(
		credentials.WithCredentialsAtLast(),
	)

	return c, nil
}
```

## Donation

If this project help you reduce time to develop, you can give me a cup of coffee :)

### Paypal donation

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://user-images.githubusercontent.com/1154587/113494222-ad8cb200-94e6-11eb-9ef3-eb883ada222a.png" width="147px" />
