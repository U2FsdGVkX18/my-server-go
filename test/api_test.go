package test

import (
	"my-server-go/api"
	"testing"
)

func TestOne(t *testing.T) {
	var msg_signature = "bb65bc2127f86d862df8e6917daa7ef4a7b1733d"
	var timestamp = "1677596202"
	var nonce = "1676975221"
	var data = "<xml><ToUserName><![CDATA[ww8d5186f5aa839ee7]]></ToUserName><Encrypt><![CDATA[MMs94DR1S7fGleh0mKyjA3RNgwuNNGu0YyGimD+99GB16gCpuhBkvWxaTt20L1PC6Ni0VBWnpSdlpUsWseUbFpmsRtt8aFkTdoyRBe8C0gx9hM8bLrOWGcdOJrtXaGIUnOF8H8UuinQXLjO/uAulBKLKE7TiMFXtvaQ62/Iuzc5UKdh8bAbGUk+iOY1nUkh3L5BSPpyWHWVKEFyLkumjUCWZV4L11lSuG9nqbDVVFhdHLT/Du3TCX/To4DW7DIUyjgpzARVjAPzBzGvYYe1Nq1Y3RkjbwdWRWz824xhgmYEpiUr4XOYlfqTWljydXOV+NdNmJBXc/WDnG4u2jo1HsUjYRsWzaYqux4CX3dm1WI6L9iDJB87F5Ldp90yoxuf5rBdb3xLtssxBu8S4zievwlZVzRnQWN33Xvg0fUKHjo0=]]></Encrypt><AgentID><![CDATA[1000002]]></AgentID></xml>"
	api.ProcessMessage(msg_signature, timestamp, nonce, []byte(data))

}
