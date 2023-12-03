file = open("inputs/inputs.txt", "r")
count = 0
max_red_count = 12
max_blue_count = 14
max_green_count = 13
while True:
    red_count = 0
    blue_count = 0
    green_count = 0
    content=file.readline()
    if not content:
        break

    game_id = content.split(":")[0].split(" ")[1]


    game_set =  content.split(":")[1]
    for game in game_set.split(";"):
        results = game.split(",")
        add = False
        for cube in results:
            if 'red' in cube:
                add = int(cube.split(" ")[1]) <= max_red_count
            elif 'blue' in cube:
                add = int(cube.split(" ")[1]) <= max_blue_count
            elif 'green' in cube:
                add = int(cube.split(" ")[1]) <= max_green_count
            if not add:
                # print(f'{game_id=} {cube=} {add=}')
                break
        if not add:
            break        
    if add :
        print(f'{game_id=} {game_set=}')
        count += int(game_id)

file.close()
print(count)