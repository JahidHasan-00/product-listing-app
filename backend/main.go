package main

import (
	"ecommerce/cmd"
)

// "ecommerce/cmd"

func main() {

	cmd.Serve()

	// var s string = "Hello World"

	// byteArr := []byte(s)

	// fmt.Println(s)
	// fmt.Println(byteArr)

	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)
	// b64Str := enc.EncodeToString(byteArr)

	// fmt.Println(b64Str)

	// decodeStr, err := enc.DecodeString(b64Str)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(decodeStr)

	// data := []byte("Hello World")

	// fmt.Println(string(data))
	// fmt.Println(data)

	// sha := sha256.Sum256(data)
	// fmt.Println(sha)

	// secrete := []byte("Secrete-Key")
	// message := []byte("Hello World")

	// h := hmac.New(sha256.New, secrete)
	// // fmt.Println(h)

	// text := h.Sum(message)
	// fmt.Println(text)

	// jwt, err := util.CreateJwt("my-secrete", util.Payload{
	// 	Sub:         45,
	// 	FirstName:   "Habibur",
	// 	LastName:    "Rahaman",
	// 	Email:       "habibur@gmail.com",
	// 	IsShopOwner: false,
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(jwt)

}
