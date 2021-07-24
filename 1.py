def find_answer(input):
  sum_map = {}
  for x in input:
    next_sum = 2020 - x
    if next_sum in sum_map:
      return next_sum * x
    else:
      sum_map[x] = True
  return "You are messing with me!" 