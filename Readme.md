# Ecommerce web API
1. **Getting all products** # method [GET]
Path:    /api/products 
Returs : 
        ## {
    "products": [
        {
            "id": 7,
            "name": "Olma",
            "price": 10.6
        }
    ]
}

2. **Adding products** method [POST]
**Path:**  /api/add/product
**Parameters:** {
    "name": "Olma",
    "price": 10.6
}

**Result:**  {
    "message": "Muvaffaqiyatli qo'shildi!"
}

3. **Updating product** method [POST]
**Path** : /api/update/product
**Parameters:** {
    "id": 1,
    "name": "yangi nok",
    "price": 10.5
}

**Result:**  {
    "message": "Muvaffaqiyatli o'zgartirildi!"
}

4. **Deleting products** method [DELETE]
**Path:** /api/delete/:id
**Result:** {"message": "Muvaffaqiyatli o'chirildi!"}

5. **Buying Product** method [POST]
**Path:** /api/buy
**Parameters:** {
    "id": id (int),
    "firstname": "John",
    "lastname": "Doe"
}
**Result:** {"message": "Muvaffaqiyatli harid qilindi!"}

6. **Getting orders excel file**
**Path:** /api/download
**Result:** excel file


