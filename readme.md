## <center> Un-Official Ovoid API Wrapper package for Go
[![Documentation](https://godoc.org/github.com/adibaulia/ovoid-go?status.svg)](http://godoc.org/github.com/adibaulia/ovoid-go) [![Test-master Actions Status](https://github.com/adibaulia/ovoid-go/workflows/Test/badge.svg)](https://github.com/adibaulia/ovoid-go/actions)
</center>

Repository berikut ini merupakan porting dari [ovoid](https://github.com/lintangtimur/ovoid) untuk Go 

### Install
```
go get github.com/adibaulia/ovoid-go
```

### Usage
Anda harus mendapatkan Token untuk menggunakan API Ovo.
<br>
<br>
<b>Langkah 1</b>
```
package main

import (
  ovo "github.com/adibaulia/ovoid-go"
)

func main(){
  login, err := ovo.NewOvoLogin("your_phone")
  if err != nil {
    ...
  }

  l, err := login.Login2FA()
  if err != nil {
    ...
  }

```

<b>Langkah 2</b>

```
  accessToken, err := login.Login2FAVerify(l.RefID, "your_verification_code")
  if err != nil {
    ...
  }
```
<b>Langkah 3</b>
```
  auth, err := login.LoginSecurityCode(accessToken.UpdateAccessToken)
  if err != nil {
      ...
  }
  YOUR_TOKEN := auth.Token

```

Lalu gunakan YOUR_TOKEN untuk menggunakan package ovoid. Contoh: 
```
package main

import (
  ovo "github.com/adibaulia/ovoid-go"
)

func main(){
  o, err := ovo.NewClient("YOUR_TOKEN")
  if err != nil {
    ...
  }

  b, err := o.GetAllBalance()
  if err != nil {
    ...
  }
}
```
### Progress

Masih dalam pengembangan. Silahkan Pull Request untuk berkontribusi. 

Thank You !
