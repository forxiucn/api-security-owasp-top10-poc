from flask import Flask, request, jsonify

app = Flask(__name__)

# 1. API1: 对象级授权失效Broken Object Level Authorization (BOLA)
@app.route('/api1/items/<item_id>', methods=['GET'])
def api1(item_id):
    return jsonify({'item_id': item_id, 'detail': 'Object info (no auth check)'})

# 2. API2: 身份认证失效Broken Authentication
@app.route('/api2/login', methods=['POST'])
def api2():
    data = request.json
    if data and data.get('username') == 'admin' and data.get('password') == '123456':
        return jsonify({'msg': 'Login success', 'token': 'fake-jwt-token'})
    return jsonify({'msg': 'Login failed'}), 401

# 3. API3: 对象属性级授权失效Broken Object Property Level Authorization (BOPLA)
@app.route('/api3/userinfo', methods=['GET'])
def api3():
    return jsonify({'username': 'alice', 'email': 'alice@example.com', 'role': 'admin', 'salary': 10000})

# 4. API4: 资源消耗无限制Unrestricted Resource Consumption
@app.route('/api4/nolimit', methods=['GET'])
def api4():
    return jsonify({'msg': 'No resource limit'})

# 5. API5: 功能级授权失效Broken Function Level Authorization
@app.route('/api5/admin', methods=['GET'])
def api5():
    return jsonify({'msg': 'Admin function accessed'})

# 6. API6: 敏感业务流程访问无限制Unrestricted Access to Sensitive Business Flows
@app.route('/api6/transfer', methods=['POST'])
def api6():
    data = request.json
    return jsonify({'msg': 'Business flow executed', 'data': data})

# 7. API7: 服务器端请求伪造（SSRF）Server Side Request Forgery (SSRF)
@app.route('/api7/ssrf', methods=['POST'])
def api7():
    url = request.json.get('url', '')
    return jsonify({'msg': f'Requested URL: {url}'})

# 8. API8: 安全配置错误Security Misconfiguration
@app.route('/api8/debug', methods=['GET'])
def api8():
    return jsonify({'debug': True, 'config': 'Sensitive config here'})

# 9. API9: 资产 inventory 管理不当Improper Inventory Management
@app.route('/api9/old-api', methods=['GET'])
def api9():
    return jsonify({'msg': 'Deprecated API still accessible'})

# 10. API10: 不安全的API调用Unsafe Consumption of APIs
@app.route('/api10/unsafe', methods=['POST'])
def api10():
    data = request.json
    return jsonify({'msg': 'Unsafe API consumed', 'data': data})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5023, debug=True)
