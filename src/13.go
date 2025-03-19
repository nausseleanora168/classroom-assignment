func randomNumber(min: Int, max: Int) -> Int {
  let range = max - min + 1
  return min + arc4random_uniform(UInt32(range))
}