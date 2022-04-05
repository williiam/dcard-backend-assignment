package util

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/itchyny/base58-go"
	uuid "github.com/nu7hatch/gouuid"
)

// UNIT TEST : 

// GetExpireTime
// input 6h 
// output 現在時間 + 6h

// sha256Of
// 自己給正確答案測 


// base58Encoded
// 自己給正確答案測 


// GenerateShortLink
// input: 長url 
// ouput: 短url 