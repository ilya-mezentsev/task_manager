package api

import (
  mock2 "mock/plugins"
  "plugins"
  "plugins/code"
  "plugins/db"
)

func init() {
  var coder = code.NewCoder("123456789012345678901234")
  testingHelper.Token, _ = coder.Encrypt(map[string]interface{}{
    "role": "admin",
  })

  InitGroupWorkerRequestHandler(plugins.NewDBProxy(testingHelper.Database))
  db.ExecQuery(testingHelper.Database, mock2.TurnOnForeignKeys)
}
