package util

import(
      "crypto/hmac"
      "crypto/sha256"
      "encoding/base64"
      "time"
)

func GenerateSign(message, key string) string{
     mac := hmac.New(sha256.New, []byte(key))
     mac.Write([]byte(message))
     return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func CompareSign(srcSign, message, key string) bool {
     return GenerateSign(message, key) == srcSign
}

func CompareTimestamp( srcTimestamp int32) bool {
     if int32(time.Now().Unix()) - srcTimestamp < 3600 {
        return true
     } else{
        return false
     }
} 
