file = open("inputs/inputs.txt", "r")
count = 0
while True:
    content=file.readline()
    line_number = ""
    list_line_number = [s for s in content if s.isdigit()]
    if len(list_line_number) == 1 :
        line_number += list_line_number[0] + list_line_number[0]
    if len(list_line_number) > 1:
        line_number = list_line_number[0] + list_line_number[-1]
    if line_number != "":
        count += int(line_number)
    if not content:
        break
file.close()
print(count)