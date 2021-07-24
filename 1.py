def find_answer(input):
  for x in input:
    next_sum = 2020 - x
    if next_sum in input:
      return next_sum * x
  return "You are messing with me!" 
