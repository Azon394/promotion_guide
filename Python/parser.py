import requests
import json
from models import *

cookies = {
    '_gid': 'GA1.2.2135665031.1716296453',
    '_ym_uid': '1716296453816682016',
    '_ym_d': '1716296453',
    '_ym_isad': '2',
    'adrdel': '1716296453021',
    'adrcid': 'AayA0UvQ5bT-gcqbncR2DiA',
    '_ym_visorc': 'w',
    'acs_3': '%7B%22hash%22%3A%223c8f85edb06b1f745fbd%22%2C%22nextSyncTime%22%3A1716382853056%2C%22syncLog%22%3A%7B%22224%22%3A1716296453056%2C%221228%22%3A1716296453056%2C%221230%22%3A1716296453056%7D%7D',
    'visitorcity': 'simferopol-krym',
    'fid': '9619a980-ad29-4329-b8ab-11a8e5b3b1b5',
    'subscribe_showed': '1',
    '_ac_oid': 'b492b9f7443160033b35c33264508296%3A1716300089594',
    'gnezdo_uid': 'XV9maWZMmwQ1Mz9cQyNLAg==',
    'cf_clearance': 'Fn4BUtwaapIhpaKtwe5D2pbOqJkC26kC18TK9j5z.wU-1716299082-1.0.1.1-ee_8A2PcQ08oAsjfaJmiGWftCJ6bolxASN1F_2UDxLWHFKyAxlbQaTYArsv.Fm3aVNtsBWn.CLXjjjZGj6XHvQ',
    'PHPSESSID': '8lo4qt6mgecevukp5v4ip6o75m',
    '_gat': '1',
    '_ga_VXB631NEPB': 'GS1.2.1716299440.4.1.1716299529.0.0.0',
    'FCNEC': '%5B%5B%22AKsRol-4b5cmK-D0PLwnL_PEEIdCaOYVZfx5XpIhqAc-HWPvFakhwtAdKwyX9zpOKYX0IXjZEEuFlNuFVzAnqcxvMuKfY8A8MwKKH6f2TblFcEZG3FomWvYOLN7bWNJlJEMO8eEwQO7i6JCnhwKh5JJ0xbl6QZirtQ%3D%3D%22%5D%5D',
    '_ga_YL18X8Z772': 'GS1.1.1716296452.1.1.1716299558.0.0.0',
    '_ga': 'GA1.1.1563663131.1716296453',
}

alcohol_response = requests.get(
    'https://skidkaonline.ru/apiv3/products/?limit=100&offset=0&pcategories_ids=499676&city_id=627&fields=id,name,shops_ids,imagefull,daystitle',
    cookies=cookies,
)

products_response = requests.get(
    'https://skidkaonline.ru/apiv3/products/?limit=1000&offset=0&pcategories_ids=412,422&city_id=627&fields=id,name,shops_ids,imagefull,daystitle',
    cookies=cookies,
)

candy_response = requests.get(
    'https://skidkaonline.ru/apiv3/products/?limit=200&offset=0&pcategories_ids=4748&city_id=627&fields=id,name,shops_ids,imagefull,daystitle',
    cookies=cookies,
)

bitovuha_response = requests.get(
    'https://skidkaonline.ru/apiv3/products/?limit=100&offset=0&pcategories_ids=630&city_id=627&fields=id,name,shops_ids,imagefull,daystitle',
    cookies=cookies,
)

meat_response = requests.get(
    'https://skidkaonline.ru/apiv3/products/?limit=150&offset=0&pcategories_ids=4628&city_id=627&fields=id,name,shops_ids,imagefull,daystitle',
    cookies=cookies,
)

coffe_response = requests.get(
    'https://skidkaonline.ru/apiv3/products/?limit=100&offset=0&pcategories_ids=1868&city_id=627&fields=id,name,shops_ids,imagefull,daystitle',
    cookies=cookies,
)

feed_response = requests.get(
    'https://skidkaonline.ru/apiv3/products/?limit=30&offset=0&pcategories_ids=3367&city_id=627&fields=id,name,shops_ids,imagefull,daystitle',
    cookies=cookies,
)

powder_response = requests.get(
    'https://skidkaonline.ru/apiv3/products/?limit=30&offset=0&pcategories_ids=2754&city_id=627&fields=id,name,shops_ids,imagefull,daystitle',
    cookies=cookies,
)

desert_response = requests.get(
    'https://skidkaonline.ru/apiv3/products/?limit=429&offset=0&pcategories_ids=2754&city_id=627&fields=id,name,shops_ids,imagefull,daystitle',
    cookies=cookies,
)


alcohol_info = Items.parse_obj(alcohol_response.json())
product_info = Items.parse_obj(products_response.json())
candy_info = Items.parse_obj(candy_response.json())
bitovuha_info = Items.parse_obj(bitovuha_response.json())
meat_info = Items.parse_obj(meat_response.json())
coffe_info = Items.parse_obj(coffe_response.json())
feed_info = Items.parse_obj(feed_response.json())
powder_info = Items.parse_obj(powder_response.json())
desert_info = Items.parse_obj(desert_response.json())

print(alcohol_info.json())

adddata = requests.post('http://localhost:6969/addalc', json=alcohol_info.json())
adddata = requests.post('http://localhost:6969/addprod', json=powder_info.json())
adddata = requests.post('http://localhost:6969/addcandy', json=candy_info.json())
adddata = requests.post('http://localhost:6969/addbit', json=bitovuha_info.json())
adddata = requests.post('http://localhost:6969/addmeat', json=meat_info.json())
adddata = requests.post('http://localhost:6969/addcof', json=coffe_info.json())
adddata = requests.post('http://localhost:6969/addfeed', json=feed_info.json())
adddata = requests.post('http://localhost:6969/addpowder', json=powder_info.json())
adddata = requests.post('http://localhost:6969/adddes', json=desert_info.json())
print(adddata)