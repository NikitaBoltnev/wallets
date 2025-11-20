#!/bin/bash

# Добавление в существующий кошелек 1000
echo "Добавление 1000 в существующий кошелек"
echo "Ожидается баланс: 1000"
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{"valletId": "ff41c5c7-4f9e-4161-a573-7e3b73d2cd55", "operationType": "DEPOSIT", "amount": 1000}'
echo ""

# Получение баланса в существующем кошельке 
echo "Получение баланса в существующем кошельке"
echo "Ожидается баланс: 1000"
curl -X GET http://localhost:8080/api/v1/wallets/ff41c5c7-4f9e-4161-a573-7e3b73d2cd55
echo ""

# Добавление в существующий кошелек 12.34
echo "Добавление 12.34 в существующий кошелек"
echo "Ожидается баланс: 1012.34"
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{"valletId": "ff41c5c7-4f9e-4161-a573-7e3b73d2cd55", "operationType": "DEPOSIT", "amount": 12.34}'
echo ""

# Получение баланса в существующем кошельке 
echo "Получение баланса в существующем кошельке"
echo "Ожидается баланс: 1012.34"
curl -X GET http://localhost:8080/api/v1/wallets/ff41c5c7-4f9e-4161-a573-7e3b73d2cd55
echo ""

# Вычитание из существующего кошелька 567
echo "Вычитание из существующего кошелька 567.08"
echo "Ожидается баланс: 445.26"
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{"valletId": "ff41c5c7-4f9e-4161-a573-7e3b73d2cd55", "operationType": "WITHDRAW", "amount": 567.08}'
echo ""

# Вычитание из существующего кошелька суммы превышающей баланс (баланс:445.26, вычитание: 500)
echo "Вычитание из существующего кошелька суммы превышающей баланс (баланс:445.26, вычитание: 500)"
echo "Ожидается код ошибки: INSUFFICIENT_FUNDS"
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{"valletId": "ff41c5c7-4f9e-4161-a573-7e3b73d2cd55", "operationType": "WITHDRAW", "amount": 500}'
echo ""

# Получение баланса в существующем кошельке 
echo "Получение баланса в существующем кошельке"
echo "Ожидается баланс: 445.26"
curl -X GET http://localhost:8080/api/v1/wallets/ff41c5c7-4f9e-4161-a573-7e3b73d2cd55
echo ""

# Попытка добавления в несуществующий кошелек 1000
echo "Попытка добавления в несуществующий кошелек 1000"
echo "Ожидается код ошибки: WALLET_NOT_FOUND"
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{"valletId": "b34c3acb-ae8c-48a8-bacb-641de595e51f", "operationType": "DEPOSIT", "amount": 1000}'
echo ""

# Попытка вычитания из несуществующего кошелька 500
echo "Попытка вычитания из несуществующего кошелька 500"
echo "Ожидается код ошибки: WALLET_NOT_FOUND"
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{"valletId": "b34c3acb-ae8c-48a8-bacb-641de595e51f", "operationType": "WITHDRAW", "amount": 500}'
echo ""

# Попытка получения баланса в несуществующем кошельке 
echo "Попытка получения баланса в несуществующем кошельке"
echo "Ожидается код ошибки: WALLET_NOT_FOUND"
curl -X GET http://localhost:8080/api/v1/wallets/b34c3acb-ae8c-48a8-bacb-641de595e51f
echo ""

# Попытка выполнить несуществущую операцию на существующем кошельке
echo "Попытка выполнить несуществущую операцию на существующем кошельке"
echo "Ожидается код ошибки: INVALID_OPERATION_TYPE"
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{"valletId": "ff41c5c7-4f9e-4161-a573-7e3b73d2cd55", "operationType": "TRANSFER", "amount": 1000}'
echo ""

# Попытка выполнить операцию с невалидным UUID
echo "Попытка выполнить операцию с невалидным UUID"
echo "Ожидается код ошибки: INVALID_WALLET_ID"
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{"valletId": "not-uuid-01234", "operationType": "DEPOSIT", "amount": 1000}'
echo ""

