file = open("inputs/inputs.txt", "r")
count = 0
while True:
    red_count = 0
    blue_count = 0
    green_count = 0
    game_count = 0
    content=file.readline()
    if not content:
        break

    game_id = content.split(":")[0].split(" ")[1]


    game_set =  content.split(":")[1]
    for game in game_set.split(";"):
        results = game.split(",")
        for cube in results:
            if 'red' in cube:
                red_count = int(cube.split(" ")[1]) if (int(cube.split(" ")[1]) > red_count) else red_count
                print(f'{red_count=} {cube}')
            elif 'blue' in cube:
                blue_count = int(cube.split(" ")[1]) if (int(cube.split(" ")[1]) > blue_count) else blue_count
                print(f'{blue_count=} {cube}')
            elif 'green' in cube:
                green_count = int(cube.split(" ")[1]) if (int(cube.split(" ")[1]) > green_count) else green_count
                print(f'{green_count=} {cube}')
    game_count = red_count * blue_count * green_count
    count += game_count
    print(f'{game_id=} {game_count=} {count=} \n')
file.close()
print(count)