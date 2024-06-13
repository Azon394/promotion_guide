from pydantic import BaseModel

# name, daystitle, shops_ids, imagefull
# Модель для парсинга
class Item(BaseModel):
    name: str
    daystitle: str
    shops_ids: list
    imagefull: dict

class Items(BaseModel):
    products: list[Item]


