package controller

import (
	"github.com/gin-gonic/gin"
	"member/router/app/content/create-order"
	"member/router/app/content/login"
	"member/router/app/content/register"
	"net/http"
)

func (c appController) Register(ctx *gin.Context) {
	req := new(register.Request)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = c.userService.Register(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register successfully!",
	})
	return
}

func (c appController) Login(ctx *gin.Context) {
	req := new(login.Request)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if req.Account == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "The account must input.",
		})
		return
	}

	if req.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "The password must input.",
		})
		return
	}
	jwtToken, err := c.userService.Login(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Login successfully!",
		"jwtToken": jwtToken,
	})
	return
}

func (c appController) Receive(ctx *gin.Context) {
	//req := new(receive.Request)
	//err := ctx.ShouldBindJSON(req)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//// 收集所有參數
	//params := make(map[string]string)
	//for key, values := range ctx.Request.PostForm {
	//	if len(values) > 0 {
	//		params[key] = values[0] // 取第一個值
	//	}
	//}
	//
	//fmt.Println("params: ", params)
	//
	//// 計算 CheckMacValue
	//received := params["CheckMacValue"]
	//
	//// 比對 CheckMacValue
	//if req.CheckMacValue != received {
	//	ctx.String(400, "CheckMacValue mismatch")
	//	return
	//}
	//
	//// 處理成功，回應綠界
	//ctx.String(200, "1|OK")

}

func (c appController) GetUserInfo(ctx *gin.Context) {
	resp, err := c.userService.GetUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Successfully",
		"data":    resp,
	})
	return
}

func (c appController) CreateOrder(ctx *gin.Context) {
	req := new(create_order.Request)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if req.MerchantID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "MerchantID can't be null!",
		})
		return
	}

	if req.MerchantTradeNo == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "MerchantTradeNo can't be null!",
		})
		return
	}

	if req.MerchantTradeDate == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "MerchantTradeDate can't be null!",
		})
		return
	}

	if req.PaymentType == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "PaymentType can't be null!",
		})
		return
	}

	if req.TotalAmount == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "TotalAmount can't be null!",
		})
		return
	}

	if req.TradeDesc == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "TradeDesc can't be null!",
		})
		return
	}

	if req.ItemName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ItemName can't be null!",
		})
		return
	}

	if req.ReturnURL == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ReturnURL can't be null!",
		})
		return
	}

	if req.EncryptType == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "EncryptType can't be null!",
		})
		return
	}

	body, err := c.userService.CreateOrder(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)

	return
}
