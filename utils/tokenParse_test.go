package utils

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestParseToken(t *testing.T) {
	//token, _ := GenerateToken("6659348")
	//spew.Dump(token)

	parse, _ := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE1OTI0MTcxMDQ2MjE5OTgwODAiLCJ1c2VyTmFtZSI6ImFkbWluIiwiZW1haWwiOiIiLCJwaG9uZSI6IjE4MzgwNDI3NDk3Iiwid3hJRCI6IjIwMjMtMDMtMDggMDk6MzE6NTguMjk1OTQ4OSArMDgwMCBDU1QgbT0rNDgyNzUuMDg1MjMwNzAxIn0.FSV3o9ytBocnGCo-2Cr0PQlPPNTGoTBfusz9kAhER_U")

	spew.Dump(parse)

	//	spew.Dump(IsTokenValid(token))
	//	spew.Dump(IsTokenValid(token + "sdlsdljkdlk"))
	//}
}
