import pandas as pd
import matplotlib.pyplot as plt

# Чтение данных из файлов
file1 = 'noindex.csv'
file2 = 'index.csv'

data1 = pd.read_csv(file1, sep=",")
data2 = pd.read_csv(file2, sep=",")

# Построение графика
plt.figure(figsize=(10, 6))
plt.plot(data1['count'], data1['time'], label='Без индекса', marker='o')
plt.plot(data2['count'], data2['time'], label='С индексом', marker='*')

# Настройка осей и заголовка
plt.xlabel('количество записей')
plt.ylabel('время, мкс')
plt.xscale('log')
plt.legend()
plt.grid(True)

plt.savefig('timeInd.pdf')

# Отображение графика
plt.show()
