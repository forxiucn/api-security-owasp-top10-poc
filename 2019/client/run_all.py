import requests
import json

BASE_URL = 'http://127.0.0.1:5001'

def test_api1():
    r = requests.get(f'{BASE_URL}/api1/items/123')
    print('API1:', r.status_code, r.json())

def test_api2():
    r = requests.post(f'{BASE_URL}/api2/login', json={'username': 'admin', 'password': '123456'})
    print('API2:', r.status_code, r.json())
    r2 = requests.post(f'{BASE_URL}/api2/login', json={'username': 'admin', 'password': 'wrong'})
    print('API2 (fail):', r2.status_code, r2.json())

def test_api3():
    r = requests.get(f'{BASE_URL}/api3/userinfo')
    print('API3:', r.status_code, r.json())

def test_api4():
    for i in range(3):
        r = requests.get(f'{BASE_URL}/api4/nolimit')
        print(f'API4 [{i+1}]:', r.status_code, r.json())

def test_api5():
    r = requests.get(f'{BASE_URL}/api5/admin')
    print('API5:', r.status_code, r.json())

def test_api6():
    r = requests.post(f'{BASE_URL}/api6/profile', json={'username': 'bob', 'is_admin': True})
    print('API6:', r.status_code, r.json())

def test_api7():
    r = requests.get(f'{BASE_URL}/api7/debug')
    print('API7:', r.status_code, r.json())

def test_api8():
    r = requests.post(f'{BASE_URL}/api8/search', json={'q': "1' OR '1'='1"})
    print('API8:', r.status_code, r.json())

def test_api9():
    r = requests.get(f'{BASE_URL}/api9/old-api')
    print('API9:', r.status_code, r.json())

def test_api10():
    r = requests.post(f'{BASE_URL}/api10/transfer', json={'from': 'alice', 'to': 'bob', 'amount': 100})
    print('API10:', r.status_code, r.json())

def main():
    test_api1()
    test_api2()
    test_api3()
    test_api4()
    test_api5()
    test_api6()
    test_api7()
    test_api8()
    test_api9()
    test_api10()

if __name__ == '__main__':
    main()
