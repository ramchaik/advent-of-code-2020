def find_answer(input):
  ans = 0
  for i in input:
    [policy, password] = i.split(': ')
    [min_max_len, char] = policy.split() 
    [min_limit, max_limit] = min_max_len.split('-')
    if char not in password:
      continue
    current_char_count = password.count(char)
    if int(min_limit) <= current_char_count and int(max_limit) >= current_char_count:
      ans += 1
  return ans

def find_answer_2(input):
  ans = 0
  for i in input:
    [policy, password] = i.split(': ')
    [indxes, char] = policy.split() 
    [idx_1, idx_2] = indxes.split('-')

    if char not in password:
      continue

    if (password[int(idx_1) - 1] == char) or (password[int(idx_2) - 1] == char):
      if password[int(idx_1) - 1] == char and password[int(idx_2) - 1] == char:
        continue
      else: 
        ans += 1
  return ans
