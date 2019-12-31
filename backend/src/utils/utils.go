package utils

func Assert(condition bool, onFalseFn func()) {
  if !condition {
    onFalseFn()
  }
}
