# pylint: disable=invalid-name
"""
TextAnalysis - анализатор темы текстов.
Использование:
    Аргумент -h             Отобразить это сообщение
    Аргумент -i <файл>      Обработать текст из файла
    add <тема>              Добавить темы
    remove <тема>           Удалить тему
    list                    Отобразить сохраненные темы
    text <text>             Обработать введенный текст
    Аргумент -t <файл>      Использовать свой словарь тем

Полное использование:
Usage:
    TextAnalysis (-h|--help|--version)
    TextAnalysis [-i=<path>]    [-t=<path>]
    TextAnalysis add <theme>    [-t=<path>]
    TextAnalysis remove <theme> [-t=<path>]
    TextAnalysis list           [-t=<path>]
    TextAnalysis text [-t=<path>] <text>

Options:
    <theme>                   Название темы
    <text>                    Обычный текст
    -i --input_file=<path>    Путь файла с текстом
    -t --themes_file=<path>   Путь файла с темами

"""
import socketserver

from docopt import docopt

from __init__ import __version__
from main import main
# Модуль socketserver для сетевого программирования
from socketserver import *

# данные сервера
host = 'localhost'
port = 777
addr = (host, port)

DEFAULT_ARG = {
    '--themes_file': './demo_data/themes.json',
    '--input_file': './demo_data/input.txt',
}


class MyTCPHandler(StreamRequestHandler):
    # функция handle делает всю работу, необходимую для обслуживания запроса.
    # доступны несколько атрибутов: запрос доступен как self.request, адрес как self.client_address, экземпляр сервера как self.server
    def handle(self):
        self.data = self.rfile.readlines()
        print('client send: ' + str(self.data))

        # sndall - отправляет сообщение
        self.request.sendall(b'Hello from server!')


if __name__ == '__main__':

    with socketserver.TCPServer(addr, MyTCPHandler) as server:
        server.serve_forever()
    args = docopt(__doc__, version=__version__)
    # Запуск без параметров
    args['use_local_files'] = 0  # Флаг использования файлов по умолчанию
    if not args['--themes_file']:
        args['--themes_file'] = DEFAULT_ARG['--themes_file']
        args['use_local_files'] += 1
    if not args['--input_file']:
        args['--input_file'] = DEFAULT_ARG['--input_file']
        args['use_local_files'] += 2
    # Запуск
    main(args)
