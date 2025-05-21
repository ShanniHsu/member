package message

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// init不需手動呼叫(與本專案的 mainInit不同)
// 特殊命名 init() -> 無參數、無回傳
// 自動執行 Go編譯器在執行main()之前會自動執行所有的init()
// 執行順序 由import tree 的依賴順序決定 https://blog.csdn.net/joeyoj/article/details/135745098
var bundle *i18n.Bundle

func init() {
	// 中央翻譯庫，用來儲存所有語系的翻譯檔
	// 在這裏設定預設語言
	bundle = i18n.NewBundle(language.English)
	// 註冊翻譯檔解析器
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	// 系統支援的語言
	supportLangs := []string{"en", "zh"}

	for _, lang := range supportLangs {
		path := fmt.Sprintf("./pkg/message/locales/%s.json", lang)
		// 必要載入翻譯檔，如果有錯誤直接讓系統崩潰
		bundle.MustLoadMessageFile(path)
	}
	fmt.Println("Init i18n successfully")
}

func GetMsg(msgID string, lang string) (message string) {
	localizer := i18n.NewLocalizer(bundle, lang)
	message, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: msgID,
	})
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("message: ", message)
	return
}