# Попытка выполнить операцию с пустым UUID
echo "Попытка выполнить операцию с пустым UUID"
echo "Ожидается код ошибки: WALLET_ID_EMPTY"
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{"valletId": "", "operationType": "DEPOSIT", "amount": 1000}'
echo ""

# Попытка выполнить DEPOSIT с отрицательным amount
echo "Попытка выполнить DEPOSIT с отрицательным amount"
echo "Ожидается код ошибки: AMOUNT_MUST_BE_POSITIVE"
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{"valletId": "ff41c5c7-4f9e-4161-a573-7e3b73d2cd55", "operationType": "DEPOSIT", "amount": -100}'
echo ""

# Попытка выполнить WITHDRAW, с отрицательным amount
echo "Попытка выполнить WITHDRAW, с отрицательным amount"
echo "Ожидается код ошибки: AMOUNT_MUST_BE_POSITIVE"
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{"valletId": "ff41c5c7-4f9e-4161-a573-7e3b73d2cd55", "operationType": "WITHDRAW", "amount": -100}'
echo ""

# Попытка выполнить операцию с пустым UUID
echo "Попытка выполнить операцию с amount превышающим 2 знака после запятой"
echo "Ожидается код ошибки: AMOUNT_PRECISION_INVALID"
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{"valletId": "ff41c5c7-4f9e-4161-a573-7e3b73d2cd55", "operationType": "DEPOSIT", "amount": 100.123}'
echo ""

# Получение баланса в существующем кошельке 
echo "Получение баланса в существующем кошельке"
echo "Ожидается баланс: 445.26"
curl -X GET http://localhost:8080/api/v1/wallets/ff41c5c7-4f9e-4161-a573-7e3b73d2cd55
echo ""

# Попытка отправить невалидный JSON
echo "Попытка отправить невалидный JSON"
echo "Ожидается код ошибки: INVALID_JSON"
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{"valletId": "ff41c5c7-4f9e-4161-a573-7e3b73d2cd55", "operationType": "DEPOSIT", "amount": 100.12,}'
echo ""

# Получение баланса в существующем кошельке 
echo "Получение баланса в существующем кошельке"
echo "Ожидается баланс: 445.26"
curl -X GET http://localhost:8080/api/v1/wallets/ff41c5c7-4f9e-4161-a573-7e3b73d2cd55
echo ""

#  1000 запросов на пополнение кошелька
echo "1000 запросов на пополнение кошелька"
echo "Каждый запрос добавляет 3.01 к балансу существующего кошелька"
echo "Ожидается баланс: 3455.26"
hey -n 1000 -c 100 \
  -m POST \
  -H "Content-Type: application/json" \
  -d '{"valletId": "ff41c5c7-4f9e-4161-a573-7e3b73d2cd55", "operationType": "DEPOSIT", "amount": 3.01}' \
  http://localhost:8080/api/v1/wallet
echo ""

# Получение баланса после 1000 операций 
echo "Получение баланса в существующем кошельке после 1000 операций"
echo "Ожидается баланс: 3455.26"
curl -X GET http://localhost:8080/api/v1/wallets/ff41c5c7-4f9e-4161-a573-7e3b73d2cd55
echo ""


#  1000 запросов на вычитание из кошелька
echo "1000 запросов на вычитание из кошелька"
echo "Каждый запрос вычитает 2.79 из баланса существующего кошелька"
echo "Ожидается баланс: 665.26"
hey -n 1000 -c 100 \
  -m POST \
  -H "Content-Type: application/json" \
  -d '{"valletId": "ff41c5c7-4f9e-4161-a573-7e3b73d2cd55", "operationType": "WITHDRAW", "amount": 2.79}' \
  http://localhost:8080/api/v1/wallet
echo ""

# Получение баланса после 1000 операций 
echo "Получение баланса в существующем кошельке после 1000 операций"
echo "Ожидается баланс: 665.26"
curl -X GET http://localhost:8080/api/v1/wallets/ff41c5c7-4f9e-4161-a573-7e3b73d2cd55
echo ""

