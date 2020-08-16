package main

import (
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
    // 暗号化したい文字列
    plaintext := []byte("My name is Astaxie")
    // 暗号化された文字列を渡すと、plaintextは渡された文字列になります。
    if len(os.Args) > 1 {
        plaintext = []byte(os.Args[1])
    }

    // aesの暗号化文字列
    keyText := "astaxie12798akljzmknm.ahkjkljl;k"
    if len(os.Args) > 2 {
        keyText = os.Args[2]
    }

    fmt.Println(len(keyText))

    // 暗号化アルゴリズムaesを作成
    c, err := aes.NewCipher([]byte(keyText))
    if err != nil {
        fmt.Printf("Error: NewCipher(%d bytes) = %s", len(keyText), err)
        os.Exit(-1)
    }

    // 暗号化文字列
    cfb := cipher.NewCFBEncrypter(c, commonIV)
    ciphertext := make([]byte, len(plaintext))
    cfb.XORKeyStream(ciphertext, plaintext)
    fmt.Printf("%s=>%x\n", plaintext, ciphertext)

    // 復号文字列
    cfbdec := cipher.NewCFBDecrypter(c, commonIV)
    plaintextCopy := make([]byte, len(plaintext))
    cfbdec.XORKeyStream(plaintextCopy, ciphertext)
    fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)
}