file = open("inputs/inputs.txt", "r")
count = 0
convert_dict = {
    "one" : '1',
    "two" : '2',
    "three" : '3',
    "four" : '4',
    "five" : '5',
    "six" : '6',
    "seven" : '7',
    "eight" : '8',
    "nine" : '9'
}
while True:
    content=file.readline()
    line_number = ""
    # list_line_number = [s for s in content if s.isdigit()]
    lst_number = []
    for i in range(len(content)):
        for word in convert_dict:
            if content[i:i+len(word)] == word:
                lst_number.append(convert_dict[word])
        if content[i].isdigit():
            lst_number.append(content[i])
    print(f'{lst_number=},{content=}')
    if len(lst_number) == 1 :
        line_number += lst_number[0] + lst_number[0]
    if len(lst_number) > 1:
        line_number = lst_number[0] + lst_number[-1]
    if line_number != "":
        count += int(line_number)
    if not content:
        break
file.close()
print(count)