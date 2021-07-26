def find_answer(input):
  # two sum
  sum_map = {}
  for x in input:
    next_sum = 2020 - x
    if next_sum in sum_map:
      return next_sum * x
    else:
      sum_map[x] = True
  return "You are messing with me!" 

def find_answer_2(input):
  # three sum, could avoid extra steps in loop but meh!
  for x in input:
    next_sum = 2020 - x
    sum_map = {}
    for y in input:
      next_next_sum = next_sum - y 
      if next_next_sum in sum_map:
        return next_next_sum * x * y
      else:
        sum_map[y] = True
  return "You are messing with me!" 