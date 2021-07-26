def find_answer(input):
  ans = 0
  for i in input:
    [policy, password] = i.split(':')
    [min_max_len, char] = policy.split() 
    [min_limit, max_limit] = min_max_len.split('-')
    if char not in password:
      continue
    current_char_count = password.count(char)
    if int(min_limit) <= current_char_count and int(max_limit) >= current_char_count:
      ans += 1
  return ans