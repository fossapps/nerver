package main

import (
	"github.com/cyberhck/nerver"
	"fmt"
)

var xmlStr = `<Bug>
	<ScreenshotsComment>
		- [ ] All screenshots have visible address bar on top
		<!-- If irrevalant, just ignore this checklist -->
	</ScreenshotsComment>
	<UserInfo id="null" username="cyberhck" />
	<Environment>
		- [ x ] Live
		- [ ] Dev
		- [ ] Staging
	</Environment>
	<Browsers>
		- [ x ] N/A
		- [ ] Chrome
		- [ ] Firefox
		- [ ] Safari
		- [ ] Opera
		- [ ] IE
		- [ ] Edge
		- [ ] All
	</Browsers>
	<Description>
		I was trying to use configurator review tool, then I came across a weird bug
	</Description>
	<ReproductionSteps>
		- Went to http://backend.crazy-factory.com/backend/
		- Hard Reloaded (Ctrl + Shift + R)
		- Clicked on Login
		- Logged In
		- went to administration and clicked on configurator review tool
		- came across a design which was weird
	</ReproductionSteps>
	<ExpectedBehavior>
		I expected that thumbnail looks good and doesn't have any weird artifacts on it.
		<!-- Briefly explain what you expected from above -->
	</ExpectedBehavior>
	<ActualBehavior>
		But thumbnail had some weird artifacts, sort of like patterns.
		<!-- Briefly explain what happened and try provide lot of info -->
	</ActualBehavior>
	<AdditionalInfo>
		[config.json](https://s3.eu-central-1.amazonaws.com/cf-shop-production-designs/designs/product_designs/JT/JTTDX33IBAGWQECE/config.json), [thumbnail](https://s3.eu-central-1.amazonaws.com/cf-shop-production-designs/designs/product_designs/JT/JTTDX33IBAGWQECE/thumbnail.png), [Original Uploaded image](https://s3.eu-central-1.amazonaws.com/cf-shop-production-designs/designs/product_design_files/2d/2dec5b738f1449478d8932a2766f9b82.png), [production.pdf](file:///home/cyberkiller/Downloads/production.pdf) (looks okay) order ID: 8681192

I've a wild guess that maybe it has something to do with filters, but it shouldn't.
	</AdditionalInfo>
</Bug>`

var schema = `{
    "name": "Bug",
    "children": [
        {"element": {"name": "ScreenshotsComment","children": []}, "required": true},
        {"element": {
                "name": "UserInfo",
                "children": [],
                "props": [
                    {"name": "id", "required": true},
                    {"name": "username", "required": true}
                ]
            }, "required": true},
        {"element": {"name": "Environment","children": []}, "required": true},
        {"element": {"name": "Browsers", "children": []}, "required": true},
        {"element": {"name": "Description", "children": []}, "required": true},
        {"element": {"name": "ReproductionSteps", "children": []}, "required": true},
        {"element": {"name": "ExpectedBehavior", "children": []}, "required": true},
        {"element": {"name": "ActualBehavior", "children": []}, "required": true},
        {"element": {"name": "ActualBehavior", "children": []}, "required": true},
        {"element": {"name": "AdditionalInfo", "children": []}, "required": true}
    ]
}`


func main() {
	v := nerver.Validation{
		Xml: xmlStr,
		Schema: schema,
	}
	ok, err := v.Ok()
	fmt.Println(ok, err, v.Tokens)
}
