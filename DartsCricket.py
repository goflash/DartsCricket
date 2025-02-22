def get_valid_throw(sector):
    while True:
        throw_value = int(input("Первый бросок: "))
        if throw_value in (sector, sector * 2, sector * 3):
            return throw_value
        else:
            print("______________________")
            print("Ошибка вводите заново: ")
            print("______________________")


sum1 = 0
sum2 = 0

for sector in range(1, 4):
    print(f"Сектор {sector}")
    a = get_valid_throw(sector)
    b = get_valid_throw(sector)
    c = get_valid_throw(sector)
    sum1 += a + b + c

print("Итоговый результат Игорька 1:", sum1)

for sector in range(1, 4):
    print(f"Сектор {sector}")
    e = get_valid_throw(sector)
    g = get_valid_throw(sector)
    h = get_valid_throw(sector)
    sum2 += e + g + h

print("Итоговый результат Игорька 1:", sum1)
print("Итоговый результат Игорька 2:", sum2)

if sum1 > sum2:
    print("Выиграл Игорёк 1")
elif sum1 < sum2:
    print("Выиграл Игорёк 2")
else:
    print("Ничья")
