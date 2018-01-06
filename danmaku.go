package bilibili

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Possible font size value of a danmaku.
const (
	ExtremeSmall FontSize = 12
	VerySmall             = 16
	Small                 = 18
	Middle                = 25
	Large                 = 36
	VeryLarge             = 45
	ExtreneLarge          = 64
)

// Possible danmaku pool type.
const (
	General DanmakuPoolType = iota
	Subtitle
	Special
)

// API configuration for decrypting senderid of danmaku to global user id.
const (
	descriptionAPI = "http://biliquery.typcn.com/api/user/hash/"
)

// Danmaku represents a single danmaku sent by user.
type Danmaku struct {

	// Flash time represents when exactly should this danmaku appear during the lifetime
	// of the video.
	FlashTime float64

	// Mode represents different mode of danmaku.
	// 1 ~ 3 is the basic danmaku, which just pass through your screen from right to left.
	// 4 is the bottom danmaku.
	// 5 is the top danmaku.
	// 6 is the reverse direction danmaku, which goes from left to right.
	// 7 is the located danmaku, which has an absolute postion where it should appear.
	// 8 is the advanced danmaku
	Mode int

	// FontSize represents how large is the text of the danmaku, 25 as its default value.
	// The size could only be 12, 16, 18, 25, 36, 45 or 64
	FontSize FontSize

	// Color represents the color of the danmaku, it is decimal not RGB.
	Color int

	// Timestamp represents when this danmaku is sent.
	Timestamp int

	// Pool represents the purpose of danmaku.
	// 0 is for general danmaku.
	// 1 is for subtitle.
	// 3 is for others.

	// TODO: This field should be check before burst into alpha version. See Issue #2.
	Pool int

	// SenderID represents the sender, it is different from the global uid of the user.
	SenderID string

	// DatabaseID represents the id of the danmaku in bilibili database.
	DatabaseID int

	// Content of the danmaku.
	Content string
}

// FontSize represents the font size of the danmaku.
type FontSize int

// DanmakuPoolType represents three pool type of danmaku.
type DanmakuPoolType int

type decryption struct {
	Error int
	Data  []map[string]int
}

// DecryptUserID currently use a third party api to decrypt the user id.
// The encryption algorithm is I.363.5.
// For more algorithm details, see http://www.itu.int/rec/T-REC-I.363.5-199608-I/en
// For more API details, see http://blog.eqoe.cn/posts/bilibili-comment-sender-digger.html
func (d *Danmaku) DecryptUserID() (int, error) {

	resp, err := http.Get(descriptionAPI + string(d.SenderID))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result decryption
	if json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	if result.Error != 0 {
		return 0, errors.New("sender id not valid")
	}

	return result.Data[0]["id"], nil
}
