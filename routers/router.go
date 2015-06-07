// Copyright 2015 iCloudFund. All Rights Reserved.

package routers

import (
	"github.com/astaxie/beego"
	"github.com/wangch/ifundmgr/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/signin", &controllers.MainController{}, "get:SigninGet;post:SigninPost")
	beego.Router("/signout", &controllers.MainController{}, "post:SignoutPost")

	beego.Router("/issue", &controllers.MainController{}, "get:IssuesGet;post:IssuesPost")
	beego.Router("/deposit", &controllers.MainController{}, "get:DepositsGet;post:DepositsPost")
	beego.Router("/redeem", &controllers.MainController{}, "get:RedeemsGet;post:RedeemsPost")
	beego.Router("/withdrawal", &controllers.MainController{}, "get:WithdrawalsGet;post:WithdrawalsPost")

	beego.Router("/issue/add", &controllers.MainController{}, "get:AddIssueGet;post:AddIssuePost")
	beego.Router("/deposit/add", &controllers.MainController{}, "get:AddDepositGet;post:AddDepositPost")

	beego.Router("/issue/verify?:id", &controllers.MainController{}, "post:VerifyIssue")
	beego.Router("/deposit/verify?:id", &controllers.MainController{}, "post:VerifyDeposit")
	beego.Router("/redeem/verify?:id", &controllers.MainController{}, "post:VerifyIssue")
	beego.Router("/withdrawal/verify?:id", &controllers.MainController{}, "post:VerifyDeposit")

	beego.Router("/api/deposit", &controllers.MainController{}, "post:DepositApi")
	beego.Router("/api/withdrawal", &controllers.MainController{}, "post:WithdrawalApi")
	beego.Router("/api/issue", &controllers.MainController{}, "post:IssueApi")
	beego.Router("/api/redeem", &controllers.MainController{}, "post:RedeemApi")
}
