from django.shortcuts import render
from django.http import HttpResponse
from django.views.decorators.csrf import csrf_exempt
import secrets
import requests
import json
from rest_framework_jwt.utils import jwt_encode_handler
from django.views.decorators.http import require_POST
import base64
import hashlib
from werkzeug.utils import secure_filename
from django.http import JsonResponse
import json
import re

adr = "http://localhost:6969"

def main(request):
    if 'user_id' in request.session:
            url = adr + "/getall?type=alc"
            response = requests.get(url)
            data_str = response.text
            data_dict = json.loads(data_str)
            first_product = data_dict[0]
            product_name = first_product['name']
            product_image_src = first_product['imagefull'][0]['src']
            products = {product_name, product_image_src}
            print(products)

            return render(request, 'main/index_auth_complete.html')
    else:
        # url = adr + "/getall?type=alc"
        # response = requests.get(url)
        # data = response.text
        # with open('output.json', 'w', encoding='utf-8') as file:
        #     json.dump(data, file, ensure_ascii=False, indent=4)
        # # категории 

        # data = " {products [[{daystitle Осталось <span>14</span> дней} {imagefull [{h 1376} {src https://skidkaonline.ru/img/p/2024/06/517581/58125754-517581-378417182722175285.jpg?t=t1718272291} {w 1024}]} {name Бренди (Коньяк)} {shops_ids [1773]}]"

        # # Регулярное выражение для поиска продуктов
        # product_pattern = r"\{.*?name\s*(.*?)\s*\}\s*{shops_ids \[\w+\]\s*}.*?daystitle\s*Осталось\s*<span>\d+</span>\s*дней.*?imagefull \[(.*?)\s*{src (.*?)\s*w (\d+)\].+?"

        # print('ДО ПОИСКА ПРОДУКТОВ')
        # # Поиск всех продуктов
        # products = re.findall(product_pattern, data)
        # print("до преобразование результато в словари")
        # # Преобразование результатов в словари
        # for product in products:
        #     name, image_full, src, width = product
        #     # Удаление лишних пробелов и переносов строк
        #     name = name.strip()
        #     image_full = image_full.strip().replace(" ", "")
        #     src = src.strip()
        #     width = int(width)
    
        # print(f"Название: {name}, Изображение: {image_full}, URL: {src}, Ширина: {width}")


        return render(request, 'main/index.html')

def reg(request):
    if request.method == 'POST':
        if 'registration' in request.POST:
            def generate_jwt_token(login, password):
                header = {
                    'alg': 'HS256',
                    'typ': 'JWT'
                }
                payload = {
                    'login': login,
                    'password': password,
                    'exp': 600
                }
                header_encoded = base64.urlsafe_b64encode(json.dumps(header).encode()).decode().strip('=')
                payload_encoded = base64.urlsafe_b64encode(json.dumps(payload).encode()).decode().strip('=')
                secret = '123456789'
                signature = hashlib.sha256(f'{header_encoded}.{payload_encoded}{secret}'.encode()).hexdigest()
                token = f'{header_encoded}.{payload_encoded}.{signature}'
                return token


            login = request.POST['login']
            password = request.POST['password']
            token = generate_jwt_token(login, password)

            url = adr + f"/reg?token={token}"
            response = requests.get(url)
            answer = response.text

            if answer == "успешная регистрация":
                request.session['user_id'] = secrets.token_hex(16)
                request.session['login'] = login
                request.session['password'] = password
                return render(request, 'main/index_auth_complete.html')
            if answer == "такой пользователь уже существует":
                text = "Такой пользователь уже гарегистрирован. Попробуйте войти в акканут или введите другие данные"
                return render(request, 'main/reg.html', {'text': text})
    else:
        return render(request, 'main/reg.html')

def auth(request):
    if request.method == 'POST':
        if 'auth' in request.POST:
            def generate_jwt_token(login, password):
                header = {
                    'alg': 'HS256',
                    'typ': 'JWT'
                }
                payload = {
                    'login': login,
                    'password': password,
                    'exp': 600
                }
                header_encoded = base64.urlsafe_b64encode(json.dumps(header).encode()).decode().strip('=')
                payload_encoded = base64.urlsafe_b64encode(json.dumps(payload).encode()).decode().strip('=')
                secret = '123456789'
                signature = hashlib.sha256(f'{header_encoded}.{payload_encoded}{secret}'.encode()).hexdigest()
                token = f'{header_encoded}.{payload_encoded}.{signature}'
                return token


            login = request.POST['login']
            password = request.POST['password']
            token = generate_jwt_token(login, password)

            url = adr + f"/auth?token={token}"
            response = requests.get(url)
            answer = response.text

            if answer == "успешная авторизация":
                request.session['user_id'] = secrets.token_hex(16)
                request.session['login'] = login
                request.session['password'] = password
                return render(request, 'main/index_auth_complete.html')
            if answer == "неверный пароль":
                text = "Неверный пароль, попробуйте еще раз"
                return render(request, 'main/auth.html', {'text': text})
            if answer == "нет такого пользователя":
                text = "Такого пользователя не существует, попробуйте зарегистрироваться"
                return render(request, 'main/auth.html', {'text': text})
    else:
        return render(request, 'main/auth.html')

def account(request):
    return render(request, "main/account.html")
